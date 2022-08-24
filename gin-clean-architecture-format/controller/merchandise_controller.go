package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"kouhei-github/sample-gin/repository"
	"kouhei-github/sample-gin/service"
)

func CreateMerchandiseHandler(c *gin.Context) {
	var requestBody repository.MerchandiseEntity
	err := c.ShouldBindJSON(&requestBody)
	if err != nil {
		myErr := service.MyError{Message: err.Error()}
		c.JSON(500, myErr)
		return
	}
	entity, err := repository.NewMerchandiseEntity(
		requestBody.Image,
		requestBody.Name,
		requestBody.Detail,
		requestBody.Status,
		requestBody.Carriage,
		requestBody.RequestRequired,
		requestBody.SellPrice,
		requestBody.DeliveryEntityID,
		requestBody.CategoryEntityID,
	)
	if err != nil {
		c.JSON(500, err)
		return
	}

	// 作成する
	err = entity.Create()
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(201, "Insert Completed")
}

func BulkInsertMerchandiseHandler(c *gin.Context) {
	buf := make([]byte, 2048)
	// ここでRequest.Bodyを読み切る
	n, _ := c.Request.Body.Read(buf)

	// リクエストBodyの内容を保存する構造体
	var requestBody []repository.MerchandiseEntity
	err := json.Unmarshal(buf[0:n], &requestBody)
	if err != nil {
		c.JSON(500, err)
		return
	}
	err = repository.CreateMerchandiseList(requestBody)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, "Batch Insert Completed")
}

func BulkUpdateMerchandiseHandler(c *gin.Context) {
	buf := make([]byte, 2048)
	// ここでRequest.Bodyを読み切る
	n, _ := c.Request.Body.Read(buf)

	// リクエストBodyの内容を保存する構造体
	var requestBody []repository.MerchandiseEntity
	err := json.Unmarshal(buf[0:n], &requestBody)
	fmt.Println(requestBody[0].IsUpload)
	if err != nil {
		c.JSON(500, err)
		return
	}
	err = repository.UpdateMerchandiseList(requestBody)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, "Batch update Completed")
}
