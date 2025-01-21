package model

import "time"

// book model

type Book struct {
	ID        string    `gorm:"primaryKey;column:id;type:varchar(255);not null" `
	Title     string    `gorm:"column:title;type:varchar(255);not null" `
	Author    string    `gorm:"column:author;type:varchar(255);not null" `
	CreatedAt time.Time `gorm:"column:created_at;type:timestamptz;not null" `
	CreatedBy string    `gorm:"column:created_by;type:varchar(255);not null" `
	UpdatedAt time.Time `gorm:"column:updated_at;type:timestamptz;not null"`
	UpdatedBy string    `gorm:"column:updated_by;type:varchar(255);not null" `
}

// TableName for Book model
func (Book) TableName() string {
	return "books"
}

// Books model
type Books []Book
