package main

import (
	"code.google.com/p/go.net/websocket"
	"time"
	"html/template"
	"log"
	"net/http"
)

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("notFoundHandler")
	r.ParseForm()

	t, err := template.ParseFiles("templates/html/error.html")
	if err != nil {
		log.Println(err)
	}
	t.Execute(w, nil)
}

func registerGet(w http.ResponseWriter, r *http.Request) {
	log.Println("registerGet")
	r.ParseForm()

	t, err := template.ParseFiles("templates/html/register.html")
	CheckError(err)
	err = t.Execute(w, nil)
	CheckError(err)
}

func loginGet(w http.ResponseWriter, r *http.Request) {
	log.Println("loginGet")
	r.ParseForm()
	if _,_,err:= CheckCookie(r); err == nil{
			url := "/main"
			http.Redirect(w, r, url, http.StatusFound)
			return
	}else{
			log.Println(err)
	}

	t, err := template.ParseFiles("templates/html/login.html")
	CheckError(err)
	err = t.Execute(w, nil)
	CheckError(err)
}

func mainGet(w http.ResponseWriter, r *http.Request) {
	log.Println("onlineFriendsGet")
	r.ParseForm()
	if name,_,err:= CheckCookie(r); err == nil{
			AddOnlineFriend(name, 0, 0)
	}else{
			url := "/login"
			http.Redirect(w, r, url, http.StatusFound)
			log.Println(err)
			return
	}

	t, err := template.ParseFiles("templates/html/main.html")
	CheckError(err)
	err = t.Execute(w, allOnlineFriend)
	CheckError(err)
}

func wsChat(ws *websocket.Conn) {
	log.Println("wsChat")

	req := ws.Request()
	if name,_, err := CheckCookie(req); err == nil{
			log.Println(name)
			InsertWsChatConnData(name, ws)
	}else{
			log.Println(err)
			return
	}

	var err error
	var toWho *websocket.Conn
	var rcvMsg string

	for {
		if err = websocket.Message.Receive(ws, &rcvMsg); err != nil {
			log.Println(err)
			break
		}
		log.Println("Received : " + rcvMsg)

		name, content, err := ParseRcvMsg(rcvMsg)
		toWho = GetConnByName(name)
		if err = websocket.Message.Send(toWho, content); err != nil {
			log.Println(err)
			break
		}
	}
}

func wsOnlineFriends(ws *websocket.Conn){
		log.Println("wsOnlineFriends")

		req := ws.Request()
		if name,_, err := CheckCookie(req); err == nil{
				log.Println(name)
				InsertWsOnlineFriendConnData(name, ws)
		}else{
				log.Println(err)
				return
		}
		var rcvMsg string
		for{
				time.Sleep(10000)
				if err := websocket.Message.Receive(ws, &rcvMsg); err != nil{
						log.Println(err)
						break
				}
		}
}

func registerPost(w http.ResponseWriter, r *http.Request) {
	log.Println("registerPost")
	r.ParseForm()
	var regRet interface{} = nil
	type reg struct {
		RegisterResult    string
		RegisterReturnMsg string
	}

	username := r.FormValue("username")
	password := r.FormValue("password")
	pwConfirm := r.FormValue("passwordConfirm")
	err := AddUser(username, password, pwConfirm)
	if  nil != err {
			log.Println(err)
			regRet = reg{RegisterResult: "registerFailed", RegisterReturnMsg:err.Error()}
	}else{
			regRet = reg{RegisterResult: "registerSuccessful", RegisterReturnMsg:"congratulation, register successfully"}
	}

	t, err := template.ParseFiles("templates/html/register.html")
	CheckError(err)
	err = t.Execute(w, regRet)
	CheckError(err)
}

func loginPost(w http.ResponseWriter, r *http.Request) {
	log.Println("loginPost")
	r.ParseForm()

	var data interface{}
	username := r.FormValue("username")
	password := r.FormValue("password")

	if CheckUserPassword(username, password) {
			SetCookie(w, username, password)
			AddOnlineFriend(username, 0, 0)
			url := "/main"
			http.Redirect(w, r, url, http.StatusFound)
	} else {
			SetCookie(w, "", "")
			type loginRet struct{
					LoginRet string
			}
			data = loginRet{"wrongUsrPw"}
	}

	t, _ := template.ParseFiles("templates/html/login.html")
	t.Execute(w, data)
}
