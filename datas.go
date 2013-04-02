package main

import (
	"code.google.com/p/go.net/websocket"
	"errors"
	"database/sql"
	_"github.com/Go-SQL-Driver/MySQL"
	"log"
)

type Location struct {
	Latitude  float64
	Longitude float64
}

type OnlineUser struct {
	Name   string
	Loc    *Location
	wsChatConn *websocket.Conn
	wsOnlineUserConn *websocket.Conn
}

type AllOnlineUser struct {
	AllUser []*OnlineUser
}

var allOnlineUser AllOnlineUser

func FindLocByName(name string) *Location {
	size := len(allOnlineUser.AllUser)

	for i := 0; i < size; i++ {
		if allOnlineUser.AllUser[i].Name == name {
			return allOnlineUser.AllUser[i].Loc
		}
	}
	return nil
}

func InsertWsChatConnData(name string, ws *websocket.Conn) {
	size := len(allOnlineUser.AllUser)

	for i := 0; i < size; i++ {
		if allOnlineUser.AllUser[i].Name == name {
			allOnlineUser.AllUser[i].wsChatConn = ws
			onlineUsersRefresh <- true
			return
		}
	}
}

func InsertWsOnlineUserConnData(name string, ws *websocket.Conn) {
	size := len(allOnlineUser.AllUser)

	for i := 0; i < size; i++ {
		if allOnlineUser.AllUser[i].Name == name {
			allOnlineUser.AllUser[i].wsOnlineUserConn= ws
			onlineUsersRefresh <- true
			return
		}
	}
}

func UpdateOnlineUsers(){
		<-onlineUsersRefresh
		size := len(allOnlineUser.AllUser)
		for i:=0; i<size; i++{
				ws := allOnlineUser.AllUser[i].wsOnlineUserConn
				if nil == ws{
						continue
				}
				if err := websocket.Message.Send(ws, "Y"); err != nil {
						log.Println(err)
						continue
				}
		}
}

func GetConnByName(name string) *websocket.Conn {
	size := len(allOnlineUser.AllUser)

	for i := 0; i < size; i++ {
		if allOnlineUser.AllUser[i].Name == name {
			return allOnlineUser.AllUser[i].wsChatConn
		}
	}

	return nil
}

func AddOnlineUser(username string, pw string, lat float64, lng float64){
		loc := Location{Latitude: lat, Longitude: lng}
		onlineUser := OnlineUser{Name: username, Loc: &loc}
		allOnlineUser.AllUser = append(allOnlineUser.AllUser, &onlineUser)
		onlineUsersRefresh <- true
}

func AddUser(username string, password string, pwConfirm string)error{
		db, e := sql.Open("mysql", "root:dumx@tcp(localhost:3306)/LBSIM?charset=utf8")
		if nil != e{
				log.Print(e)
				return e
		}
		defer db.Close()
		querySql := "select name from LBSIM.users WHERE name = ' " + username + "'"
		rows, e := db.Query(querySql)
		if nil != e{
				log.Print(e)
				return e
		}
		if rows.Next(){
				return errors.New("user exsited")
		}

		insertSql := "INSERT LBSIM.users SET name=?, password=?"
		stmt, e := db.Prepare(insertSql)
		if nil != e{
				log.Print(e)
				return e
		}
		defer stmt.Close()

		_, e = stmt.Exec(username, password)
		if nil != e{
				log.Print(e)
				return e
		}

		return nil
}
