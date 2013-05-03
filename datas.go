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

type OnlineFriend struct {
	Name   string
	Loc    *Location
	wsChatConn *websocket.Conn
	wsOnlineFriendConn *websocket.Conn
	wsRoutConn *websocket.Conn
}

type AllOnlineFriend struct {
	AllUser []*OnlineFriend
}

var allOnlineFriend AllOnlineFriend

func FindLocByName(name string) *Location {
	size := len(allOnlineFriend.AllUser)

	for i := 0; i < size; i++ {
		if allOnlineFriend.AllUser[i].Name == name {
			return allOnlineFriend.AllUser[i].Loc
		}
	}
	return nil
}

func InsertWsChatConnData(name string, ws *websocket.Conn) {
	size := len(allOnlineFriend.AllUser)

	for i := 0; i < size; i++ {
		if allOnlineFriend.AllUser[i].Name == name {
			allOnlineFriend.AllUser[i].wsChatConn = ws
			return
		}
	}
}

func InsertWsOnlineFriendConnData(name string, ws *websocket.Conn) {
	log.Println("InsertWsOnlineFriendConnData()")
	size := len(allOnlineFriend.AllUser)

	for i := 0; i < size; i++ {
		if allOnlineFriend.AllUser[i].Name == name {
			if nil == allOnlineFriend.AllUser[i].wsOnlineFriendConn{
				log.Println("onlineUsersRefresh<-name")
				onlineUsersRefresh <- name
			}
			allOnlineFriend.AllUser[i].wsOnlineFriendConn= ws
			return
		}
	}
}

func observeOnlineFriends(){
	for{
		name :=<-onlineUsersRefresh
		log.Println("observeOnlineFriends", name)
		size := len(allOnlineFriend.AllUser)
		for i:=0; i<size; i++{
			ws := allOnlineFriend.AllUser[i].wsOnlineFriendConn
			if allOnlineFriend.AllUser[i].Name == name{
				continue
			}
			if nil == ws{
				continue
			}
			if err := websocket.Message.Send(ws, "R"); err != nil {
				log.Println(err)
				continue
			}
		}
	}
}

func GetConnByName(name string) *websocket.Conn {
	size := len(allOnlineFriend.AllUser)

	for i := 0; i < size; i++ {
		if allOnlineFriend.AllUser[i].Name == name {
			return allOnlineFriend.AllUser[i].wsChatConn
		}
	}

	return nil
}

func AddOnlineFriend(username string, lat float64, lng float64){
	for i := 0; i <len(allOnlineFriend.AllUser); i++{
		if allOnlineFriend.AllUser[i].Name == username{
			if ( lat > 0) && (lng > 0){
				allOnlineFriend.AllUser[i].Loc.Latitude = lat
				allOnlineFriend.AllUser[i].Loc.Longitude = lng
			}
			return
		}
	}

	loc := Location{Latitude: lat, Longitude: lng}
	onlineUser := OnlineFriend{Name: username, Loc: &loc}
	allOnlineFriend.AllUser = append(allOnlineFriend.AllUser, &onlineUser)
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
