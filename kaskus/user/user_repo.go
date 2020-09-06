package user

import "kaskus/model"

// UserRepo Is interface 
type UserRepo interface {
	ViewAll() (*[]model.User, error)
	AddUser(user *model.User) (*model.User, error)
	ViewByEmail(email string) (*model.User, error)
	ViewById(id int) (*model.User, error)
}