package main

import (
	//"code.google.com/p/go.net/websocket"
	"controllers"
	"logger"
	"net/http"
	"strings"
)

const (
	rootUrl    = "/"
	friendsUrl = "/friends/"
	accountUrl = "/account/"
	adminUrl   = "/admin/"
	wsUrl      = "/ws/"
)

func router(rw http.ResponseWriter, req *http.Request) {
	urlPath := req.URL.Path
	println(urlPath)
	logger.Debugln(urlPath)
	switch {
	case rootUrl == urlPath || strings.HasPrefix(urlPath, friendsUrl):
		friends := &controllers.Friends{}
		friends.Handler(rw, req)
	case strings.HasPrefix(urlPath, accountUrl):
		account := &controllers.Account{}
		account.Handler(rw, req)
	case strings.HasPrefix(urlPath, adminUrl):
		admin := &controllers.Admin{}
		admin.Handler(rw, req)
	case strings.HasPrefix(urlPath, wsUrl):
		ws := &controllers.Ws{}
		ws.Handler(rw, req)
	case strings.HasPrefix(urlPath, "public/"): //static files
		http.ServeFile(rw, req, urlPath)
	case urlPath == "/favicon.ico": //the request which browser send automatically
		http.ServeFile(rw, req, "public/images/favicon.ico")
	default:
		controllers.NotFoundHandler(rw, req)
	}
}
