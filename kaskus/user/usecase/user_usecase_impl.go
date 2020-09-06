package usecase

import (
	"fmt"
	"kaskus/middleware"
	"kaskus/model"
	"kaskus/user"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)


type UserUsecase struct {
	userRepo user.UserRepo
}

func CreateUserUsecase(userRepo user.UserRepo) user.UserUsecase {
	return &UserUsecase{userRepo}
}

func (e *UserUsecase)ViewAll() (*[]model.User, error) {
	return e.userRepo.ViewAll()
}

func (e *UserUsecase) ViewByEmail(email string) (*model.User, error) {
	return e.userRepo.ViewByEmail(email)
}

func (e *UserUsecase) ViewById(id int) (*model.User, error) {
	return e.userRepo.ViewById(id)
}

func (e *UserUsecase)AddUser(user *model.User) (*model.User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		fmt.Printf("[UserUsecaseimpl.AddUser] error to generate password %v \n", err)
		return nil, fmt.Errorf("Oopss server someting wrong")
	}
	user.Password = string(hash)
	return e.userRepo.AddUser(user)
}

func (e *UserUsecase) ValidUser(c *gin.Context) (*model.User, error) {
	userAuth, err := middleware.ExtractTokenAuth(c)
	if err != nil {
		return nil, err
	}
	user, err := e.userRepo.ViewById(userAuth.ID)
	if err != nil {
		return nil, err
	}
	if user.Password != userAuth.Password {
		return nil, fmt.Errorf("invalid token, please sign out and sign in again")
	}
	return user, nil
}

func (e *UserUsecase) AdminOnly(c *gin.Context) (*model.User, error) {
	user, err := e.ValidUser(c)
	if err != nil {
		return nil, err
	}
	if user.Role != "admin" {
		return nil, fmt.Errorf("You are not an admin")
	}
	return user, nil
}