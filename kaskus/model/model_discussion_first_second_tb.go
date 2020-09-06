package model

import "github.com/jinzhu/gorm"

type DiscussionFirst struct {
	gorm.Model
	UserID       uint   `gorm:"not null" json:"user_id"`
	DiscussionID uint   `gorm:"not null" json:"discussion_id"`
	Message      string `gorm:"type:text; not null" json:"message"`
	Date         string `gorm:"type:date; not null" json:"date"`
	File         string `json:"file"`
	Image        string `json:"image"`
}

type DiscussionSecond struct {
	gorm.Model
	UserID            uint   `gorm:"not null" json:"user_id"`
	DiscussionFirstID uint   `gorm:"not null" json:"discussion_first_id"`
	Message           string `gorm:"type:text; not null" json:"message"`
	Date              string `gorm:"type:date; not null" json:"date"`
	File              string `json:"file"`
	Image             string `json:"image"`
}

func (e *DiscussionFirst) TableName() string {
	return "discussion_first"
}

func (e *DiscussionSecond) TableName() string {
	return "discussion_second"
}
