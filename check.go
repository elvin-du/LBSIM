package main

import (
	"log"
	"net/http"
	"database/sql"
	_"github.com/Go-SQL-Driver/MySQL"
)

func CheckUserPassword(user string, password string) bool {
	db, e := sql.Open("mysql","root:dumx@tcp(localhost:3306)/LBSIM?charset=utf8")
	if nil != e{
			log.Print(e)
			return false
	}
	defer db.Close()

	querySql := "select name, password from LBSIM.users where name = '" + user + "'"
	rows, e := db.Query(querySql)
	if nil != e{
			log.Print(e)
			return false
	}

	if rows.Next(){
			var usr, pw string
			rows.Scan(&usr, &pw)
			if usr == user && pw == password{
					return true
			}
	}

	return false
}

func SetCookie(w http.ResponseWriter, user, pw string) {
	cookie := http.Cookie{Name: appName, Value: user}
	http.SetCookie(w, &cookie)
}

func CheckCookie(r *http.Request) (name ,pw string, err error){
	cookie, err := r.Cookie(appName)
	if nil == err{
			name = cookie.Value
			pw = ""
			//if "" == name{
			//j		err = errors.New("the value of cookie is empty")
		//	}
	}

	return
}
