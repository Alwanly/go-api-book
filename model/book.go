package model

// book model

type Book struct {
	ID    int    `gorm:"primaryKey;column:id;type:int(11);not null" `
	Title string `gorm:"column:title;type:varchar(255);not null" `
}

// TableName for Book model
func (Book) TableName() string {
	return "books"
}

// Books model
type Books []Book
