package handler

import (
	"fmt"
	"kaskus/discussion"
	"kaskus/middleware"
	"kaskus/model"
	"kaskus/user"
	"kaskus/utils"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type DiscussionHandler struct {
	discussionUsecase discussion.DiscussionUsecase
	userUsecase       user.UserUsecase
}

func CreateDiscussionHandler(r *gin.Engine, discussionUsecase discussion.DiscussionUsecase, UserUsecase user.UserUsecase) {
	discussionHandler := DiscussionHandler{discussionUsecase, UserUsecase}

	r2 := r.Group("/catagory").Use(middleware.TokenVerifikasiMiddleware())
	r2.POST("", discussionHandler.addCatagory)
	r2.GET("", discussionHandler.viewAllCatagory)
	r2.DELETE("/:id", discussionHandler.deleteCatagory)

	r3 := r.Group("/discussion").Use(middleware.TokenVerifikasiMiddleware())
	r3.POST("", discussionHandler.addDiscussion)
	r3.POST("/answerf/:discussion_id", discussionHandler.addDiscussionFirst)
	r3.POST("/answers/:discussion_first_id", discussionHandler.addDiscussionSecond)
	r3.GET("", discussionHandler.viewAllDiscussion)
	r3.GET("/:idDiscussion", discussionHandler.viewDiscussionDetailByID)
	r3.DELETE("/:discussionID", discussionHandler.deleteDiscussion)
	r3.PUT("/:discussionID", discussionHandler.editDiscussion)
}

func (e *DiscussionHandler) addCatagory(c *gin.Context) {
	user, err := e.userUsecase.AdminOnly(c)
	if err != nil {
		utils.HandleError(c, http.StatusUnauthorized, err.Error())
		return
	}

	var catagory = model.Catagory{}
	err = c.Bind(&catagory)
	if err != nil {
		fmt.Printf("[DiscussionHandler.addCatagory] error bind data body %v \n", err)
		utils.HandleError(c, http.StatusInternalServerError, "Oppss server has be wrong")
		return
	}

	if catagory.Catagory == "" {
		utils.HandleError(c, http.StatusBadRequest, "field are required")
		return
	}

	if catagory.ID != 0 || catagory.UserID != 0 {
		utils.HandleError(c, http.StatusBadRequest, "input is not allowed")
		return
	}

	catagory.UserID = user.ID
	outCatagory, err := e.discussionUsecase.AddCatagory(&catagory)
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.HandleSuccess(c, outCatagory)
}

func (e *DiscussionHandler) deleteCatagory(c *gin.Context) {
	_, err := e.userUsecase.AdminOnly(c)
	if err != nil {
		utils.HandleError(c, http.StatusUnauthorized, err.Error())
		return
	}
	strID := c.Param("id")
	id, err := strconv.Atoi(strID)
	if err != nil {
		utils.HandleError(c, http.StatusBadRequest, "id has be number")
		return
	}
	err = e.discussionUsecase.DeleteCatagory(id)
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.HandleSuccess(c, "success delete catagory")
}

func (e *DiscussionHandler) viewAllCatagory(c *gin.Context) {
	catagories, err := e.discussionUsecase.ViewAllCatagory()
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, err.Error())
		return
	}
	if len((*catagories)) == 0 {
		utils.HandleError(c, http.StatusNotFound, "Catagoriy is empty")
		return
	}
	utils.HandleSuccess(c, catagories)
}

