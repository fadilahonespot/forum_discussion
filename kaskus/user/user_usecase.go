package user

import (
	"kaskus/model"
	"github.com/gin-gonic/gin"
)


type UserUsecase interface {
	ViewAll() (*[]model.User, error)
	AddUser(user *model.User) (*model.User, error)
	ViewByEmail(email string) (*model.User, error)
	ViewById(id int) (*model.User, error)
	ValidUser(c *gin.Context) (*model.User, error)
	AdminOnly(c *gin.Context) (*model.User, error)
}