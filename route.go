package main

import (
	"code.google.com/p/go.net/websocket"
	"fmt"
	"html/template"
	"log"
	"os"
	"net/http"
	"strconv"
)

type Location struct{
		Latitude float64
		Longitude float64
}

type  OnlineUser struct{
		Name string
		Loc  *Location
		wsConn *websocket.Conn
}

type AllOnlineUser struct{
		AllUser []*OnlineUser
}

var allOnlineUser  AllOnlineUser

func FindLocByName(name string) *Location{
	size := len(allOnlineUser.AllUser)

	for i := 0; i < size; i++ {
			if allOnlineUser.AllUser[i].Name == name{
					return allOnlineUser.AllUser[i].Loc
			}
	}

	return nil
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		http.Redirect(w, r, "/login", http.StatusFound)
	}

	t, err := template.ParseFiles("templates/html/error.html")
	if err != nil {
		log.Println(err)
	}
	t.Execute(w, nil)
}

func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Login")
	r.ParseForm()
	if r.Method == "POST" {
		username := r.FormValue("username")
		password := r.FormValue("password")
		lat,_ := strconv.ParseFloat(r.FormValue("latitude"), 64)
		lng,_ := strconv.ParseFloat(r.FormValue("longitude"), 64)
		//fmt.Println(lat)
		//fmt.Println(lng)

		if CheckUserPassword(username, password) {
			SetCookie(w, username, password)

			loc := Location{Latitude:lat, Longitude:lng}
			onlineUser := OnlineUser{Name: username, Loc:&loc}
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

func OnlineUsers(w http.ResponseWriter, r *http.Request){
		fmt.Println("OnlineUsers")
		r.ParseForm()

		if r.Method == "POST"{
				withWho := r.FormValue("onlineUser")
				var urlWithWho string
				chatOrRoute := r.FormValue("chatOrRoute")
				if "chat" == chatOrRoute{
					urlWithWho = "/chat?withWho=" + withWho;
				}else if "route" == chatOrRoute{
					urlWithWho = "/route?withWho=" + withWho;
				}

				http.Redirect(w, r, urlWithWho, http.StatusFound)
				return
		}

		t,err := template.ParseFiles("templates/html/onlineUser.html")
		checkError(err)

		err = t.Execute(w, allOnlineUser)
		checkError(err)
}

func Chat(w http.ResponseWriter, r *http.Request){
		fmt.Println("Chat")
		r.ParseForm()

		withWho := r.Form.Get("withWho")
		type ToWho struct{
				Name string
		}

		toWho := ToWho{Name: withWho}

		t, err := template.ParseFiles("templates/html/chat.html")
		checkError(err)
		err = t.Execute(w , toWho)
		checkError(err)
}

func WebsocketChat(ws *websocket.Conn){
		var err error
		fmt.Println("WebsocketChat")

		for{
				var reply string

				if err = websocket.Message.Receive(ws, &reply); err != nil {
						fmt.Println("Can't receive")
						fmt.Println(err)
						break
				}

				fmt.Println("Received back from client: " + reply)

				msg := "welcome to websocket do by pp"
				fmt.Println("Sending to client: " + msg)

				if err = websocket.Message.Send(ws, msg); err != nil {
						fmt.Println("Can't send")
						break
				}
		}
}

func Route(w http.ResponseWriter, r *http.Request){
		fmt.Println("Route")
		r.ParseForm()

		withWho := r.Form.Get("withWho")
		if "" == withWho{
				return
		}
		loc := FindLocByName(withWho)

		t, err := template.ParseFiles("templates/html/route.html")
		checkError(err)

		//fmt.Println(loc.Longitude, loc.Latitude)
		end := Location{Longitude: loc.Longitude, Latitude:loc.Latitude}
		err = t.Execute(w, end)
		checkError(err)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
