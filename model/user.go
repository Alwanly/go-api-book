package model

// User model
type User struct {
	ID        int    `gorm:"primaryKey;column:id;type:int(11);not null" `
	Name      string `gorm:"column:name;type:varchar(255);not null" `
	Email     string `gorm:"column:email;type:varchar(255);not null" `
	Password  string `gorm:"column:password;type:varchar(255);not null" `
	CreatedAt string `gorm:"column:created_at;type:timestamp;not null" `
}

// TableName for User model
func (User) TableName() string {
	return "users"
}

// Users model
type Users []User
