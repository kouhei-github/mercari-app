package controller

import (
	"github.com/gin-gonic/gin"
	"kouhei-github/sample-gin/repository"
)

func InsertAuthenticateHandler(c *gin.Context) {
	var requestBody repository.AuthenticationEntity
	// リクエストボディを構造体に格納
	err := c.ShouldBindJSON(&requestBody)
	if err != nil {
		c.JSON(500, err)
		return
	}

	// BlogEntity構造体の生成
	entity, err := repository.NewAuthenticationEntity(requestBody.Token, requestBody.Cookie)
	if err != nil {
		c.JSON(500, err)
		return
	}

	// 作成
	err = entity.Create()
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(201, "Insert Completed")
}

func UpdateAuthenticateHandler(c *gin.Context) {
	var requestBody repository.AuthenticationEntity
	// リクエストボディを構造体に格納
	err := c.ShouldBindJSON(&requestBody)
	if err != nil {
		c.JSON(500, err)
		return
	}
	// BlogEntity構造体の生成
	entity, err := repository.NewAuthenticationEntity(requestBody.Token, requestBody.Cookie)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	entity.ID = requestBody.ID
	// 更新
	err = entity.Update()
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(201, "Update Completed")
}
