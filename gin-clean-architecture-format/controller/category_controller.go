package controller

import (
	"github.com/gin-gonic/gin"
	"kouhei-github/sample-gin/repository"
)

func InsertCategoryHandler(c *gin.Context) {
	var requestBody repository.CategoryEntity
	// リクエストボディを構造体に格納
	err := c.ShouldBindJSON(&requestBody)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	// BlogEntity構造体の生成
	entity := repository.NewCategoryEntity(requestBody.Name, requestBody.CategoryRakumaId)
	if err != nil {
		c.JSON(500, err)
		return
	}

	err = entity.Create()
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(201, "Insert Completed")
}
