package models

import (
	"gorm.io/gorm"
	"time"
)

// ID 自增ID主键
type ID struct {
	ID uint `json:"id" gorm:"primaryKey"`
}

// Timestamps 创建、更新时间
type Timestamps struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// SoftDeletes 软删除
type SoftDeletes struct {
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

// Pagination  分页
type Pagination struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
	Count    int64 `json:"count"`
}
