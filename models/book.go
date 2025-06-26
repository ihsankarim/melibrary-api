package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model

	Title       string   `json:"title"`
	Author      string   `json:"author"`
	Description string   `json:"description"`
	CategoryId  uint     `json:"category_id"`
	Category    Category `json:"category" gorm:"foreignKey:CategoryID"`
}