func (e *DiscussionHandler) addDiscussion(c *gin.Context) {
	user, err := e.userUsecase.ValidUser(c)
	if err != nil {
		utils.HandleError(c, http.StatusUnauthorized, err.Error())
		return
	}
	var mDiscussion = []string{"catagoryId", "title", "message"}
	var valueDiscussion []string
	for i := 0; i < len(mDiscussion); i++ {
		value := c.PostForm(mDiscussion[i])
		valueDiscussion = append(valueDiscussion, value)
	}
	catagoryID, err := strconv.Atoi(valueDiscussion[0])
	if err != nil {
		utils.HandleError(c, http.StatusBadRequest, "catagory id has be number")
		return
	}
	_, err = e.discussionUsecase.ViewCatagoryById(catagoryID)
	if err != nil {
		utils.HandleError(c, http.StatusNotFound, err.Error())
		return
	}
	var now = time.Now()
	discussion := model.Discussion{
		UserID:     user.ID,
		CatagoryID: uint(catagoryID),
		Title:      valueDiscussion[1],
		Message:    valueDiscussion[2],
		Date:       now.Format("2006-01-02"),
	}
	outDiscussion, err := e.discussionUsecase.AddDiscussion(c, &discussion)
	if err != nil {
		utils.HandleError(c, http.StatusForbidden, err.Error())
		return
	}
	utils.HandleSuccess(c, outDiscussion)
}

func (e *DiscussionHandler) addDiscussionFirst(c *gin.Context) {
	user, err := e.userUsecase.ValidUser(c)
	if err != nil {
		utils.HandleError(c, http.StatusUnauthorized, err.Error())
		return
	}
	idStr := c.Param("discussion_id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.HandleError(c, http.StatusBadRequest, "id has be number")
		return
	}
	_, err = e.discussionUsecase.ViewDiscussionById(id)
	if err != nil {
		utils.HandleError(c, http.StatusNotFound, err.Error())
		return
	}
	message := c.PostForm("message")
	var now = time.Now()
	var discussionfirst = model.DiscussionFirst{
		UserID:       user.ID,
		DiscussionID: uint(id),
		Message:      message,
		Date:         now.Format("2006-01-02"),
	}
	mDiscussionFirst, err := e.discussionUsecase.AddDiscussionFisrt(c, &discussionfirst)
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.HandleSuccess(c, mDiscussionFirst)
}

func (e *DiscussionHandler) addDiscussionSecond(c *gin.Context) {
	user, err := e.userUsecase.ValidUser(c)
	if err != nil {
		utils.HandleError(c, http.StatusUnauthorized, err.Error())
		return
	}
	idStr := c.Param("discussion_first_id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.HandleError(c, http.StatusBadRequest, "id has be number")
		return
	}
	_, err = e.discussionUsecase.ViewDiscussionFirstById(id)
	if err != nil {
		utils.HandleError(c, http.StatusNotFound, err.Error())
		return
	}
	message := c.PostForm("message")
	var now = time.Now()
	var discussionSecond = model.DiscussionSecond{
		UserID:            user.ID,
		DiscussionFirstID: uint(id),
		Message:           message,
		Date:              now.Format("2006-01-02"),
	}
	mDiscussionSecond, err := e.discussionUsecase.AddDiscussionSecond(c, &discussionSecond)
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.HandleSuccess(c, mDiscussionSecond)
}

func (e *DiscussionHandler) viewAllDiscussion(c *gin.Context) {
	allDiscussion, err := e.discussionUsecase.ViewAllDiscussion()
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, err.Error())
		return
	}
	if len(*allDiscussion) == 0 {
		utils.HandleError(c, http.StatusNotFound, "discussion is empty")
		return
	}
	var arrDiscussionShow []model.DiscussionShow
	var total int
	for i := 0; i < len(*allDiscussion); i++ {
		user, err := e.userUsecase.ViewById(int((*allDiscussion)[i].UserID))
		if err != nil {
			utils.HandleError(c, http.StatusNotFound, err.Error())
			return
		}
		discussionFirst, err := e.discussionUsecase.ViewDiscussionFirstByDiscussionId(int((*allDiscussion)[i].ID))
		if err != nil {
			utils.HandleError(c, http.StatusInternalServerError, err.Error())
			return
		}
		total += len(*discussionFirst)
		for k := 0; k < len(*discussionFirst); k++ {
			discussionSecond, err := e.discussionUsecase.ViewDiscussionSecondByDiscussionFirstId(int((*discussionFirst)[k].ID))
			if err != nil {
				utils.HandleError(c, http.StatusInternalServerError, err.Error())
				return
			}
			total += len(*discussionSecond)
		}
		var discussionShow = model.DiscussionShow{
			ID:           (*allDiscussion)[i].ID,
			Name:         user.Name,
			Date:         (*allDiscussion)[i].Date,
			Message:      (*allDiscussion)[i].Message,
			ProfileImage: user.Image,
			Total:        total,
		}
		arrDiscussionShow = append(arrDiscussionShow, discussionShow)
		total = 0
	}
	utils.HandleSuccess(c, arrDiscussionShow)
}

