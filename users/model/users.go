package model

import "github.com/jinzhu/gorm"

type User struct{
	gorm.Model
	Username string `json:"username" gorm:"type:varchar(255);unique_index"`
	Password string `json:"password" gorm:"size:255"`
	Nickname string `json:"nickname"`
}

func UserTable() *gorm.DB {
	return NewDB().Table("users")
}


