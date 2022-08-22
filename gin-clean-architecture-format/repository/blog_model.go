package repository

import (
	"gorm.io/gorm"
)

type BlogEntity struct {
	gorm.Model
	Title string
	Body  string
}

func GetAll() (datas []BlogEntity) {
	result := db.Find(datas)
	if result.Error != nil {
		panic(result.Error)
	}
	return
}

func GetOne(id int) (data BlogEntity) {
	result := db.First(&data, id)
	if result.Error != nil {
		panic(result.Error)
	}
	return
}

func (b *BlogEntity) Create() {
	result := db.Create(b)
	if result.Error != nil {
		panic(result.Error)
	}
}

func (b *BlogEntity) Update() {
	result := db.Save(b)
	if result.Error != nil {
		panic(result.Error)
	}
}

func (b *BlogEntity) Delete() {
	result := db.Delete(b)
	if result.Error != nil {
		panic(result.Error)
	}
}
