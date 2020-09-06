package usecase

import (
	"fmt"
	"kaskus/discussion"
	"kaskus/model"
	"kaskus/user"
	"kaskus/utils"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/xlzd/gotp"
)

type DiscussionUsecaseImpl struct {
	discussionRepo discussion.DiscussionRepo
	userRepo       user.UserRepo
}

func CreateDiscussionUsecase(discussionRepo discussion.DiscussionRepo, userRepo user.UserRepo) discussion.DiscussionUsecase {
	return &DiscussionUsecaseImpl{discussionRepo, userRepo}
}

func (e *DiscussionUsecaseImpl) AddCatagory(catagory *model.Catagory) (*model.Catagory, error) {
	return e.discussionRepo.AddCatagory(catagory)
}

func (e *DiscussionUsecaseImpl) DeleteCatagory(id int) error {
	return e.discussionRepo.DeleteCatagory(id)
}

func (e *DiscussionUsecaseImpl) ViewAllCatagory() (*[]model.Catagory, error) {
	return e.discussionRepo.ViewAllCatagory()
}

func (e *DiscussionUsecaseImpl) ViewCatagoryById(id int) (*model.Catagory, error) {
	return e.discussionRepo.ViewCatagoryById(id)
}

func fileUploadDiscussionMultiPart(c *gin.Context) ([]string, []string, error) {
	form, err := c.MultipartForm()
	if err != nil {
		fmt.Printf("[discussionUsecaseImpl.uploadFileDiscussionMultipart] error multipart form %v \n", err)
		return nil, nil, fmt.Errorf("failed upload image")
	}
	images := form.File["images"]
	var partImages []string
	for _, image := range images {
		err = utils.ValidationImages(image.Filename, int(image.Size))
		if err != nil {
			utils.RollbackFiles(partImages)
			return nil, nil, err
		}
		name := gotp.RandomSecret(10)
		part := viper.GetString("asset.images") + name + ".jpeg"
		err := c.SaveUploadedFile(image, part)
		if err != nil {
			fmt.Printf("[DiscussionUsecaseImpl.uploadFileDiscussionMultiPart] error save file upload %v \n", err)
			utils.RollbackFiles(partImages)
			return nil, nil, fmt.Errorf("failed upload image")
		}
		partImages = append(partImages, part)
	}

	var partFiles []string
	files := form.File["files"]
	for _, file := range files {
		err := utils.ValidationFiles(file.Filename, int(file.Size))
		if err != nil {
			utils.RollbackFiles(partFiles)
			utils.RollbackFiles(partImages)
			return nil, nil, err
		}
		part := viper.GetString("asset.files") + file.Filename
		err = c.SaveUploadedFile(file, part)
		if err != nil {
			fmt.Printf("[DiscussionUsecaseImpl.AddDiscussion] error save file upload %v \n", err)
			utils.RollbackFiles(partFiles)
			utils.RollbackFiles(partImages)
			return nil, nil, fmt.Errorf("failed upload file")
		}
		partFiles = append(partFiles, part)
	}
	return partImages, partFiles, nil
}

