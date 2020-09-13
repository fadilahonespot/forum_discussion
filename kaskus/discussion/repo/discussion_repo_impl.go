package repo

import (
	"fmt"
	"kaskus/discussion"
	"kaskus/model"

	"github.com/jinzhu/gorm"
)

type DiscussionRepoImpl struct {
	DB *gorm.DB
}

func CreateDiscussionRepo(DB *gorm.DB) discussion.DiscussionRepo {
	return &DiscussionRepoImpl{DB}
}

func (e *DiscussionRepoImpl) BeginTrans() *gorm.DB {
	return e.DB.Begin()
}

func (e *DiscussionRepoImpl) AddCatagory(catagory *model.Catagory) (*model.Catagory, error) {
	err := e.DB.Save(&catagory).Error
	if err != nil {
		fmt.Printf("[discussionRepoImpl.AddCatagory] error execute query %v \n", err)
		return nil, fmt.Errorf("failed add data catagory")
	}
	return catagory, nil
}

func (e *DiscussionRepoImpl) ViewAllCatagory() (*[]model.Catagory, error) {
	var catagories []model.Catagory
	err := e.DB.Find(&catagories).Error
	if err != nil {
		fmt.Printf("[DiscussionRepoImpl.ViewAllCatagory] error execute query %v \n", err)
		return nil, fmt.Errorf("failed show all catagory")
	}
	return &catagories, nil
}

func (e *DiscussionRepoImpl) ViewCatagoryById(id int) (*model.Catagory, error) {
	var catagory = model.Catagory{}
	err := e.DB.Table("catagory").Where("id = ?", id).First(&catagory).Error
	if err != nil {
		fmt.Printf("[DiscussionRepoImpl.ViewCatagoryById] error execute query %v \n", err)
		return nil, fmt.Errorf("catagory is not exsis")
	}
	return &catagory, nil
}

func (e *DiscussionRepoImpl) UpdateCatagoryById(id int, catagory *model.Catagory) (*model.Catagory, error) {
	var upCatagory = model.Catagory{}
	err := e.DB.Table("catagory").Where("id = ?", id).First(&upCatagory).Update(&catagory).Error
	if err != nil {
		fmt.Printf("[DiscussionRepoImpl.UpdateCatagorById] error execute query %v \n", err)
		return nil, fmt.Errorf("failed to update catagory")
	}
	return &upCatagory, nil
}

func (e *DiscussionRepoImpl) DeleteCatagoryById(id int) error {
	var catagory = model.Catagory{}
	err := e.DB.Table("catagory").Where("id = ?", id).First(&catagory).Delete(&catagory).Error
	if err != nil {
		fmt.Printf("[DiscussionRepo.DeleteCatagory] error execute query %v \n", err)
		return fmt.Errorf("catagory is not exsis")
	}
	return nil
}

func (e *DiscussionRepoImpl) ViewAllDiscussion() (*[]model.Discussion, error) {
	var discussion []model.Discussion
	err := e.DB.Find(&discussion).Error
	if err != nil {
		fmt.Printf("[DiscussionRepoImpl.ViewAll] error execute query %v \n", err)
		return nil, fmt.Errorf("failed to show all discussion")
	}
	return &discussion, nil
}

func (e *DiscussionRepoImpl) ViewDiscussionById(id int) (*model.Discussion, error) {
	var discussion = model.Discussion{}
	err := e.DB.Table("discussion").Where("id = ?", id).First(&discussion).Error
	if err != nil {
		fmt.Printf("[discussionRepoImpl.ViewDiscussionById] error execute query %v \n", err)
		return nil, fmt.Errorf("discussion is not exsis")
	}
	return &discussion, nil
}

func (e *DiscussionRepoImpl) ViewDiscussionImageByDiscussionID(id int) (*[]model.DiscussionImages, error) {
	var discussionImage []model.DiscussionImages
	err := e.DB.Table("discussion_images").Where("discussion_id = ?", id).Find(&discussionImage).Error
	if err != nil {
		fmt.Printf("[DiscussionRepoImpl.ViewDiscussionImageByDiscussionID] error execute query %v \n", err)
		return nil, fmt.Errorf("failed to show images")
	}
	return &discussionImage, nil
}

