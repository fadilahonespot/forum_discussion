package handler

import (
	"fmt"
	"kaskus/middleware"
	"kaskus/model"
	"kaskus/user"
	"kaskus/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/xlzd/gotp"
	"golang.org/x/crypto/bcrypt"
)

// UserHandler struct data usecase
type UserHandler struct {
	userUsecase user.UserUsecase
}

// CreateUserHandler is constructor
func CreateUserHandler(r *gin.Engine, userUsecase user.UserUsecase) {
	userHandler := UserHandler{userUsecase}

	r2 := r.Group("/user").Use(middleware.TokenVerifikasiMiddleware())
	r2.GET("", userHandler.viewAllUser)
	r2.GET("/profile", userHandler.profile)
	
	r.POST("/register", userHandler.register)
	r.POST("/login", userHandler.login)
}

func (e *UserHandler) viewAllUser(c *gin.Context) {
	users, err := e.userUsecase.ViewAll()
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, err.Error())
		return
	}
	if len((*users)) == 0 {
		utils.HandleError(c, http.StatusNotFound, "data user is empty")
		return
	}
	utils.HandleSuccess(c, users)
}

func (e *UserHandler) profile(c *gin.Context) {
	user, err := e.userUsecase.ValidUser(c)
	if err != nil {
		utils.HandleError(c, http.StatusUnauthorized, err.Error())
		return
	}
	utils.HandleSuccess(c, user)
}

func (e *UserHandler) register(c *gin.Context) {
	file, err := c.FormFile("image")
	if err != nil {
		utils.HandleError(c, http.StatusBadRequest, "image field are required")
		return
	}

	err = utils.ValidationImages(file.Filename, int(file.Size))
	if err != nil {
		utils.HandleError(c, http.StatusBadRequest, err.Error())
		return
	}

	name := gotp.RandomSecret(12)
	path := viper.GetString("asset.avatar") + name + ".jpg"
	err = c.SaveUploadedFile(file, path)
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, "Upload image failed")
		return
	}

	var dataUser []string
	var param = []string{"name", "email", "password"}
	for i := 0; i < len(param); i++ {
		if c.PostForm(param[i]) == "" {
			utils.HandleError(c, http.StatusBadRequest, "field are required")
			utils.RollbackFile(path)
			return
		}
		dataUser = append(dataUser, c.PostForm(param[i]))
	}

	_, err = e.userUsecase.ViewByEmail(dataUser[1])
	if err == nil {
		utils.HandleError(c, http.StatusConflict, "Email already exsis, use another email")
		utils.RollbackFile(path)
		return
	}

	var user = model.User{
		Name: dataUser[0],
		Image: path,
		Email: dataUser[1],
		Password: dataUser[2],
	}

	inUser, err := e.userUsecase.AddUser(&user)
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, err.Error())
		utils.RollbackFile(path)
		return
	}
	utils.HandleSuccess(c, inUser)
}

func (e *UserHandler) login(c *gin.Context) {
	var inUser = model.User{}
	err := c.Bind(&inUser)
	if err != nil {
		fmt.Printf("[UserHandler.login] error bind data body")
		utils.HandleError(c, http.StatusInternalServerError, "Oopss server someting wrong")
		return
	}
	if inUser.Email == "" || inUser.Password == "" {
		utils.HandleError(c, http.StatusBadRequest, "field are required")
		return
	}
	if inUser.ID != 0 || inUser.Name != "" || inUser.Image != "" {
		utils.HandleError(c, http.StatusBadRequest, "input is not allowed")
		return
	}
	user, err := e.userUsecase.ViewByEmail(inUser.Email)
	if err != nil {
		utils.HandleError(c, http.StatusNotFound, err.Error())
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(inUser.Password))
	if err != nil {
		utils.HandleError(c, http.StatusBadRequest, "Password is wrong")
		return
	}
	token, err := utils.GenerateToken(int(user.ID), user.Password)
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, err.Error())
		return
	}
	var outToken = model.Auth{
		Token: token,
	}

	utils.HandleSuccess(c, outToken)
}