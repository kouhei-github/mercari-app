package usecase

import (
	"fmt"
	"kouhei-github/sample-gin/repository"
	"kouhei-github/sample-gin/service"
	"os"
	"time"
)

type UploadToRakumaUseCase struct {
}

type UseCase interface {
	Handle() error
}

func (_self *UploadToRakumaUseCase) Handle() error {
	// 認証情報の取得
	auth := repository.AuthenticationEntity{}
	err := auth.FindByLatest()
	if err != nil {
		return err
	}

	// 現在のルートディレクトリの取得
	cur, err := os.Getwd()
	if err != nil {
		myErr := service.MyError{Message: err.Error()}
		return myErr
	}

	// []MerchandiseEntityの取得
	merchandiseEntities, err := repository.FindByUploadTarget("target")
	if err != nil {
		myErr := service.MyError{Message: err.Error()}
		return myErr
	}

	for _, merchandiseEntity := range merchandiseEntities {
		// 画像ファイルをURLから取得
		fileName, err := service.DownloadImage(merchandiseEntity.Image)
		if err != nil {
			myErr := service.MyError{Message: err.Error()}
			return myErr
		}
		// 画像のファイルパスの生成
		filePath := cur + "/" + fileName

		imageId, err := service.ImageToBrowserCache(filePath, auth.Token, auth.Cookie)
		if err != nil {
			myErr := service.MyError{Message: err.Error()}
			return myErr
		}
		fmt.Println(imageId)

		delivery, err := repository.FindByDeliveryId(merchandiseEntity.DeliveryEntityID)
		if err != nil {
			return err
		}
		category, err := repository.FindByCategoryId(merchandiseEntity.CategoryEntityID)
		if err != nil {
			return err
		}

		err = merchandiseEntity.ValidateMerchandiseBeforeUpload(
			imageId,
			auth.Token,
			auth.Cookie,
			category.CategoryRakumaId,
			delivery.Method,
			delivery.Date,
			delivery.Area,
		)
		if err != nil {
			return err
		}
		time.Sleep(2 * time.Second) // 2秒待つ
		err = merchandiseEntity.PostToRakuma(
			imageId,
			auth.Token,
			auth.Cookie,
			category.CategoryRakumaId,
			delivery.Method,
			delivery.Date,
			delivery.Area,
		)
		if err != nil {
			return err
		}
		var deleteInterface service.Delete
		deleteInterface = &service.Image{FileName: filePath}
		err = deleteInterface.Delete()
		if err != nil {
			return err
		}

		fmt.Println("DOne")
		//repository.Upload(repository.MerchandiseEntity{}, &MyStruct{})
	}
	return nil
}
