package main

import (
	"fmt"
	"log"
	"net/http"
	"code.google.com/p/go.net/websocket"
)

func main() {
	log.SetFlags(log.Llongfile | log.LstdFlags)
	http.HandleFunc("/login", Login)
	http.HandleFunc("/", NotFoundHandler)
	http.HandleFunc("/register", Register)
	http.HandleFunc("/onlineUsers", OnlineUsers)
	http.HandleFunc("/route", Route)
	http.HandleFunc("/chat", Chat)
	http.Handle("/websocketChat", websocket.Handler(WebsocketChat))
	fmt.Println("listen on port 8888")

	if err := http.ListenAndServe(":8888", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func checkError(err error) {
	if err != nil {
		log.Fatal("checkError", err.Error())
	}
}