func (e *DiscussionUsecaseImpl) AddDiscussion(c *gin.Context, discussion *model.Discussion) (*model.DiscussionPost, error) {
	tx := e.discussionRepo.BeginTrans()
	outDiscussion, err := e.discussionRepo.AddDiscussion(discussion, tx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	partImages, partFiles, err := fileUploadDiscussionMultiPart(c)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	for i := 0; i < len(partImages); i++ {
		var discussionImage = model.DiscussionImages{
			DiscussionID: outDiscussion.ID,
			UserID:       outDiscussion.UserID,
			Image:        partImages[i],
		}
		_, err := e.discussionRepo.AddDiscussionImages(&discussionImage, tx)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	for k := 0; k < len(partFiles); k++ {
		var discussionFile = model.DiscussionFiles{
			DiscussionID: outDiscussion.ID,
			UserID:       outDiscussion.UserID,
			File:         partFiles[k],
		}
		_, err := e.discussionRepo.AddDiscussionFiles(&discussionFile, tx)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
	}
	var discussionPost = model.DiscussionPost{
		DiscussionID: outDiscussion.ID,
		UserID:       outDiscussion.UserID,
		CatagoryID:   outDiscussion.CatagoryID,
		Title:        outDiscussion.Title,
		Date:         outDiscussion.Date,
		Message:      outDiscussion.Message,
		Images:       partImages,
		Files:        partFiles,
	}
	tx.Commit()
	return &discussionPost, nil
}

func fileUploadDiscussionAnswear(c *gin.Context) (string, string, error) {
	var pathFile string
	var pathImage string
	image, _ := c.FormFile("image")
	if image != nil {
		err := utils.ValidationImages(image.Filename, int(image.Size))
		if err != nil {
			return "", "", err
		}

		name := gotp.RandomSecret(10)
		pathImage = viper.GetString("asset.images") + name + ".jpeg"
		err = c.SaveUploadedFile(image, pathImage)
		if err != nil {
			fmt.Printf("[DiscussionUsecaseImpl.FileUploadDiscussionAnswer] error upload image %v \n", err)
			return "", "", fmt.Errorf("failed to upload image")
		}
	}

	file, _ := c.FormFile("file")
	if file != nil {
		err := utils.ValidationFiles(file.Filename, int(file.Size))
		if err != nil {
			return "", "", err
		}
		pathFile = viper.GetString("asset.files") + file.Filename
		err = c.SaveUploadedFile(file, pathFile)
		if err != nil {
			fmt.Printf("[DiscussionUsecaseImpl.FileUploadDiscussionAnswer] error upload file %v \n", err)
			return "", "", fmt.Errorf("failed to upload file")
		}
	}

	return pathImage, pathFile, nil
}

func (e *DiscussionUsecaseImpl) AddDiscussionFisrt(c *gin.Context, discussionFirst *model.DiscussionFirst) (*model.DiscussionFirst, error) {
	pathImage, pathFile, err := fileUploadDiscussionAnswear(c)
	if err != nil {
		return nil, err
	}
	discussionFirst.Image = pathImage
	discussionFirst.File = pathFile
	return e.discussionRepo.AddDiscussionFisrt(discussionFirst)
}

func (e *DiscussionUsecaseImpl) AddDiscussionSecond(c *gin.Context, discussionSecond *model.DiscussionSecond) (*model.DiscussionSecond, error) {
	pathImage, pathFile, err := fileUploadDiscussionAnswear(c)
	if err != nil {
		return nil, err
	}
	discussionSecond.Image = pathImage
	discussionSecond.File = pathFile
	return e.discussionRepo.AddDiscussionSecond(discussionSecond)
}

func (e *DiscussionUsecaseImpl) ViewDiscussionById(id int) (*model.Discussion, error) {
	return e.discussionRepo.ViewDiscussionById(id)
}

func (e *DiscussionUsecaseImpl) ViewDiscussionFirstById(id int) (*model.DiscussionFirst, error) {
	return e.discussionRepo.ViewDiscussionFirstById(id)
}

func (e *DiscussionUsecaseImpl) ViewAllDiscussion() (*[]model.Discussion, error) {
	return e.discussionRepo.ViewAllDiscussion()
}

func (e *DiscussionUsecaseImpl) ViewDiscussionFirstByDiscussionId(id int) (*[]model.DiscussionFirst, error) {
	return e.discussionRepo.ViewDiscussionFirstByDiscussionId(id)
}

func (e *DiscussionUsecaseImpl) ViewDiscussionSecondByDiscussionFirstId(id int) (*[]model.DiscussionSecond, error) {
	return e.discussionRepo.ViewDiscussionSecondByDiscussionFirstId(id)
}

func (e *DiscussionUsecaseImpl) ViewDiscussionDetailByID(id int) (*model.DiscussionDetailShow, error) {
	discussion, err := e.discussionRepo.ViewDiscussionById(id)
	if err != nil {
		return nil, err
	}
	user, err := e.userRepo.ViewById(int(discussion.UserID))
	if err != nil {
		return nil, err
	}
	discussionImages, err := e.discussionRepo.ViewDiscussionImageByDiscussionID(int(discussion.ID))
	if err != nil {
		return nil, err
	}
	var images []string
	for i := 0; i < len(*discussionImages); i++ {
		images = append(images, (*discussionImages)[i].Image)
	}
	discussionFile, err := e.discussionRepo.ViewDiscussionFileByDiscussionID(int(discussion.ID))
	if err != nil {
		return nil, err
	}
	var files []string
	for k := 0; k < len(*discussionFile); k++ {
		files = append(files, (*discussionFile)[k].File)
	}
	discussionFirst, err := e.detailDiscussionFirst(int(discussion.ID))
	if err != nil {
		return nil, err
	}
	var discussionDetailShow = model.DiscussionDetailShow{
		ID:             discussion.ID,
		Name:           user.Name,
		Date:           discussion.Date,
		Message:        discussion.Message,
		ProfileImage:   user.Image,
		Images:         images,
		Files:          files,
		FirsDiscussion: (*discussionFirst),
	}
	return &discussionDetailShow, nil
}

func (e *DiscussionUsecaseImpl) detailDiscussionFirst(discussionID int) (*[]model.DiscussionFirstDetailShow, error) {
	discussionFirst, err := e.discussionRepo.ViewDiscussionFirstByDiscussionId(discussionID)
	if err != nil {
		return nil, err
	}
	var arrDiscussionDetailFirst []model.DiscussionFirstDetailShow
	for i := 0; i < len(*discussionFirst); i++ {
		user, err := e.userRepo.ViewById(int((*discussionFirst)[i].UserID))
		if err != nil {
			return nil, err
		}
		discussionSecond, err := e.detailDiscussionSecond(int((*discussionFirst)[i].ID))
		if err != nil {
			return nil, err
		}
		var discussionFirst = model.DiscussionFirstDetailShow{
			ID:               (*discussionFirst)[i].ID,
			Name:             user.Name,
			Date:             (*discussionFirst)[i].Date,
			Message:          (*discussionFirst)[i].Message,
			ProfileImage:     user.Image,
			File:             (*discussionFirst)[i].File,
			Image:            (*discussionFirst)[i].Image,
			SecondDiscussion: (*discussionSecond),
		}
		arrDiscussionDetailFirst = append(arrDiscussionDetailFirst, discussionFirst)
	}
	return &arrDiscussionDetailFirst, nil
}

func (e *DiscussionUsecaseImpl) detailDiscussionSecond(firstDiscussionID int) (*[]model.DiscussionSecondDetailShow, error) {
	discussionSecond, err := e.discussionRepo.ViewDiscussionSecondByDiscussionFirstId(firstDiscussionID)
	if err != nil {
		return nil, err
	}
	var arrShowDiscussionSecond []model.DiscussionSecondDetailShow
	for i := 0; i < len(*discussionSecond); i++ {
		user, err := e.userRepo.ViewById(int((*discussionSecond)[i].UserID))
		if err != nil {
			return nil, err
		}
		var showDiscussionSecond = model.DiscussionSecondDetailShow{
			ID:           (*discussionSecond)[i].ID,
			Name:         user.Name,
			Message:      (*discussionSecond)[i].Message,
			Date:         (*discussionSecond)[i].Date,
			ProfileImage: user.Image,
			Image:        (*discussionSecond)[i].Image,
			File:         (*discussionSecond)[i].File,
		}
		arrShowDiscussionSecond = append(arrShowDiscussionSecond, showDiscussionSecond)
	}
	return &arrShowDiscussionSecond, nil
}

func (e *DiscussionUsecaseImpl) DeleteDiscussionByID(id int) error {
	tx := e.discussionRepo.BeginTrans()
	var files []string
	var images []string
	discussionImages, err := e.discussionRepo.ViewImagesByDiscussionID(id)
	if err != nil {
		return err
	}
	for l := 0; l < len(*discussionImages); l++ {
		images = append(images, (*discussionImages)[l].Image)
	}
	discussionFiles, err := e.discussionRepo.ViewFilesByDiscussionID(id)
	if err != nil {
		return err
	}
	for j := 0; j < len(*discussionFiles); j++ {
		files = append(files, (*discussionFiles)[j].File)
	}
	discussionFirst, err := e.discussionRepo.ViewDiscussionFirstByDiscussionId(id)
	if err != nil {
		return err
	}
	for i := 0; i < len(*discussionFirst); i++ {
		files = append(files, (*discussionFirst)[i].File)
		images = append(images, (*discussionFirst)[i].Image)
		err := e.discussionRepo.DeleteDiscussionFirstByID(int((*discussionFirst)[i].ID), tx)
		if err != nil {
			tx.Rollback()
			return err
		}
		discussionSecond, err := e.discussionRepo.ViewDiscussionSecondByDiscussionFirstId(int((*discussionFirst)[i].ID))
		if err != nil {
			tx.Rollback()
			return err
		}
		for k := 0; k < len(*discussionSecond); k++ {
			files = append(files, (*discussionSecond)[k].File)
			images = append(images, (*discussionSecond)[k].Image)
			err := e.discussionRepo.DeleteDiscussionSecondByID(int((*discussionSecond)[k].ID), tx)
			if err != nil {
				tx.Rollback()
				return err
			}
		}
	}
	err = e.discussionRepo.DeleteDiscussionFilesByDiscussionID(id, tx)
	if err != nil {
		tx.Rollback()
		return err
	}
	err = e.discussionRepo.DeleteDiscussionImagesByDiscussionID(id, tx)
	if err != nil {
		tx.Rollback()
		return err
	}
	err = e.discussionRepo.DeleteDiscussionById(id, tx)
	if err != nil {
		tx.Rollback()
		return err
	}
	utils.RollbackFiles(files)
	utils.RollbackFiles(images)
	tx.Commit()

	return nil
}

func (e *DiscussionUsecaseImpl) UpdateDiscussionById(id int, c *gin.Context, discussion *model.Discussion) (*model.DiscussionPost, error) {
	tx := e.discussionRepo.BeginTrans()
	partImages, partFiles, err := fileUploadDiscussionMultiPart(c)
	if err != nil {
		return nil, err
	}

	if partImages != nil {
		err := e.discussionRepo.DeleteDiscussionImagesByDiscussionID(id, tx)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
		for i := 0; i < len(partImages); i++ {
			var discussionImages = model.DiscussionImages{
				DiscussionID: uint(id),
				UserID:       discussion.UserID,
				Image:        partImages[i],
			}
			_, err := e.discussionRepo.AddDiscussionImages(&discussionImages, tx)
			if err != nil {
				tx.Rollback()
				return nil, err
			}
		}
		utils.RollbackFiles(partImages)
	}

	if partFiles != nil {
		err := e.discussionRepo.DeleteDiscussionFilesByDiscussionID(id, tx)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
		for i := 0; i < len(partFiles); i++ {
			var discussionFiles = model.DiscussionFiles{
				DiscussionID: uint(id),
				UserID:       discussion.UserID,
				File:         partFiles[i],
			}
			_, err := e.discussionRepo.AddDiscussionFiles(&discussionFiles, tx)
			if err != nil {
				tx.Rollback()
				return nil, err
			}
		}
		utils.RollbackFiles(partFiles)
	}
	mDiscussion, err := e.discussionRepo.UpdateDiscussionById(id, discussion, tx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	var discussionPost = model.DiscussionPost{
		DiscussionID: mDiscussion.ID,
		UserID:       mDiscussion.UserID,
		CatagoryID:   mDiscussion.CatagoryID,
		Title:        mDiscussion.Title,
		Date:         mDiscussion.Date,
		Message:      mDiscussion.Message,
		Images:       partImages,
		Files:        partFiles,
	}
	tx.Commit()
	return &discussionPost, nil
}
