package repository

import (
	"gorm.io/gorm"
	"kouhei-github/sample-gin/service"
	"unicode/utf8"
)

type BlogEntity struct {
	gorm.Model
	Title string `json:"title" binding:"required"`
	Body  string `json:"body"  binding:"required"`
}

func NewBlogEntity(title string, body string) (*BlogEntity, error) {
	if utf8.RuneCountInString(title) <= 2 {
		err := service.MyError{Message: "タイトルを入力してください"}
		return &BlogEntity{}, err
	}
	if body == "" {
		err := service.MyError{Message: "本文を入力してください"}
		return &BlogEntity{}, err
	}
	return &BlogEntity{Title: title, Body: body}, nil
}

func (blog *BlogEntity) CreateBlogEntity() error {
	result := db.Create(blog)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (blog *BlogEntity) FindByTitle() ([]*BlogEntity, error) {
	var users []*BlogEntity
	result := db.Where("title = ?", blog.Title).Find(&users)
	if result.Error != nil {
		return []*BlogEntity{}, result.Error
	}
	return users, nil
}

//func GetAll() (datas []BlogEntity) {
//	result := db.Find(datas)
//	if result.Error != nil {
//		panic(result.Error)
//	}
//	return
//}
//
//func GetOne(id int) (data BlogEntity) {
//	result := db.First(&data, id)
//	if result.Error != nil {
//		panic(result.Error)
//	}
//	return
//}
//
//func (b *BlogEntity) Create() {
//	result := db.Create(b)
//	if result.Error != nil {
//		panic(result.Error)
//	}
//}
//
//func (b *BlogEntity) Update() {
//	result := db.Save(b)
//	if result.Error != nil {
//		panic(result.Error)
//	}
//}
//
//func (b *BlogEntity) Delete() {
//	result := db.Delete(b)
//	if result.Error != nil {
//		panic(result.Error)
//	}
//}
