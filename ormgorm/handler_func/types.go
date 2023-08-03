package handler_func

import (
	"gorm.io/gorm"
)

type App struct {
	Db *gorm.DB
}

type Person struct {
	Name       string `json:"name"`
	Profession string `json:"profession"`
	Contact    string `json:"contact"`
	Email      string `json:"email"`
	Experience int    `json:"experience"`
}
