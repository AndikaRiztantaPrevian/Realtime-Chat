package models

import "gorm.io/gorm"

type Message struct {
	ID      uint   `json:"id" gorm:"primary_key"`
	Message string `json:"message"`
}

func Migrate(db *gorm.DB) {
    db.AutoMigrate(&Message{})
}