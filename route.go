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
		url := "/onlinefriends"
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

func onlineFriendsGet(w http.ResponseWriter, r *http.Request) {
	log.Println("onlineFriendsGet")
	r.ParseForm()
	if name,_,err:= CheckCookie(r); err == nil{
		AddOnlineFriend(name, -1, -1)//-1 means that do not insert GPS data
	}else{
		url := "/login"
		http.Redirect(w, r, url, http.StatusFound)
		log.Println(err)
		return
	}

	t, err := template.ParseFiles("templates/html/onlineFriends.html")
	CheckError(err)
	err = t.Execute(w, allOnlineFriend)
	CheckError(err)
}

func routeGet(w http.ResponseWriter, r *http.Request) {
	log.Println("routeGet")
	r.ParseForm()
	name,_,err:= CheckCookie(r)
	if err == nil{
		AddOnlineFriend(name, -1, -1)//-1 means that do not insert GPS data
	}else{
		url := "/login"
		http.Redirect(w, r, url, http.StatusFound)
		log.Println(err)
		return
	}
	t, err := template.ParseFiles("templates/html/route.html")
	CheckError(err)

	end := Location{-1,-1}
	withWho := r.Form.Get("withwho")
	if withWho != ""{
		loc := FindLocByName(withWho)//if no withWho, loc is empty
		end = Location{Longitude: loc.Longitude, Latitude: loc.Latitude}
	}
	log.Println(withWho, " is in lng:",end.Longitude,"lat:",end.Latitude)
	err = t.Execute(w, end)
	CheckError(err)
}

func chatGet(w http.ResponseWriter, r *http.Request) {
	log.Println("chatGet")
	r.ParseForm()
	if name,_,err:= CheckCookie(r); err == nil{
		AddOnlineFriend(name, -1, -1)//-1 means that do not insert GPS data
	}else{
		url := "/login"
		http.Redirect(w, r, url, http.StatusFound)
		log.Println(err)
		return
	}

	type withwho struct{
		Name string
	}
	who := r.Form.Get("withwho")

	t, err := template.ParseFiles("templates/html/chat.html")
	CheckError(err)
	err = t.Execute(w, withwho{who})
	CheckError(err)
}

func wsChat(ws *websocket.Conn) {
	log.Println("wsChat")
	req := ws.Request()
	name,_, err := CheckCookie(req)
	if err == nil{
		log.Println(name,"connected by websocket")
		UpdateWsConn(name, "chat", ws)
	}else{
		log.Println(err)
		return
	}

	var toWhoConn *websocket.Conn
	var rcvMsg string
	for {
		if err = websocket.Message.Receive(ws, &rcvMsg); err != nil {
			log.Println(err)
			break
		}
		log.Println("Received : " + rcvMsg)

		toWhoName, content, err := ParseRcvMsg(rcvMsg)
		toWhoConn = GetConnByName(toWhoName)
		content = name + ":" + content
		if nil == toWhoConn{
			continue
		}
		if err = websocket.Message.Send(toWhoConn, content); err != nil {
			log.Println(err)
			//break
		}
	}
}

func wsOnlineFriends(ws *websocket.Conn){
	log.Println("wsOnlineFriends")
	req := ws.Request()
	username,_, err := CheckCookie(req)
	if err == nil{
		log.Println(username ," connected on websocket")
		UpdateWsConn(username,"online", ws)
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
		lng,lat,_:= ParseLngLat(rcvMsg)
		AddOnlineFriend(username,lng,lat)
		log.Println(username,lng,lat)
	}
	//ws disconnected process
}

func registerPost(w http.ResponseWriter, r *http.Request) {
	log.Println("registerPost")
	r.ParseForm()
	var regRet interface{} = nil
	type reg struct {
		RegisterResult    string
		RegisterReturnMsg string
	}

	username := r.FormValue("username-register")
	password := r.FormValue("password-register")
	pwConfirm := r.FormValue("passwordConfirm-register")
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
	username := r.FormValue("username-login")
	password := r.FormValue("password-login")
	log.Println(username, ":", password)

	if CheckUserPassword(username, password) {
		SetCookie(w, username, password)
		AddOnlineFriend(username, -1, -1)//-1 means that do not insert GPS data
		url := "/onlinefriends"
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
