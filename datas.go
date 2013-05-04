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
	Page	string//determine parameter wsConn to connected which page 
	wsConn *websocket.Conn
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

func UpdateWsConn(name,page string, ws *websocket.Conn){
	log.Println("UpdateWsConn()")
	size := len(allOnlineFriend.AllUser)
	for i := 0; i < size; i++ {
		//if do not match,jump to next
		if allOnlineFriend.AllUser[i].Name != name {
			continue
		}

		switch page{
		case "route","chat":
		case "online":
			if nil == allOnlineFriend.AllUser[i].wsConn{
				log.Println("onlineUsersRefresh<-name")
				onlineUsersRefresh <- name
			}
		}
		allOnlineFriend.AllUser[i].Page = page
		allOnlineFriend.AllUser[i].wsConn= ws
		return
	}
}

func observeOnlineFriends(){
	for{
		name :=<-onlineUsersRefresh
		log.Println("observeOnlineFriends", name)
		size := len(allOnlineFriend.AllUser)
		for i:=0; i<size; i++{
			ws := allOnlineFriend.AllUser[i].wsConn
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
			return allOnlineFriend.AllUser[i].wsConn
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
