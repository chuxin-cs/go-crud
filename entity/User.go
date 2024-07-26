package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Id        int    `json:"id"`
	Name      string `json:"name"`
	NickName  string `json:"NickName"`
	Avatar    string `json:"avatar"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	Mobile    string `json:"mobile"`
	DelStatus int    `json:"delStatus"`
}

func (User) TableName() string {
	return "sys_user"
}
