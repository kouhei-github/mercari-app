package repository

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
	"io/ioutil"
	"kouhei-github/sample-gin/service"
	"mime/multipart"
	"net/http"
	"strconv"
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

func Upload(entity MerchandiseEntity, myError error) {
	fmt.Println(myError)
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

func FindByUploadTarget(search string) ([]MerchandiseEntity, error) {
	var merchandiseEntities []MerchandiseEntity
	var merchandiseEntity MerchandiseEntity
	result := db.Model(merchandiseEntity).Where("is_upload = ?", search).Find(&merchandiseEntities)
	if result.Error != nil {
		myErr := service.MyError{Message: result.Error.Error()}
		return []MerchandiseEntity{}, myErr
	}
	fmt.Println(merchandiseEntities[0])
	return merchandiseEntities, nil
}

type validateResponse struct {
	Result bool `json:"result"`
}

func (entity *MerchandiseEntity) ValidateMerchandiseBeforeUpload(
	imageId int,
	token string,
	cookie string,
	rakumaId int,
	deliveryMethod int,
	deliveryDate int,
	deliveryArea int,
) error {
	url := "https://fril.jp/item/validate"
	method := "POST"
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("authenticity_token", token)
	_ = writer.WriteField("utf8", "✓")
	_ = writer.WriteField("item_img_ids[]", strconv.Itoa(imageId))
	_ = writer.WriteField("item_img_ids[]", "")
	_ = writer.WriteField("item_img_ids[]", "")
	_ = writer.WriteField("item_img_ids[]", "")
	_ = writer.WriteField("item_img_ids[]", "")
	_ = writer.WriteField("item_img_ids[]", "")
	_ = writer.WriteField("item_img_ids[]", "")
	_ = writer.WriteField("item_img_ids[]", "")
	_ = writer.WriteField("item_img_ids[]", "")
	_ = writer.WriteField("item_img_ids[]", "")
	_ = writer.WriteField("item_img_ids[]", "")
	_ = writer.WriteField("item_img_ids[]", "")
	_ = writer.WriteField("item_img_ids[]", "")
	_ = writer.WriteField("item_img_ids[]", "")
	_ = writer.WriteField("item_img_ids[]", "")
	_ = writer.WriteField("item_img_ids[]", "")
	_ = writer.WriteField("item_img_ids[]", "")
	_ = writer.WriteField("item_img_ids[]", "")
	_ = writer.WriteField("item_img_ids[]", "")
	_ = writer.WriteField("item_img_ids[]", "")
	_ = writer.WriteField("item[category_id]", strconv.Itoa(rakumaId))
	_ = writer.WriteField("item[brand_id]", "")
	_ = writer.WriteField("belonging_hash_id", "")
	_ = writer.WriteField("item[name]", entity.Name)
	_ = writer.WriteField("item[detail]", entity.Detail)
	_ = writer.WriteField("item[status]", strconv.Itoa(entity.Status))
	_ = writer.WriteField("item[carriage]", strconv.Itoa(entity.Carriage))
	_ = writer.WriteField("item[delivery_method]", strconv.Itoa(deliveryMethod))
	_ = writer.WriteField("item[delivery_date]", strconv.Itoa(deliveryDate))
	_ = writer.WriteField("item[delivery_area]", strconv.Itoa(deliveryArea))
	_ = writer.WriteField("item[request_required]", "0")
	_ = writer.WriteField("item[sell_price]", strconv.Itoa(entity.SellPrice))
	err = writer.Close()
	if err != nil {
		fmt.Println("First")
		fmt.Println(err)
		return err
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println("Second")
		fmt.Println(err)
		return err
	}
	req.Header.Add("cookie", cookie)
	req.Header.Add("x-csrf-token", token)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Third")
		fmt.Println(err)
		return err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Fourth")
		fmt.Println(err)
		return err
	}

	var response validateResponse
	fmt.Println(string(body))
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println("Fifth")
		//fmt.Println(err)
		return err
	}
	fmt.Println(response)
	if !response.Result {
		err := service.MyError{Message: "商品登録時バリデーションに失敗しました"}
		return err
	}
	return nil
}

func (entity MerchandiseEntity) PostToRakuma(
	imageId int,
	token string,
	cookie string,
	rakumaId int,
	deliveryMethod int,
	deliveryDate int,
	deliveryArea int,
) error {
	url := "https://fril.jp/item/create_with_tmp_img"
	method := "POST"

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("authenticity_token", token)
	_ = writer.WriteField("utf8", "✓")
	_ = writer.WriteField("tmp_img_ids[]", strconv.Itoa(imageId))
	_ = writer.WriteField("item_img_ids[]", "")
	_ = writer.WriteField("tmp_img_ids[]", "")
	_ = writer.WriteField("item_img_ids[]", "")
	_ = writer.WriteField("tmp_img_ids[]", "")
	_ = writer.WriteField("item_img_ids[]", "")
	_ = writer.WriteField("tmp_img_ids[]", "")
	_ = writer.WriteField("item_img_ids[]", "")
	_ = writer.WriteField("tmp_img_ids[]", "")
	_ = writer.WriteField("item_img_ids[]", "")
	_ = writer.WriteField("tmp_img_ids[]", "")
	_ = writer.WriteField("item_img_ids[]", "")
	_ = writer.WriteField("tmp_img_ids[]", "")
	_ = writer.WriteField("item_img_ids[]", "")
	_ = writer.WriteField("tmp_img_ids[]", "")
	_ = writer.WriteField("item_img_ids[]", "")
	_ = writer.WriteField("tmp_img_ids[]", "")
	_ = writer.WriteField("item_img_ids[]", "")
	_ = writer.WriteField("item[category_id]", strconv.Itoa(rakumaId))
	_ = writer.WriteField("item[brand_id]", "")
	_ = writer.WriteField("belonging_hash_id", "")
	_ = writer.WriteField("item[name]", entity.Name)
	_ = writer.WriteField("item[detail]", entity.Detail)
	_ = writer.WriteField("item[status]", strconv.Itoa(entity.Status))
	_ = writer.WriteField("item[carriage]", strconv.Itoa(entity.Carriage))
	_ = writer.WriteField("item[delivery_method]", strconv.Itoa(deliveryMethod))
	_ = writer.WriteField("item[delivery_date]", strconv.Itoa(deliveryDate))
	_ = writer.WriteField("item[delivery_area]", strconv.Itoa(deliveryArea))
	_ = writer.WriteField("item[request_required]", "0")
	_ = writer.WriteField("item[sell_price]", strconv.Itoa(entity.SellPrice))
	err := writer.Close()
	if err != nil {
		fmt.Println(err)
		return err
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return err
	}
	req.Header.Add("cookie", cookie)
	req.Header.Add("x-csrf-token", token)

	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(string(body))

	var response validateResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println("Fifth")
		//fmt.Println(err)
		return err
	}
	fmt.Println(response)
	if !response.Result {
		err := service.MyError{Message: "商品登録に失敗しました"}
		return err
	}
	return nil
}
