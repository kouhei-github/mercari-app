package repository

import (
	"fmt"
	"gorm.io/gorm"
	"kouhei-github/sample-gin/service"
)

type DeliveryEntity struct {
	gorm.Model
	Method int `json:"method"`
	Date   int `json:"date"`
	Area   int `json:"area"`
}

func NewDeliveryEntity(method int, date int, area int) (*DeliveryEntity, error) {
	return &DeliveryEntity{Method: method, Date: date, Area: area}, nil
}

func (entity *DeliveryEntity) Create() error {
	result := db.Create(entity)
	if result.Error != nil {
		myErr := service.MyError{Message: result.Error.Error()}
		return myErr
	}
	return nil
}

func FindByDeliveryId(id uint) (DeliveryEntity, error) {
	entity := DeliveryEntity{}
	entity.ID = id
	fmt.Println(id)
	result := db.First(&entity)
	fmt.Println(entity)
	if result.Error != nil {
		err := service.MyError{Message: result.Error.Error()}
		return DeliveryEntity{}, err
	}
	return entity, nil
}
