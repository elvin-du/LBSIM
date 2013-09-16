package controllers

import (
	"html/template"
	"logger"
	"net/http"
	"utils"
)

type Friends struct {
	*Controller
}

func NewFriends() *Friends {
	return &Friends{
		Controller: &Controller{},
	}
}

func (this *Friends) Handler(rw http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	urlPath := req.URL.Path

	switch {
	case urlPath == "/" || urlPath == "/friends/online":
		this.OnlineHandler(rw, req)
	default:
		NotFoundHandler(rw, req)
	}
}

func (this *Friends) OnlineHandler(rw http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	switch req.Method {
	case "GET":
		t, err := template.ParseFiles(utils.BaseHtmlTplFile, "public/html/friends/online.html")
		if nil != err {
			logger.Errorln(err)
			break
		}
		err = t.Execute(rw, nil)
		if nil != err {
			logger.Errorln(err)
		}
	case "POST":
	default:
		NotFoundHandler(rw, req)
	}
}

func (this *Friends) ChatHandler(rw http.ResponseWriter, req *http.Request) {
}

func (this *Friends) AddHandler(rw http.ResponseWriter, req *http.Request) {
}

func (this *Friends) DelHandler(rw http.ResponseWriter, req *http.Request) {
}

func (this *Friends) SettingHandler(rw http.ResponseWriter, req *http.Request) {
}
