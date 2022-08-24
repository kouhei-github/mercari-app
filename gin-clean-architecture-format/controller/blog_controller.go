package controller

import (
	"github.com/gin-gonic/gin"
	"kouhei-github/sample-gin/repository"
)

func ShowHelloWorld(c *gin.Context) {
	c.JSON(200, "Hello Gin App on Docker.")
}

func InsertBlog(c *gin.Context) {
	var requestBody repository.BlogEntity
	// リクエストボディを構造体に格納
	err := c.ShouldBindJSON(&requestBody)
	if err != nil {
		c.JSON(500, err)
		return
	}
	// BlogEntity構造体の生成
	entity, err := repository.NewBlogEntity(requestBody.Title, requestBody.Body)
	if err != nil {
		c.JSON(500, err)
		return
	}

	// タイトルが既に存在しないか確認
	entities, err := entity.FindByTitle()

	if err != nil {
		c.JSON(500, err)
		return
	}
	if len(entities) > 0 {
		c.JSON(500, "タイトルが既に存在してます")
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
