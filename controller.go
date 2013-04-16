package main

import (
	"code.google.com/p/go.net/websocket"
	"fmt"
	"log"
	"net/http"
	"strings"
)

var appName string = "LBSIM"
var onlineUsersRefresh = make(chan string, 1)

func init(){
	//defer close(onlineUsersRefresh)
	log.SetFlags(log.Llongfile | log.LstdFlags)
}

func router(rw http.ResponseWriter, req *http.Request) {
	registerURL := "/register"
	loginURL := "/login"
	rootPathURL := "/"
	mainURL := "/main"
	wsOnlineFriendsURL := "/wsOnlineFriends"
	wsChatURL := "/wsChat"

	urlPath := req.URL.Path
	log.Println(urlPath)
	switch req.Method {
	case "GET":
		switch {
		case strings.HasPrefix(urlPath, registerURL):
			registerGet(rw, req)
		case rootPathURL == urlPath || loginURL == urlPath:
			loginGet(rw, req)
		case strings.HasPrefix(urlPath, mainURL):
			mainGet(rw, req)
		case wsOnlineFriendsURL == urlPath:
			websocket.Handler(wsOnlineFriends).ServeHTTP(rw, req)
		case wsChatURL == urlPath:
			websocket.Handler(wsChat).ServeHTTP(rw, req)
		case strings.HasPrefix(urlPath, "/templates"):
			http.ServeFile(rw, req,urlPath[1:])
		case urlPath == "/favicon.ico":
			http.NotFound(rw, req)
		default:
			notFoundHandler(rw, req)
		}
	case "POST":
		switch {
		case rootPathURL == urlPath || loginURL == urlPath:
				loginPost(rw, req)
		case strings.HasPrefix(urlPath, registerURL):
				registerPost(rw, req)
		}
	}
}

func main() {
	//http.Handle("/wsOnlineFriends", websocket.Handler(wsOnlineFriends))
	//http.Handle("/wsChat", websocket.Handler(wsChat))

	fmt.Println("listen on port 8888")

	go observeOnlineFriends()
	if err := http.ListenAndServe(":8888", http.HandlerFunc(router)); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
};
