package users

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	Email    string  `gorm:"size:255;not null; uniqueIndex:idx_email"`
	Password *string `gorm:"size:100;not null"`
	Name     string  `gorm:"size:255;not null"`
}
