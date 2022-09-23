package services

import (
	"errors"
	"go-test/app/common/request"
	"go-test/app/models"
	"go-test/conf"
	"strconv"
)

type article struct {
}

var ArticleService = new(article)

func (articleService *article) ArticleCreate(params request.ArticleCreate) (err error, articleId string) {
	article := models.Article{
		CategoryId: params.CategoryId,
		Title:      params.Title,
		Content:    params.Content,
	}
	err = conf.Global.DB.Create(&article).Error
	if err != nil {
		err = errors.New("添加失败")
	}
	articleId = "添加成功"
	return
}

func (articleService *article) ArticleDel(params request.ArticleDel) error {
	articleId, _ := strconv.Atoi(params.Id)
	if err := conf.Global.DB.Where("id = ?", articleId).Unscoped().Delete(&models.Article{}).Error; err != nil {
		err = errors.New("删除失败")
		return err
	}
	return nil
}

func (articleService *article) ArticleModify(params request.ArticleModify) (err error) {

	result := conf.Global.DB.Where("id = ?", params.Id).First(&models.Article{})
	if result.RowsAffected == 0 {
		err := errors.New("该文章不存在")
		return err
	}
	article := models.Article{
		CategoryId: params.CategoryId,
		Title:      params.Title,
		Content:    params.Content,
	}
	err = conf.Global.DB.Model(&article).Where("id", params.Id).Updates(article).Error
	if err != nil {
		err = errors.New("修改失败")
	}
	return err
}

type ArticleListRes struct {
	articles []models.ArticleJoinCategory
	models.Pagination
}

func (articleService *article) ArticleList(params request.ArticleList) (err error, article []models.ArticleJoinCategory, pageInfo ArticleListRes) {

	page := params.Page
	pageSize := params.PageSize
	categoryId := params.CategoryId

	if page == 0 {
		page = 1
	}
	if pageSize == 0 {
		pageSize = 10
	}
	offset := (page - 1) * pageSize


	result := conf.Global.DB.
		Order("articles.id desc").
		Table("articles").
		Select("articles.id,articles.content,articles.title,articles.category_id,categories.name as category_name,articles.created_at").
		Joins("left join categories on categories.id = articles.category_id").
		Offset(offset).
		Limit(pageSize).
		Scopes(models.WhereCategory(categoryId)).
		Find(&article)

		conf.Global.
		DB.Table("articles").Count(&pageInfo.Count)

	if result.Error != nil {
		return
	}

	pageInfo.Page = page
	pageInfo.PageSize = pageSize
	return
}

func (articleService *article) ArticleDetail(params request.ArticleDel) (err error, article models.ArticleJoinCategory) {
	result := conf.Global.
		DB.
		Table("articles").
		Where("articles.id = ?", params.Id).
		Select("articles.id,articles.content,articles.title,articles.category_id,categories.name as category_name,articles.created_at").
		Joins("left join categories on categories.id = articles.category_id").
		First(&article)

	if result.Error != nil {
		return
	}

	return
}
