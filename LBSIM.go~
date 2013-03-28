package main

import (
	"fmt"
	"log"
	"net/http"
	"code.google.com/p/go.net/websocket"
)

func main() {
	http.HandleFunc("/login", Login)
	http.HandleFunc("/", Login)
	http.HandleFunc("/register", Register)
	http.HandleFunc("/onlineUsers", OnlineUsers)
	http.HandleFunc("/templates/", assetsHandler)
	http.HandleFunc("/route", Route)
	http.HandleFunc("/chat", Chat)
	http.Handle("/websocketChat", websocket.Handler(WebsocketChat))
	fmt.Println("listen on port 8888")

	if err := http.ListenAndServe(":8888", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func assetsHandler(w http.ResponseWriter, r *http.Request) {
	//Println(r.URL.Path[len("/")])
	http.ServeFile(w, r, r.URL.Path[len("/"):])
}
