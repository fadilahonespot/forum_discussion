package discussion

import (
	"kaskus/model"
	"github.com/jinzhu/gorm"
)


type DiscussionRepo interface {
	BeginTrans() *gorm.DB
	ViewAllDiscussion() (*[]model.Discussion, error)
	AddCatagory(catagory *model.Catagory) (*model.Catagory, error)
	DeleteCatagory(id int) error
	ViewAllCatagory()(*[]model.Catagory, error)
	ViewCatagoryById(id int)(*model.Catagory, error)
	AddDiscussion(discussion *model.Discussion, tx *gorm.DB) (*model.Discussion, error)
	AddDiscussionImages(discussionImages *model.DiscussionImages, tx *gorm.DB)(*model.DiscussionImages, error)
	AddDiscussionFiles(discussionFiles *model.DiscussionFiles, tx *gorm.DB)(*model.DiscussionFiles, error) 
	AddDiscussionFisrt(discussionFirst *model.DiscussionFirst)(*model.DiscussionFirst, error)
	AddDiscussionSecond(discussionSecond *model.DiscussionSecond)(*model.DiscussionSecond, error)
	ViewDiscussionById(id int) (*model.Discussion, error)
	ViewDiscussionFirstById(id int) (*model.DiscussionFirst, error)
	ViewDiscussionFirstByDiscussionId(id int) (*[]model.DiscussionFirst, error)
	ViewDiscussionSecondByDiscussionFirstId(id int) (*[]model.DiscussionSecond, error)
	ViewDiscussionImageByDiscussionID(id int) (*[]model.DiscussionImages, error)
	ViewDiscussionFileByDiscussionID(id int) (*[]model.DiscussionFiles, error)
	DeleteDiscussionById(id int, tx *gorm.DB) error
	DeleteDiscussionFirstByID(id int, tx *gorm.DB) error
	DeleteDiscussionSecondByID(id int, tx *gorm.DB) error 
	ViewImagesByDiscussionID(id int) (*[]model.DiscussionImages, error)
	ViewFilesByDiscussionID(id int) (*[]model.DiscussionFiles, error)
	DeleteDiscussionFilesByDiscussionID(id int, tx *gorm.DB) error
	DeleteDiscussionImagesByDiscussionID(id int, tx *gorm.DB) error
	UpdateDiscussionById(id int, discussion *model.Discussion, tx *gorm.DB) (*model.Discussion, error)
}