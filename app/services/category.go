package services

import (
	"errors"
	"go-test/app/common/request"
	"go-test/app/models"
	"go-test/conf"
	"strconv"
)

type categoryService struct {

}

var CategoryService = new(categoryService)

func (categoryService *categoryService)CategoryCreate(params request.CategoryCreate) (err error, category models.Category) {
	var result = conf.Global.DB.Where("name = ?", params.Name).Select("id,name").First(&models.Category{})
	if result.RowsAffected != 0 {
		err = errors.New("分类已存在")
		return
	}
	category = models.Category{Name: params.Name}
	err = conf.Global.DB.Create(&category).Error
	return
}

func (categoryService *categoryService) CategoryList() (err error,category []models.Category) {

	var result = conf.Global.DB.Select("id,name").Find(&category)
	if result.Error != nil {
		err = errors.New("分类查询出错")
		return
	}
	return
}

func (categoryService *categoryService)CategoryDel(params request.CategoryDel) (err error) {

	categoryId,_ := strconv.Atoi(params.Id)
	if err:= conf.Global.DB.Where("id = ?", categoryId).Delete(&models.Category{}).Error;err != nil {
		err = errors.New("删除失败")
		return err
	}
	return nil
}

func (categoryService *categoryService) CategoryModify(params request.CategoryModify) (err error, category models.Category) {
	var result = conf.Global.DB.Where("name = ?", params.Name).Select("id,name").First(&models.Category{})
	if result.RowsAffected != 0 {
		err = errors.New("分类已存在")
		return
	}
	category = models.Category{Name: params.Name}
	category.ID.ID = uint(params.Id)
	err = conf.Global.DB.Model(&category).Where("id",params.Id).Update("name",params.Name).Error
	return
}