package main

import (
	"code.google.com/p/go.net/websocket"
	"fmt"
	"log"
	"net/http"
	"strings"
)

var appName string = "LBSIM"
var onlineUsersRefresh = make(chan bool, 100)

func init(){
	//defer close(onlineUsersRefresh)
	log.SetFlags(log.Llongfile | log.LstdFlags)
}

func router(rw http.ResponseWriter, req *http.Request) {
	registerURL := "/register"
	loginURL := "/login"
	rootPathURL := "/"
	onlineFriendsURL := "/onlineFriends"
	routeToFriendURL := "/routeToFriend"
	chatURL := "/chat"

	urlPath := req.URL.Path
	log.Println(urlPath)
	switch req.Method {
	case "GET":
		switch {
		case urlPath == "/style.css" || urlPath == "/favicon.ico":
			http.ServeFile(rw, req, "templates/html/css")
		case strings.HasPrefix(urlPath, registerURL):
			registerGet(rw, req)
		case strings.HasPrefix(urlPath, onlineFriendsURL):
			onlineFriendsGet(rw, req)
		case strings.HasPrefix(urlPath, routeToFriendURL):
			routeToFriendGet(rw, req)
		case strings.HasPrefix(urlPath, chatURL):
			chatGet(rw, req)
		default:
			loginGet(rw, req)
			//notFoundHandler(rw, req)
		}
	case "POST":
		switch {
		case strings.HasPrefix(urlPath, rootPathURL) || strings.HasPrefix(urlPath, loginURL):
				loginPost(rw, req)
		case strings.HasPrefix(urlPath, registerURL):
				registerPost(rw, req)
		case strings.HasPrefix(urlPath, onlineFriendsURL):
				onlineFriendsPost(rw, req)
		}
	}
}

func main() {
	http.Handle("/wsOnlineFriends", websocket.Handler(wsOnlineFriends))
	http.Handle("/wsChat", websocket.Handler(wsChat))

	fmt.Println("listen on port 8888")

	go UpdateOnlineFriends()
	if err := http.ListenAndServe(":8888", http.HandlerFunc(router)); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
};
