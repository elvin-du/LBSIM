package main

import (
	"code.google.com/p/go.net/websocket"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("NotFoundHandler")
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
		ret, err := AddUser(username, password, pwConfirm)
		if  nil != err {
			log.Println(err)
			regRet = reg{RegisterResult: "registerFailed", RegisterReturnMsg:ret}
		}
		log.Print("fdsafaf")
		log.Print(ret)
		regRet = reg{RegisterResult: "registerSuccessful", RegisterReturnMsg:ret}
	}

	t, err := template.ParseFiles("templates/html/register.html")
	checkError(err)
	err = t.Execute(w, regRet)
	checkError(err)
}

func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Login")
	r.ParseForm()
	if r.Method == "POST" {
		username := r.FormValue("username")
		password := r.FormValue("password")
		lat, _ := strconv.ParseFloat(r.FormValue("latitude"), 64)
		lng, _ := strconv.ParseFloat(r.FormValue("longitude"), 64)
		//fmt.Println(lat)
		//fmt.Println(lng)

		if CheckUserPassword(username, password) {
			SetCookie(w, "LBSIM", username)

			loc := Location{Latitude: lat, Longitude: lng}
			onlineUser := OnlineUser{Name: username, Loc: &loc}
			allOnlineUser.AllUser = append(allOnlineUser.AllUser, &onlineUser)

			http.Redirect(w, r, "/onlineUsers", http.StatusFound)
			return
		} else {
			fmt.Println("suername or password is wrong, please input again!")
		}
	}

	t, _ := template.ParseFiles("templates/html/login.html")
	t.Execute(w, nil)
}

func OnlineUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("OnlineUsers")
	r.ParseForm()

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

	toWho := ToWho{Name: withWho}

	t, err := template.ParseFiles("templates/html/chat.html")
	checkError(err)
	err = t.Execute(w, toWho)
	checkError(err)
}

func WebsocketChat(ws *websocket.Conn) {
	var err error
	var toWho *websocket.Conn
	var rcvMsg string
	fmt.Println("WebsocketChat")
	request := ws.Request()
	cookie, err := request.Cookie("LBSIM")
	name := cookie.Value
	InsertConnData(name, ws)

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

func Route(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Route")
	r.ParseForm()

	withWho := r.Form.Get("withWho")
	if "" == withWho {
		return
	}
	loc := FindLocByName(withWho)

	t, err := template.ParseFiles("templates/html/route.html")
	checkError(err)

	//fmt.Println(loc.Longitude, loc.Latitude)
	end := Location{Longitude: loc.Longitude, Latitude: loc.Latitude}
	err = t.Execute(w, end)
	checkError(err)
}
