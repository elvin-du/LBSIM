package main

import (
	"errors"
	"strings"
)

func ParseRcvMsg(rcvMsg string)(name string, content string, err error){
		index := strings.Index(rcvMsg, ":")
		if -1 == index{
				return "","",errors.New("can not find :")
		}
		name = rcvMsg[:index]
		content = rcvMsg[index:]
		return
}

