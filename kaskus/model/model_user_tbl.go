package model

import "github.com/jinzhu/gorm"

// User use in table user
type User struct {
	gorm.Model
	Name     string `gorm:"not null" json:"name"`
	Image    string `gorm:"not null" json:"image"`
	Email    string `gorm:"not null" json:"email"`
	Role     string `gorm:"type:enum('user', 'admin'); default:'user'" json:"role"`
	Password string `gorm:"not null" json:"password"`
}

// TableName Rename table user in database
func (e *User) TableName() string {
	return "user"
}