func (e *DiscussionHandler) viewDiscussionDetailByID(c *gin.Context) {
	idStr := c.Param("idDiscussion")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.HandleError(c, http.StatusBadRequest, "id has be number")
		return
	}
	discussionDetail, err := e.discussionUsecase.ViewDiscussionDetailByID(id)
	if err != nil {
		utils.HandleError(c, http.StatusForbidden, err.Error())
		return
	}
	utils.HandleSuccess(c, discussionDetail)
}

func (e *DiscussionHandler) deleteDiscussion(c *gin.Context) {
	user, err := e.userUsecase.ValidUser(c)
	if err != nil {
		utils.HandleError(c, http.StatusUnauthorized, err.Error())
		return
	}
	idStr := c.Param("discussionID")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.HandleError(c, http.StatusBadRequest, "id has be number")
		return
	}
	discussion, err := e.discussionUsecase.ViewDiscussionById(id)
	if err != nil {
		utils.HandleError(c, http.StatusNotFound, err.Error())
		return
	}
	if discussion.UserID != user.ID {
		utils.HandleError(c, http.StatusBadRequest, "You cannot delete discussion that are not your")
		return
	}
	err = e.discussionUsecase.DeleteDiscussionByID(id)
	if err != nil {
		utils.HandleError(c, http.StatusForbidden, err.Error())
		return
	}
	utils.HandleSuccess(c, "Success delete discussion")
}

func (e *DiscussionHandler) editDiscussion(c *gin.Context) {
	user, err := e.userUsecase.ValidUser(c)
	if err != nil {
		utils.HandleError(c, http.StatusUnauthorized, err.Error())
		return
	}
	idStr := c.Param("discussionID")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.HandleError(c, http.StatusBadRequest, "id has be number")
		return
	}
	var input = []string{"catagoryId", "title", "message"}
	var mDiscussion []string
	for i := 0; i < len(input); i++ {
		output := c.PostForm(input[i])
		mDiscussion = append(mDiscussion, output)
	}
	if mDiscussion[0] == "" || mDiscussion[1] == "" || mDiscussion[2] == "" {
		utils.HandleError(c, http.StatusBadRequest, "field are required")
		return
	}
	catagoryID, err := strconv.Atoi(mDiscussion[0])
	if err != nil {
		utils.HandleError(c, http.StatusBadRequest, "id catagory has be number")
		return
	}

	discussion, err := e.discussionUsecase.ViewDiscussionById(id)
	if err != nil {
		utils.HandleError(c, http.StatusNotFound, err.Error())
		return
	}
	if discussion.UserID != user.ID {
		utils.HandleError(c, http.StatusBadRequest, "You cannot edit discussion that are not your")
		return
	}
	_, err = e.discussionUsecase.ViewCatagoryById(catagoryID)
	if err != nil {
		utils.HandleError(c, http.StatusNotFound, err.Error())
		return
	}
	discussion.CatagoryID = uint(catagoryID)
	discussion.Title = mDiscussion[1]
	discussion.Message = mDiscussion[2]

	discussionShow, err := e.discussionUsecase.UpdateDiscussionById(id, c, discussion)
	if err != nil {
		utils.HandleError(c, http.StatusForbidden, err.Error())
		return
	}
	utils.HandleSuccess(c, discussionShow)
}

