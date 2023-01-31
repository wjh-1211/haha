package models

import "github.com/jinzhu/gorm"

type UserBasic struct{
	gorm.Model
	Name string
	Password string
	Phone string
	Email string
	identity string
	ClentIP string
}