func (e *DiscussionRepoImpl) ViewDiscussionFileByDiscussionID(id int) (*[]model.DiscussionFiles, error) {
	var discussionFile []model.DiscussionFiles
	err := e.DB.Table("discussion_files").Where("discussion_id = ?", id).Find(&discussionFile).Error
	if err != nil {
		fmt.Printf("[DiscussionRepoImpl.ViewDiscussionfileByDiscussionID] error execute query %v \n", err)
		return nil, fmt.Errorf("failed to show files")
	}
	return &discussionFile, nil
}

func (e *DiscussionRepoImpl) ViewDiscussionFirstById(id int) (*model.DiscussionFirst, error) {
	var discussionFirst = model.DiscussionFirst{}
	err := e.DB.Table("discussion_first").Where("id = ?", id).First(&discussionFirst).Error
	if err != nil {
		fmt.Printf("[DiscussionRepoImpl.ViewDiscussionFirstById] error execute query %v \n", err)
		return nil, fmt.Errorf("discussion answer is not exsis")
	}
	return &discussionFirst, nil
}

func (e *DiscussionRepoImpl) AddDiscussion(discussion *model.Discussion, tx *gorm.DB) (*model.Discussion, error) {
	err := tx.Save(&discussion).Error
	if err != nil {
		fmt.Printf("[DiscussionRepoImpl.AddDiscussion] error execute query %v \n", err)
		return nil, fmt.Errorf("failed add new catagory")
	}
	return discussion, nil
}

func (e *DiscussionRepoImpl) DeleteDiscussionById(id int, tx *gorm.DB) error {
	var discussion = model.Discussion{}
	err := tx.Table("discussion").Where("id = ?", id).First(&discussion).Delete(&discussion).Error
	if err != nil {
		fmt.Printf("[DiscussionRepoImpl.DeleteDiscussion] error execute query %v \n", err)
		return fmt.Errorf("discussion is not exsis")
	}
	return nil
}

func (e *DiscussionRepoImpl) DeleteDiscussionFirstByID(id int, tx *gorm.DB) error {
	var discussionFirst = model.DiscussionFirst{}
	err := tx.Table("discussion_first").Where("id = ?", id).First(&discussionFirst).Delete(&discussionFirst).Error
	if err != nil {
		fmt.Printf("[DiscussionRepoImpl.deletediscussionfirst] error execute query %v \n", err)
		return fmt.Errorf("discussion first is not exsis")
	}
	return nil
}

func (e *DiscussionRepoImpl) DeleteDiscussionSecondByID(id int, tx *gorm.DB) error {
	var discussionSecond = model.DiscussionSecond{}
	err := tx.Table("discussion_second").Where("id = ?", id).First(&discussionSecond).Delete(&discussionSecond).Error
	if err != nil {
		fmt.Printf("[DiscussionRepoImpl.DeleteDiscussionSecondByDiscussionFirst] error execute query %v \n", err)
		return fmt.Errorf("failed to delete discussion second")
	}
	return nil
}
func (e *DiscussionRepoImpl) AddDiscussionImages(discussionImages *model.DiscussionImages, tx *gorm.DB) (*model.DiscussionImages, error) {
	err := tx.Save(&discussionImages).Error
	if err != nil {
		fmt.Printf("[DiscussionRepoImpl.AdddiscussionImages] error execute query %v \n", err)
		return nil, fmt.Errorf("failed to add images")
	}
	return discussionImages, nil
}

func (e *DiscussionRepoImpl) AddDiscussionFiles(discussionFiles *model.DiscussionFiles, tx *gorm.DB) (*model.DiscussionFiles, error) {
	err := tx.Save(&discussionFiles).Error
	if err != nil {
		fmt.Printf("[DiscussionRepoImpl.AddDiscussionFiles] error execute query %v \n", err)
		return nil, fmt.Errorf("failed to add files")
	}
	return discussionFiles, nil
}

func (e *DiscussionRepoImpl) AddDiscussionFisrt(discussionFirst *model.DiscussionFirst) (*model.DiscussionFirst, error) {
	err := e.DB.Save(&discussionFirst).Error
	if err != nil {
		fmt.Printf("[DiscussionRepoImpl.AddDiscussionFirst] error execute query %v \n", err)
		return nil, fmt.Errorf("failed to answer discussion")
	}
	return discussionFirst, nil
}

