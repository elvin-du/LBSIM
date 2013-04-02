package main

import (
	"code.google.com/p/go.net/websocket"
	"fmt"
	"time"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("NotFoundHandler")
	r.ParseForm()
	if r.URL.Path == "/" {
		http.Redirect(w, r, "/login", http.StatusFound)
	}

	t, err := template.ParseFiles("templates/html/error.html")
	if err != nil {
		log.Println(err)
	}
	t.Execute(w, nil)
}

func Register(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Register")
	r.ParseForm()
	var regRet interface{} = nil
	type reg struct {
		RegisterResult    string
		RegisterReturnMsg string
	}

	if r.Method == "POST" {
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
	}

	t, err := template.ParseFiles("templates/html/register.html")
	checkError(err)
	err = t.Execute(w, regRet)
	checkError(err)
}

func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Login")
	r.ParseForm()
	if ret := CheckCookie(r); ret != ""{
			AddOnlineUser(ret, "", 0, 0)
			url := "/onlineUsers?who=" + ret
			http.Redirect(w,r, url, http.StatusFound)
			return
	}else{
			log.Println(ret)
	}
	var data interface{}
	if r.Method == "POST" {
		username := r.FormValue("username")
		password := r.FormValue("password")
		lat, _ := strconv.ParseFloat(r.FormValue("latitude"), 64)
		lng, _ := strconv.ParseFloat(r.FormValue("longitude"), 64)
		//fmt.Println(lat)
		//fmt.Println(lng)

		if CheckUserPassword(username, password) {
			SetCookie(w, username)
			AddOnlineUser(username, password, lat, lng)
			url := "/onlineUsers?who=" + username
			http.Redirect(w, r, url, http.StatusFound)
		} else {
			SetCookie(w, "")
			type loginRet struct{
					LoginRet string
			}
			data = loginRet{"wrongUsrPw"}
			//fmt.Println("suername or password is wrong, please input again!")
		}
	}

	t, _ := template.ParseFiles("templates/html/login.html")
	t.Execute(w, data)
}

func OnlineUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("OnlineUsers")
	r.ParseForm()
	who := r.Form.Get("who")
	if "" == who{
			http.Redirect(w,r,"/login", http.StatusFound)
			return //thie sentence is important, following line will exexute when no this line
	}
	if r.Method == "POST" {
		withWho := r.FormValue("onlineUser")
		var urlWithWho string
		chatOrRoute := r.FormValue("chatOrRoute")
		if "chat" == chatOrRoute {
			urlWithWho = "/chat?withWho=" + withWho
		} else if "route" == chatOrRoute {
			urlWithWho = "/route?withWho=" + withWho
		}

		http.Redirect(w, r, urlWithWho, http.StatusFound)
		return
	}

	t, err := template.ParseFiles("templates/html/onlineUser.html")
	checkError(err)

	err = t.Execute(w, allOnlineUser)
	checkError(err)
}

func Chat(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Chat")
	r.ParseForm()
	withWho := r.Form.Get("withWho")

	type ToWho struct {
		Name string
	}

	if withWho == ""{
			url := "/login"
			http.Redirect(w,r, url, http.StatusFound)
			return
	}

	toWho := ToWho{Name: withWho}

	t, err := template.ParseFiles("templates/html/chat.html")
	checkError(err)
	err = t.Execute(w, toWho)
	checkError(err)
}

func WsChat(ws *websocket.Conn) {
	var err error
	var toWho *websocket.Conn
	var rcvMsg string
	fmt.Println("WebsocketChat")
	request := ws.Request()
	cookie, err := request.Cookie(appName)
	name := cookie.Value
	InsertWsChatConnData(name, ws)

	for {
		if err = websocket.Message.Receive(ws, &rcvMsg); err != nil {
			fmt.Println("Can't receive")
			fmt.Println(err)
			break
		}
		fmt.Println("Received : " + rcvMsg)

		name, content, err := ParseRcvMsg(rcvMsg)
		toWho = GetConnByName(name)

		if err = websocket.Message.Send(toWho, content); err != nil {
			fmt.Println("Can't send")
			fmt.Println(err)
			break
		}
	}
}

func WsOnlineUsers(ws *websocket.Conn){
	//	var err error
	//	var rcvMsg string
		request := ws.Request()
		cookie, _:= request.Cookie(appName)
		name := cookie.Value
		InsertWsOnlineUserConnData(name, ws)
		log.Println("WsOnlineUser")
		for{
				time.Sleep(5000)
				if err = websocket.Message.Receive(ws, &rcvMsg); err != nil{
					log.Println(err)
					break
			}
			//if err = websocket.Message.Send(ws, "Y"); err != nil {
				//		log.Println(err)
				//		break
				//}
		}
}

func Route(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Route")
	r.ParseForm()

	withWhom := r.Form.Get("withWho")
	if "" == withWhom{
			http.Redirect(w,r,"/login", http.StatusFound)
			return //thie sentence is important, following line will exexute when no this line
	}

	loc := FindLocByName(withWhom)

	t, err := template.ParseFiles("templates/html/route.html")
	checkError(err)

	//fmt.Println(loc.Longitude, loc.Latitude)
	end := Location{Longitude: loc.Longitude, Latitude: loc.Latitude}
	err = t.Execute(w, end)
	checkError(err)
}
