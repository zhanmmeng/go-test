package models

import (
	"gorm.io/gorm"
	"strconv"
	"time"
)

type Article struct {
	gorm.Model
	Category   Category `json:"category"`
	CategoryId int      `json:"category_id"`
	Content    string   `json:"content" gorm:"type:text;comment:文章内容"`
	Title      string   `json:"title" gorm:"comment: 文章标题 "`
}

func (article Article) GetUid() string {
	return strconv.Itoa(int(article.ID))
}

type ArticleJoinCategory struct {
	Id           uint      `json:"id"`
	CategoryId   uint      `json:"category_id"`
	CategoryName string    `json:"category_name"`
	Content      string    `json:"content"`
	Title        string    `json:"title"`
	CreatedAt    time.Time `json:"created_at"`
}

func WhereCategory(categoryId int) func (db *gorm.DB) *gorm.DB {

	if categoryId == 0 {
		return func(db *gorm.DB) *gorm.DB {
			return db
		}
	}

	return func (db *gorm.DB) *gorm.DB {
		return db.Where("articles.category_id = ?", categoryId)
	}
}