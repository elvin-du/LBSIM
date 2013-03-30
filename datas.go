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
	wsConn *websocket.Conn
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

func InsertConnData(name string, ws *websocket.Conn) {
	size := len(allOnlineUser.AllUser)

	for i := 0; i < size; i++ {
		if allOnlineUser.AllUser[i].Name == name {
			allOnlineUser.AllUser[i].wsConn = ws
			return
		}
	}
}

func GetConnByName(name string) *websocket.Conn {
	size := len(allOnlineUser.AllUser)

	for i := 0; i < size; i++ {
		if allOnlineUser.AllUser[i].Name == name {
			return allOnlineUser.AllUser[i].wsConn
		}
	}

	return nil
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
