package config

import (
	"kaskus/model"
	"log"

	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)


func DbConfig() *gorm.DB {
	DB, err := gorm.Open("mysql", viper.GetString("database.mysql"))
	if err != nil {
		log.Fatal(err)
	}

	DB.Debug().AutoMigrate(
		model.User{},
		model.Discussion{},
		model.Catagory{},
		model.DiscussionImages{},
		model.DiscussionFiles{},
		model.DiscussionFirst{},
		model.DiscussionSecond{},
	)

	DB.Model(&model.Discussion{}).AddForeignKey("user_id", "user(id)", "CASCADE", "CASCADE")
	DB.Model(&model.Catagory{}).AddForeignKey("user_id", "user(id)", "CASCADE", "CASCADE")
	DB.Model(&model.Discussion{}).AddForeignKey("catagory_id", "catagory(id)", "CASCADE", "CASCADE")
	DB.Model(&model.DiscussionImages{}).AddForeignKey("user_id", "user(id)", "CASCADE", "CASCADE")
	DB.Model(&model.DiscussionImages{}).AddForeignKey("discussion_id", "discussion(id)", "CASCADE", "CASCADE")
	DB.Model(&model.DiscussionFiles{}).AddForeignKey("user_id", "user(id)", "CASCADE", "CASCADE")
	DB.Model(&model.DiscussionFiles{}).AddForeignKey("discussion_id", "discussion(id)", "CASCADE", "CASCADE")
	DB.Model(&model.DiscussionFirst{}).AddForeignKey("discussion_id", "discussion(id)", "CASCADE", "CASCADE")
	DB.Model(&model.DiscussionFirst{}).AddForeignKey("user_id", "user(id)", "CASCADE", "CASCADE")
	DB.Model(&model.DiscussionSecond{}).AddForeignKey("user_id", "user(id)", "CASCADE", "CASCADE")
	DB.Model(&model.DiscussionSecond{}).AddForeignKey("discussion_first_id", "discussion_first(id)", "CASCADE", "CASCADE")

	return DB
}