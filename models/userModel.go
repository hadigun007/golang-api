package models

import "gorm.io/gorm"

type UserModel struct {
	gorm.Model
	Name   string
	Gender string
	Agr    uint8
	Hobby  []string
}
