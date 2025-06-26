package models

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model

	Title       string   `json:"title"`
	Author      string   `json:"author"`
	Description string   `json:"description"`
	CategoryID  uint     `json:"category_id"`
	Category    Category `json:"category" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}
