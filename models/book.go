package models

import (
	"github.com/jinzhu/gorm"
)

type Book struct {
	gorm.Model
	// ID     uint   `json:"id" gorm:"primary_key"`
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
}

type UpdateBookInput struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

// func (b *Book) BeforeCreate(tx *gorm.DB) (err error) {
// 	b.ID = b.ID + 1

// 	return
// }

// func (b *Book) BeforeCreate(tx *gorm.DB) (err error) {

// 	u.UUID = uuid.New()

// 	if u.Role == "admin" {
// 		return errors.New("invalid role")
// 	}
// 	return
// }
