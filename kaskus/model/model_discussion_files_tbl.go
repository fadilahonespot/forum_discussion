package model

import "github.com/jinzhu/gorm"

type DiscussionImages struct {
	gorm.Model
	DiscussionID uint   `gorm:"not null" json:"discussion_id"`
	UserID       uint   `gorm:"not null" json:"user_id"`
	Image        string `gorm:"not null" json:"image"`
}

type DiscussionFiles struct {
	gorm.Model
	DiscussionID uint   `gorm:"not null" json:"discussion_id"`
	UserID       uint   `gorm:"not null" json:"user_id"`
	File         string `gorm:"not null" json:"file"`
}

func (e *DiscussionImages) TableName() string {
	return "discussion_images"
}

func (e *DiscussionFiles) TableName() string {
	return "discussion_files"
}
