package model

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	Name   string
	IsDone bool
}
