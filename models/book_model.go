package models

import "gorm.io/gorm"

type BookInput struct {
	Content    string `json:"content" binding:"required"`
	Author     string `json:"author" binding:"required"`
	PageNumber int    `json:"page_number" binding:"required"`
}

type Book struct {
	gorm.Model
	Content    string `json:"content" gorm:"not null"`
	Author     string `json:"author"`
	PageNumber int    `json:"page_number"`
}
