package models

import (
	"utils"
)

type Users struct {
	Uid      int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	*Dao
}

func (this *Users) Login(name, passwd string) {
	fmt.Println("test")
}
