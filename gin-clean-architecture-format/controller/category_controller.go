package controller

import (
	"encoding/json"
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

func InsertBatchCategoriesHandler(c *gin.Context) {
	buf := make([]byte, 2048)
	// ここでRequest.Bodyを読み切る
	n, _ := c.Request.Body.Read(buf)

	// リクエストBodyの内容を保存する構造体
	var requestBody []repository.CategoryEntity
	err := json.Unmarshal(buf[0:n], &requestBody)
	if err != nil {
		c.JSON(500, err)
		return
	}
	err = repository.CreateCategoryList(requestBody)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, "Batch Insert Completed")
}
