package controller

import (
	"github.com/gin-gonic/gin"
	"kouhei-github/sample-gin/repository"
	"kouhei-github/sample-gin/service"
)

func InsertDeliveryHandler(c *gin.Context) {
	var requestBody repository.DeliveryEntity
	err := c.ShouldBindJSON(&requestBody)
	if err != nil {
		myErr := service.MyError{Message: err.Error()}
		c.JSON(500, myErr)
		return
	}
	// DeliveryEntityの生成
	entity, err := repository.NewDeliveryEntity(requestBody.Method, requestBody.Date, requestBody.Area)
	if err != nil {
		c.JSON(500, err)
		return
	}
	// 保存の開始
	err = entity.Create()
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, "Insert Completed")
}
