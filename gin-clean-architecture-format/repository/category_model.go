package repository

import (
	"gorm.io/gorm"
	"kouhei-github/sample-gin/service"
)

type CategoryEntity struct {
	gorm.Model
	Name             string `json:"name"`
	CategoryRakumaId int    `json:"category_rakuma_id"  gorm:"unique"`
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

func CreateCategoryList(categories []CategoryEntity) error {
	result := db.Create(&categories)
	if result.Error != nil {
		myErr := service.MyError{
			Message: result.Error.Error(),
		}
		return myErr
	}
	return nil
}

func FindByCategoryId(id uint) (CategoryEntity, error) {
	entity := CategoryEntity{}
	entity.ID = id
	result := db.First(&entity)
	if result.Error != nil {
		err := service.MyError{Message: result.Error.Error()}
		return CategoryEntity{}, err
	}
	return entity, nil
}

func FindByCategoryName(name string) (*CategoryEntity, error) {
	var entity CategoryEntity
	result := db.Where("name = ?", name).First(&entity)
	if result.Error != nil {
		err := service.MyError{Message: result.Error.Error()}
		return &CategoryEntity{}, err
	}
	return &entity, nil
}
