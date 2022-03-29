package model

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name      string `gorm:"type:varchar(20);not null"`
	Telephone string `gorm:"varchar(110;not null;unique"`
	Password  string `gorm:"size:255;not null"`
}

type GinFile struct {
	ID         uuid.UUID `json:"id" gorm:"type:char(36);primary_key"`
	UserId     uint      `json:"user_id" gorm:"not null"`
	CategoryId uint      `json:"category_id" gorm:"not null"`
	Category   *Category
	Title      string `json:"title" gorm:"type:varchar(50);not null"`
	HeadImg    string `json:"head_img"`
	Content    string `json:"content" gorm:"tyep:text;not null"`
	ModelTime
}

type Category struct {
	ID   uint   `json:"id" gorm:"primary_key"`
	Name string `json:"name" gorm:"type:varchar(50);not null;unique"`
	ModelTime
}

type ModelTime struct {
	CreatedAt Time `json:"created_at" gorm:"type:timestamp"`
	UpdatedAt Time `json:"updated_at" gorm:"type:timestamp"`
}
