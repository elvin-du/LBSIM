package controllers

import (
	"html/template"
	"logger"
	"net/http"
	"utils"
)

type Account struct {
	*Controller
}

func NewAccount() *Account {
	return &Account{
		Controller: &Controller{},
	}
}

/*
所有[/account/*]路由的请求，都要经过这里进行转发
*/
func (this *Account) Handler(rw http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	switch req.URL.Path {
	case "/account/login":
		this.LoginHandler(rw, req)
	case "/account/register":
		this.RegisterHandler(rw, req)
	default:
		NotFoundHandler(rw, req)
	}
}

func (this *Account) LoginHandler(rw http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	switch req.Method {
	case "GET":
		t, err := template.ParseFiles(utils.BaseHtmlTplFile, "public/html/account/login.html")
		if nil != err {
			logger.Errorln(err)
			break
		}
		err = t.Execute(rw, nil)
		if nil != err {
			logger.Errorln(err)
		}
	case "POST":
		username := req.FormValue("username-login")
		password := req.FormValue("password-login")

	default:
		NotFoundHandler(rw, req)
	}
}

func (this *Account) RegisterHandler(rw http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	switch req.Method {
	case "GET":
	case "POST":
	default:
		NotFoundHandler(rw, req)
	}
}
