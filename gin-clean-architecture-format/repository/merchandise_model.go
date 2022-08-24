package repository

import (
	"fmt"
	"gorm.io/gorm"
	"kouhei-github/sample-gin/service"
	"unicode/utf8"
)

type MerchandiseEntity struct {
	gorm.Model
	Image            string `json:"image" binding:"required"`
	Name             string `json:"name" binding:"required"`
	Detail           string `json:"detail" binding:"required"`
	Status           int    `json:"status" binding:"required"`
	Carriage         int    `json:"carriage" binding:"required"`
	RequestRequired  int    `json:"request_required" binding:"required"`
	SellPrice        int    `json:"sell_price" binding:"required"`
	IsUpload         string `json:"is_upload"`
	IsPurchased      string `json:"is_purchased"`
	DeliveryEntityID uint   `json:"delivery_id" binding:"required"`
	DeliveryEntity   DeliveryEntity
	CategoryEntityID uint `json:"category_id" binding:"required"`
	CategoryEntity   CategoryEntity
}

func NewMerchandiseEntity(
	image string,
	name string,
	detail string,
	status int,
	carriage int,
	requestRequired int,
	sellPrice int,
	deliveryEntityID uint,
	categoryEntityID uint,
) (*MerchandiseEntity, error) {
	if utf8.RuneCountInString(image) <= 1 {
		err := service.MyError{Message: "画像を入力してください"}
		return &MerchandiseEntity{}, err
	}
	if utf8.RuneCountInString(name) <= 1 {
		err := service.MyError{Message: "商品名を入力してください"}
		return &MerchandiseEntity{}, err
	}
	if utf8.RuneCountInString(detail) <= 1 {
		err := service.MyError{Message: "商品の説明を入力してください"}
		return &MerchandiseEntity{}, err
	}
	entity := MerchandiseEntity{
		Image:            image,
		Name:             name,
		Detail:           detail,
		Status:           status,
		Carriage:         carriage,
		RequestRequired:  requestRequired,
		SellPrice:        sellPrice,
		DeliveryEntityID: deliveryEntityID,
		CategoryEntityID: categoryEntityID,
	}
	return &entity, nil
}

func (receiver *MerchandiseEntity) Create() error {
	// カテゴリーのEntityの作成

	categoryEntity, err := FindByCategoryId(receiver.CategoryEntityID)
	if err != nil {
		return err
	}
	receiver.CategoryEntity = categoryEntity
	fmt.Println(categoryEntity)

	// カテゴリーのEntityの作成
	deliveryEntity, err := FindByDeliveryId(receiver.DeliveryEntityID)
	if err != nil {
		return err
	}
	receiver.DeliveryEntity = deliveryEntity
	// DBの作成
	fmt.Println("HERE")
	result := db.Create(receiver)
	if result.Error != nil {
		fmt.Println(result.Error.Error())
		err := service.MyError{Message: result.Error.Error()}
		return err
	}
	return nil
}

func CreateMerchandiseList(merchandises []MerchandiseEntity) error {
	result := db.Create(&merchandises)
	if result.Error != nil {
		myErr := service.MyError{
			Message: result.Error.Error(),
		}
		return myErr
	}
	return nil
}

func UpdateMerchandiseList(merchandises []MerchandiseEntity) error {
	result := db.Save(&merchandises)
	if result.Error != nil {
		myErr := service.MyError{
			Message: result.Error.Error(),
		}
		return myErr
	}
	return nil
}
