package repository

import (
	"gorm.io/gorm"
	"kouhei-github/sample-gin/service"
	"unicode/utf8"
)

type AuthenticationEntity struct {
	gorm.Model
	Token  string `json:"token" binding:"required"`
	Cookie string `json:"cookie"  binding:"required"`
}

func NewAuthenticationEntity(token string, cookie string) (*AuthenticationEntity, error) {
	if utf8.RuneCountInString(token) <= 2 {
		err := service.MyError{Message: "tokenを入力してください"}
		return &AuthenticationEntity{}, err
	}
	if utf8.RuneCountInString(cookie) <= 2 {
		err := service.MyError{Message: "Cookieを入力してください"}
		return &AuthenticationEntity{}, err
	}
	return &AuthenticationEntity{Token: token, Cookie: cookie}, nil
}

func (authentication *AuthenticationEntity) Create() error {
	result := db.Create(authentication)
	if result.Error != nil {
		err := service.MyError{Message: "モデルの作成ができませんでした。"}
		return err
	}
	return nil
}

func (authentication *AuthenticationEntity) Update() error {
	token := authentication.Token
	cookie := authentication.Cookie

	// レコードが存在するか確認
	result := db.First(authentication)
	if result.Error != nil {
		err := service.MyError{Message: result.Error.Error()}
		return err
	}
	// 存在したらUpdate
	authentication.Token = token
	authentication.Cookie = cookie

	// 更新した構造体で保存する
	result = db.Save(authentication)
	if result.Error != nil {
		err := service.MyError{Message: result.Error.Error()}
		return err
	}
	return nil
}

func (authentication *AuthenticationEntity) FindByLatest() error {
	result := db.First(authentication)
	if result.Error != nil {
		myErr := service.MyError{Message: result.Error.Error()}
		return myErr
	}
	return nil
}
