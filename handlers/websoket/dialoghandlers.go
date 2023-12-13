package handlers_websoket

import (
	handlers_http "chat_v2/handlers/http"
	"chat_v2/models"
	"chat_v2/scylla"
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

var (
	clienPage = template.Must(template.ParseFiles("client.html"))
)

type clients struct {
	Sender_id   string
	Reciever_id string
	Sender      models.User
	Reciever    models.User
	Messages    []models.Message
}

type CurrentMessage struct {
  Username string `json:"username"`
  Message  string `json:"message"`
}

func DialogPageHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	sender, _ := scylla.FindCurrentUser(vars["user_id"], scylla.Session(), scylla.Logger())
	reciever, _ := scylla.FindCurrentUser(vars["reciever_id"], scylla.Session(), scylla.Logger())

	clienPage.Execute(w, clients{
		vars["user_id"],
		vars["reciever_id"],
		sender,
		reciever,
		scylla.FindAllMessages(vars["user_id"], vars["reciever_id"], scylla.Session(), scylla.Logger()),
	})
}

func PeerHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("мы в обработчике диалога")
	vars := mux.Vars(r)
	sender_id := vars["sender_id"]     // "50e0b9f3-a7ab-42eb-b8b6-6c56b9c05acd"
	reciever_id := vars["reciever_id"] // "1d52892f-e753-41b8-b201-0c22d27ef512"
	conn, err := handlers_http.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Ошибка обновления соединения:", err)
		return
	}
	handlers_http.MX.Lock()
	handlers_http.Users[sender_id] = handlers_http.Connection{State: handlers_http.Dialog, Conn: conn}
	handlers_http.MX.Unlock()

	// err = conn.WriteMessage(websocket.TextMessage, []byte("Hello, test"))
	// if err != nil {
	// 	log.Println("err in write message: ", err)
	// } else {
	// 	log.Println("ошибки нет в записи в сокет")
	// }
	log.Printf("sender_id = %s\nreciever_id = %s\n", sender_id, reciever_id)
	time.Sleep(time.Second * 2)

  curentUser, err := scylla.FindCurrentUser(sender_id, scylla.Session(), scylla.Logger())
  if err != nil{
    log.Fatalf("Не удалось вытащить пользователя. Ошибка: %s\n", err)
  }
  message := CurrentMessage{Username: curentUser.Username}

	for {
		// Чтение сообщения от клиента
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("Ошибка чтения сообщения:", err)
			break
		}
		err = scylla.CreateMessage(sender_id, reciever_id, string(msg), scylla.Session(), scylla.Logger())
		if err != nil {
			log.Println("error in CreateMessage: ", err)
		}
		log.Println("msg: ", string(msg))
		
    message.Message = string(msg)

    message_byte, err := json.Marshal(message)
    if err != nil {
      log.Println("error in marshal: ", err)
    }

    err = conn.WriteMessage(websocket.TextMessage, message_byte)
		if err != nil {
			log.Println("Ошибка отправки сообщения НА ОТПРАВИТЕЛЯ:", err)
			// break
		}

		handlers_http.MX.Lock()
		reciever_conn, ok := handlers_http.Users[reciever_id]
		handlers_http.MX.Unlock()
		if ok {
			//log.Fatal("not found in map user_id: ", reciever_id)
			if reciever_conn.State == handlers_http.Dialog {
				err := reciever_conn.Conn.WriteMessage(websocket.TextMessage, message_byte)
				if err != nil {
					log.Println("Ошибка отправки сообщения НА ПОЛУЧАТЕЛЯ:", err)
				}

			}
		}

		// Пример проверки, закрыто ли соединение с клиентом
		// ticker := time.NewTicker(10 * time.Second)
		// defer ticker.Stop()

		// for range ticker.C {
		// 	err := conn.WriteMessage(websocket.PingMessage, nil)
		// 	if err != nil {
		// 		log.Println("Соединение с клиентом закрыто")
		// 		break
		// 	} else {
		// 		//log.Println("пинг прошел успешно")
		// 		err := conn.WriteMessage(websocket.PingMessage, nil)
		// 		if err != nil {
		// 			log.Println("Соединение с клиентом закрыто")
		// 			break
		// 		} else {
		// 			log.Println("пинг прошел успешно")
		// 		}
		//
		// 	}
		// }
		// Отправка обратного сообщения клиенту
	}

}
