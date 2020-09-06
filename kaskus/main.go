package main

import (
	"log"

	"kaskus/config"
	userHandler "kaskus/user/handler"
	userRepo "kaskus/user/repo"
	userUsecase "kaskus/user/usecase"
	discussionHandler "kaskus/discussion/handler"
	discussionRepo "kaskus/discussion/repo"
	discussionUsecase "kaskus/discussion/usecase"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile("config/config.json")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}


func main() {
	db := config.DbConfig()
	defer db.Close()

	router := gin.Default()

	userRepo := userRepo.CreateUserRepo(db)
	userUsecase := userUsecase.CreateUserUsecase(userRepo)
	discussionRepo := discussionRepo.CreateDiscussionRepo(db)
	discussionUsecase := discussionUsecase.CreateDiscussionUsecase(discussionRepo, userRepo)

	userHandler.CreateUserHandler(router, userUsecase)
	discussionHandler.CreateDiscussionHandler(router, discussionUsecase, userUsecase)

	err := router.Run(":" + viper.GetString("port.router"))
	if err != nil {
		log.Fatal(err)
	}
}