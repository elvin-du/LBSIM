package main

import (
	"fmt"
	"log"
	"net/http"
	"code.google.com/p/go.net/websocket"
)

var appName string = "LBSIM"
var onlineUsersRefresh = make(chan bool,100)

func main() {
	defer close(onlineUsersRefresh)
	log.SetFlags(log.Llongfile | log.LstdFlags)
	http.HandleFunc("/login", Login)
	http.HandleFunc("/", NotFoundHandler)
	http.HandleFunc("/register", Register)
	http.HandleFunc("/onlineUsers", OnlineUsers)
	http.HandleFunc("/route", Route)
	http.HandleFunc("/chat", Chat)
	http.Handle("/wsOnlineUsers", websocket.Handler(WsOnlineUsers))
	http.Handle("/wsChat", websocket.Handler(WsChat))
	fmt.Println("listen on port 8888")

	go UpdateOnlineUsers()
	if err := http.ListenAndServe(":8888", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func checkError(err error) {
	if err != nil {
		log.Fatal("checkError", err.Error())
	}
}
