package main

import (
	"errors"
	"runtime"
	"log"
	"strings"
)

func ParseRcvMsg(rcvMsg string) (name string, content string, err error) {
	index := strings.Index(rcvMsg, ":")
	if -1 == index {
		return "", "", errors.New("can not find :")
	}
	name = rcvMsg[:index]
	content = rcvMsg[index:]
	return
}

func CheckError(err error) {
	if nil != err {
		_, file, line, ok := runtime.Caller(1)
		if ok {
			log.Print(file, ":", line)
		}
		log.Println(err)
	}
}
