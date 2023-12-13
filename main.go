package main

import (
	handlers_http "chat_v2/handlers/http"
	handlers_websoket "chat_v2/handlers/websoket"
	"chat_v2/scylla"

	"log"
	"net/http"

	"github.com/gocql/gocql"
  "github.com/gorilla/mux"
)

var (
	// кластер не тот, поменять не забудь
  cluster = scylla.CreateCluster(gocql.One, "my_keyspace", "192.168.31.126:9042")//"192.168.1.30:9042")// "scylla-node1", "scylla-node2", "scylla-node3")
)


func init() {
  scylla.Init(*cluster)
}

func main() {
  router := mux.NewRouter()
  router.HandleFunc("/ws", handlers_http.MainHandler)
  router.HandleFunc("/", handlers_http.HomeHandler)
	// router.HandleFunc("/ws", handlers_websoket.WebsocketHandler)
  router.HandleFunc("/{user_id}/dialog/{reciever_id}", handlers_websoket.DialogPageHandler)
	router.HandleFunc("/styles.css", handlers_http.StylesHandler)
  router.HandleFunc("/magic.js", handlers_http.MagicHandler)
  router.HandleFunc("/dialog.js", handlers_http.DialogJSHandler)
  router.HandleFunc("/dialog.ts", handlers_http.DialogTSHandler)
  router.HandleFunc("/peer/{sender_id}/{reciever_id}", handlers_websoket.PeerHandler)
	http.Handle("/", router)//handlers_http.HomeHandler)

	// fs := http.FileServer(http.Dir("./"))
	// http.Handle("/static/", http.StripPrefix("/static/", fs))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Ошибка запуска сервера:", err)
	}
}


