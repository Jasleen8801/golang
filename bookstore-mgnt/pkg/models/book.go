package models

import (
  "gorm.io/gorm"
  "github.com/Jasleen8801/golang/bookstore-mgnt/pkg/config"
)

var db *gorm.DB

type Book struct {
  gorm.model 
  Name string `gorm:""json:"name"`
  Author string `json:"author"`
  Publication string `json:"publication"`
}

func init() {
  config.Connect()
  db = config.GetDB()
  db.AutoMigrate(&Book{})
}