func (e *DiscussionRepoImpl) AddDiscussionSecond(discussionSecond *model.DiscussionSecond) (*model.DiscussionSecond, error) {
	err := e.DB.Save(&discussionSecond).Error
	if err != nil {
		fmt.Printf("[DiscussionRepoImpl.AddDiscussionSecond] error execute query %v \n", err)
		return nil, fmt.Errorf("failed to answer discussion")
	}
	return discussionSecond, nil
}

func (e *DiscussionRepoImpl) ViewDiscussionFirstByDiscussionId(id int) (*[]model.DiscussionFirst, error) {
	var discussionFirst []model.DiscussionFirst
	err := e.DB.Table("discussion_first").Where("discussion_id = ?", id).Find(&discussionFirst).Error
	if err != nil {
		fmt.Printf("[DiscussionRepoImpl.ViewDiscussionFirstByDiscussionId] error execute query %v \n", err)
		return nil, fmt.Errorf("failed to show discussion")
	}
	return &discussionFirst, nil
}

func (e *DiscussionRepoImpl) ViewDiscussionSecondByDiscussionFirstId(id int) (*[]model.DiscussionSecond, error) {
	var discussionSecond []model.DiscussionSecond
	err := e.DB.Table("discussion_second").Where("discussion_first_id = ?", id).Find(&discussionSecond).Error
	if err != nil {
		fmt.Printf("[DiscussionRepoImpl.ViewDiscussionSecondByDiscussionFirstId] error execute query %v \n", err)
		return nil, fmt.Errorf("failed to show discussion")
	}
	return &discussionSecond, nil
}

func (e *DiscussionRepoImpl) ViewImagesByDiscussionID(id int) (*[]model.DiscussionImages, error) {
	var images []model.DiscussionImages
	err := e.DB.Table("discussion_images").Where("discussion_id = ?", id).Find(&images).Error
	if err != nil {
		fmt.Printf("[DiscussionRepoImpl.ViewImagesByDiscussionID] error execute query %v \n", err)
		return nil, fmt.Errorf("failed to show images discussion")
	}
	return &images, nil
}

func (e *DiscussionRepoImpl) ViewFilesByDiscussionID(id int) (*[]model.DiscussionFiles, error) {
	var files []model.DiscussionFiles
	err := e.DB.Table("discussion_files").Where("discussion_id = ?", id).Find(&files).Error
	if err != nil {
		fmt.Printf("[DiscussionRepoImpl.ViewFilesByDiscussionID] error execute query %v \n", err)
		return nil, fmt.Errorf("failed to show files discussion")
	}
	return &files, nil
}

func (e *DiscussionRepoImpl) DeleteDiscussionFilesByDiscussionID(id int, tx *gorm.DB) error {
	var discussionFiles []model.DiscussionFiles
	err := tx.Table("discussion_files").Where("discussion_id = ?", id).Find(&discussionFiles).Delete(&discussionFiles).Error
	if err != nil {
		fmt.Printf("[DiscussionRepoImp.DeleteDiscussinFilesByDiscussionID] error execute query %v \n", err)
		return fmt.Errorf("failed to delete discussion files")
	}
	return nil
}

func (e *DiscussionRepoImpl) DeleteDiscussionImagesByDiscussionID(id int, tx *gorm.DB) error {
	var discussionImages []model.DiscussionImages
	err := tx.Table("discussion_images").Where("discussion_id = ?", id).Find(&discussionImages).Delete(&discussionImages).Error
	if err != nil {
		fmt.Printf("[DiscussionRepoImpl.DeleteDiscussionImagesByDiscussionID] error execute query %v \n", err)
		return fmt.Errorf("failed to delete discussion images")
	}
	return nil
}

func (e *DiscussionRepoImpl) UpdateDiscussionById(id int, discussion *model.Discussion, tx *gorm.DB) (*model.Discussion, error) {
	var upDiscussion = model.Discussion{}
	err := tx.Table("discussion").Where("id = ?", id).First(&upDiscussion).Update(&discussion).Error
	if err != nil {
		fmt.Printf("[DiscussionRepoImpl.UpdateDiscussionByID] error execute query %v \n", err)
		return nil, fmt.Errorf("failed update discussion")
	}
	return &upDiscussion, nil
}
