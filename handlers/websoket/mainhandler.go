package handlers_websoket

import (
	handlers_http "chat_v2/handlers/http"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)


func WebsocketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := handlers_http.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Ошибка обновления соединения:", err)
		return
	}
	defer conn.Close()

	// for {
	// Чтение сообщения от клиента
	_, msg, err := conn.ReadMessage()
	if err != nil {
		log.Println("Ошибка чтения сообщения:", err)
		// break
	}

	log.Println("msg: ", string(msg))
	// Пример проверки, закрыто ли соединение с клиентом
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		err := conn.WriteMessage(websocket.PingMessage, nil)
		if err != nil {
			log.Println("Соединение с клиентом закрыто")
			break
		} else {
			//log.Println("пинг прошел успешно")
			err := conn.WriteMessage(websocket.PingMessage, nil)
			if err != nil {
				log.Println("Соединение с клиентом закрыто")
				break
			} else {
				log.Println("пинг прошел успешно")
			}

		}
	}
	// Отправка обратного сообщения клиенту
	err = conn.WriteMessage(websocket.TextMessage, msg)
	if err != nil {
		log.Println("Ошибка отправки сообщения:", err)
		// break
	}
	// }
}

