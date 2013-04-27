package main

import (
	"errors"
	"runtime"
	"strconv"
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

func ParseLngLat(rcvMsg string)(lng float64,lat float64,err error){
	index := strings.Index(rcvMsg, ":")
	if -1 == index {
		return -1, -1, errors.New("Lng lat can not be found:")
	}
	lng,_= strconv.ParseFloat(rcvMsg[:index],64)
	lat,_ = strconv.ParseFloat(rcvMsg[index+1:],64)
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
