package main

import (
	"log"
	"net/http"
	"database/sql"
	"errors"
	_"github.com/Go-SQL-Driver/MySQL"
	"encoding/base64"
	"strings"
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
	encryptedUsrpw := string(encrypt(user + ":" + pw))
	cookie := http.Cookie{Name: appName, Value: encryptedUsrpw}
	http.SetCookie(w, &cookie)
}

func base64Encode(src []byte) []byte {
    return []byte(base64.StdEncoding.EncodeToString(src))
}

func base64Decode(src []byte) ([]byte, error) {
    return base64.StdEncoding.DecodeString(string(src))
}

func encrypt(src string) []byte{
	return base64Encode([]byte(src))
}

func decrypt(src string) ([]byte,error){
	return base64Decode([]byte(src))
}

func CheckCookie(r *http.Request) (string,string,error){
	var name, pw string
	cookie, err := r.Cookie(appName)
	if nil == err{
		val,err := decrypt(cookie.Value)
		namepw := string(val)
		if "" == namepw{
			err = errors.New("the value of cookie is empty")
			return "","", err
		}
		index := strings.Index(namepw, ":")
		name = namepw[:index]
		pw = namepw[index+1:]
	}

	if(!CheckUserPassword(name,pw)){
		return "","",errors.New("wrong password or username")
	}

	return name, pw, err
}
