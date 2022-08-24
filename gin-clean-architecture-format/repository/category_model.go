package repository

import (
	"gorm.io/gorm"
	"kouhei-github/sample-gin/service"
)

type CategoryEntity struct {
	gorm.Model
	Name             string `json:"name" binding:"required"`
	CategoryRakumaId int    `json:"category_rakuma_id"  binding:"required"`
}

func NewCategoryEntity(name string, rakumaCategoryId int) *CategoryEntity {
	return &CategoryEntity{
		Name:             name,
		CategoryRakumaId: rakumaCategoryId,
	}
}

func (category *CategoryEntity) Create() error {
	result := db.Create(category)
	if result.Error != nil {
		myErr := service.MyError{Message: result.Error.Error()}
		return myErr
	}
	return nil
}
