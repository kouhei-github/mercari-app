package repository

import "gorm.io/gorm"

type AuthenticationEntity struct {
	gorm.Model
	Token  string `json:"token" binding:"required"`
	Cookie string `json:"cookie"  binding:"required"`
}
