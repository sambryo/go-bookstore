package models

import (
	"github.com/jinzhu/gorm"
	"github.com/sambryo/go-bookstore/pkg/config"
)

var db *gorm.DB

type (
	Book struct {
		gorm.model
		Name         string `gorm:""json:"name"`
		Author       string `json:"author"`
		Publications string `json:"publications"`
	}
)

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}
