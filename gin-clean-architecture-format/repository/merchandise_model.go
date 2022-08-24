package repository

import "gorm.io/gorm"

type MerchandiseEntity struct {
	gorm.Model
	Image            string `json:"image" binding:"required"`
	Name             string `json:"name" binding:"required"`
	Detail           string `json:"detail" binding:"required"`
	Status           int    `json:"status" binding:"required"`
	Carriage         int    `json:"carriage" binding:"required"`
	RequestRequired  int    `json:"request_required" binding:"required"`
	SellPrice        int    `json:"sell_price" binding:"required"`
	DeliveryEntityID int    `json:"delivery_id"`
	DeliveryEntity   DeliveryEntity
	CategoryEntityID int `json:"category_id"`
	CategoryEntity   CategoryEntity
}
