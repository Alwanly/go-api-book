package book

import "gorm.io/gorm"

type Books struct {
	gorm.Model
	Title    string `gorm:"size:255;not null"`
	Author   string `gorm:"size:255;not null"`
	Headline string `gorm:"size:50;not null"`
	Tag      string `gorm:"size:255;not null"`
}
