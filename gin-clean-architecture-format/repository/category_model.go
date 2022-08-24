package repository

import "gorm.io/gorm"

type CategoryEntity struct {
	gorm.Model
	Name             string `json:"name" binding:"required"`
	CategoryRakumaId string `json:"category_rakuma_id"  binding:"required"`
}
