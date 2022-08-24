package repository

import "gorm.io/gorm"

type DeliveryEntity struct {
	gorm.Model
	Method int `json:"method" binding:"required"`
	Date   int `json:"date" binding:"required"`
	Area   int `json:"area" binding:"required"`
}
