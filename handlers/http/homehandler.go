package handlers_http

import (
	"chat_v2/models"
	"chat_v2/scylla"
	"html/template"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

var (
	mainpage = template.Must(template.ParseFiles("mainpage.html"))
)
var Upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type CurrentUser struct {
	Me      models.User
	Friends []models.User
}

type Connection struct {
	State string
	Conn  *websocket.Conn
}

const (
	Chat   = "chat"
	Dialog = "dialog"
)

var (
	Users = make(map[string]Connection)
	MX    = sync.Mutex{}
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	var user_id string
	cookie, err := r.Cookie("user_id")
	if err != nil {
		// надо будет просить зарегаться, но пока без этого
		user_id = uuid.NewString()
		username := "Opera Browser"

		if err := scylla.CreateUser(user_id, username, scylla.Session(), scylla.Logger()); err != nil {
			scylla.Logger().Warn("ошибка в СОЗДАНИИ пользователя. Err: %s\n", zap.Error(err))
			return
		}
		// Почему то устанавливается всё равно до закрытия браузера, а не на постоянную основу
		http.SetCookie(w, &http.Cookie{Name: "user_id",
			Value:    user_id,
			MaxAge:   30 * 24 * 3600,
			Expires:  time.Now().Add(30 * 24 * time.Hour),
			HttpOnly: true,
			Secure:   false,
		})

	} else {
		user_id = cookie.Value
	}
	// user, err := scylla.FindCurrentUser(user_id, scylla.Session(), scylla.Logger())
	// if err != nil {
	// 	// не знаю когда так может произойти, но может и появятся случаи
	// 	// например удалят чела если
	//     scylla.Logger().Warn("ошибка в получении пользователя из бд")
	// 	return
	// }
	index := -1
	users := scylla.FindAll(scylla.Session(), scylla.Logger())
	for i, v := range users {
		if v.User_id == user_id {
			index = i
			break
		}
	}
	if index == -1 {
		scylla.Logger().Error("not found user_id in users slice. нужно на авторизацию кинуть")
		return
	}
	log.Println("end")
	u := models.User{User_id: users[index].User_id, Username: users[index].Username}
	log.Printf("user: {%s, %s}; id = %d\n", u.User_id, u.Username, index)
	err = mainpage.Execute(w, CurrentUser{Me: u, Friends: append(users[:index], users[index+1:]...)})
	scylla.Logger().Info("err execute mainpage = %s", zap.Error(err))
}

func MainHandler(w http.ResponseWriter, r *http.Request) {
	var user_id string
	cookie, err := r.Cookie("user_id")
	if err != nil {
		log.Fatal("не смог вытащить куки для вебсокета")
	}
	user_id = cookie.Value

	log.Println("значение уже в мапе")
	scylla.Logger().Info("логер говорит что мы в вебсокет обработчике")
	conn, err := Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Ошибка обновления соединения:", err)
		return
	}
	MX.Lock()
	Users[user_id] = Connection{Chat, conn}
	MX.Unlock()
}
