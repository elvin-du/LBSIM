package main

import (
	"fmt"
	"net/http"
)

func CheckUserPassword(user string, password string) bool {
		//if "dumx" == user && "dumx" == password{
		//		return true
		//}
	return true 
}

func SetCookie(w http.ResponseWriter, user string, password string) {
	cookie := http.Cookie{Name: user, Value: password}
	http.SetCookie(w, &cookie)
}

func CheckCookie(r *http.Request) bool {
	cookie, _ := r.Cookie("dumx")
	fmt.Println(cookie.Value)
	if "dumx" == cookie.Value{
			return true
	}

	return false
}
