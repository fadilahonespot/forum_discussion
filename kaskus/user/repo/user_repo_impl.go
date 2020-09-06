package repo

import (
	"fmt"
	"kaskus/model"
	"kaskus/user"

	"github.com/jinzhu/gorm"
)

// UserRepoImpl struct
type UserRepoImpl struct {
	DB *gorm.DB
}

//CreateUserRepo is constructor
func CreateUserRepo(DB *gorm.DB) user.UserRepo {
	return &UserRepoImpl{DB}
}

// ViewAll function view all user
func (e *UserRepoImpl)ViewAll() (*[]model.User, error) {
	var user []model.User
	err := e.DB.Find(&user).Error
	if err != nil {
		fmt.Printf("[UserRepoImpl.ViewAll] error execute query %v", err)
		return nil, fmt.Errorf("Oppss server someting wrong")
	}
	return &user, nil
}

// ViewByEmail function
func (e *UserRepoImpl) ViewByEmail(email string) (*model.User, error) {
	var user = model.User{}
	err := e.DB.Table("user").Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, fmt.Errorf("Email user is not found")
	}
	return &user, nil
}

func (e *UserRepoImpl) ViewById(id int) (*model.User, error) {
	var user = model.User{}
	err := e.DB.Table("user").Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, fmt.Errorf("id user is not found")
	}
	return &user, nil
}


// AddUser function add user
func (e *UserRepoImpl) AddUser(user *model.User) (*model.User, error) {
	err := e.DB.Save(&user).Error
	if err != nil {
		fmt.Printf("[UserRepoImpl.AddUser] error execute query %v", err)
		return nil, fmt.Errorf("failed insert data user")
	}
	return user, nil
}