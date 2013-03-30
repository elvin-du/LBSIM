package main

import (
	"code.google.com/p/go.net/websocket"
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

func AddUser(username string, password string, pwConfirm string)(ret string, err error){
		ret = "congratulation! register successful"
		return
}
