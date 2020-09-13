package discussion

import (
	"kaskus/model"
	"github.com/gin-gonic/gin"
)


type DiscussionUsecase interface {
	AddCatagory(catagory *model.Catagory) (*model.Catagory, error)
	DeleteCatagory(id int) error
	ViewCatagoryById(id int)(*model.Catagory, error)
	UpdateCatagoryById(id int, catagory *model.Catagory) (*model.Catagory, error)
	ViewAllCatagory()(*[]model.Catagory, error)
	ViewAllDiscussion() (*[]model.Discussion, error)
	AddDiscussion(c *gin.Context, discussion *model.Discussion)(*model.DiscussionPost, error)
	AddDiscussionFisrt(c *gin.Context, DiscussionFirst *model.DiscussionFirst)(*model.DiscussionFirst, error)
	AddDiscussionSecond(c *gin.Context, discussionSecond *model.DiscussionSecond)(*model.DiscussionSecond, error)
	ViewDiscussionById(id int) (*model.Discussion, error)
	ViewDiscussionFirstById(id int) (*model.DiscussionFirst, error)
	ViewDiscussionFirstByDiscussionId(id int) (*[]model.DiscussionFirst, error)
	ViewDiscussionSecondByDiscussionFirstId(id int) (*[]model.DiscussionSecond, error)
	ViewDiscussionDetailByID(id int) (*model.DiscussionDetailShow, error)
	DeleteDiscussionByID(id int) error
	UpdateDiscussionById(id int, c *gin.Context, discussion *model.Discussion) (*model.DiscussionPost, error)
}