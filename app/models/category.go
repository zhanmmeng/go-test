package models

type Category struct {
	ID
	Name     string `json:"name" gorm:"comment:分类名称"`
	Timestamps
}

type CategoryList []Category