// Author：macs		dumanxiang@gmail.com

package main

import (
	. "config"
	"log"
	"logger"
	"net/http"
)

func init() {
	log.SetFlags(log.Llongfile | log.LstdFlags)
}

func main() {
	log.Println("Server Started, and listen on port: 8888")
	logger.Infoln("Server Started, listen on port: 8888")

	if err := http.ListenAndServe(Config["host"], http.HandlerFunc(router)); err != nil {
		log.Println("ListenAndServe:", err)
		logger.Errorln("ListenAndServe:", err)
		return
	}
}
