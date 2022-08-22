package controller

import (
	"fmt"
	"kouhei-github/sample-gin/repository"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ShowHelloWorld(c *gin.Context) {
	c.JSON(200, "Hello World !!")
}

func ShowAllBlog(c *gin.Context) {
	datas := repository.GetAll()
	c.JSON(200, datas)
}

func ShowOneBlog(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data := repository.GetOne(id)
	c.JSON(200, data)
}

func ShowCreateBlog(c *gin.Context) {
	c.JSON(200, "Done")
}

func CreateBlog(c *gin.Context) {
	title := c.PostForm("title")
	body := c.PostForm("body")
	data := repository.BlogEntity{Title: title, Body: body}
	data.Create()
	c.JSON(200, data)
}

func ShowEditBlog(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data := repository.GetOne(id)
	c.JSON(200, data)
}

func EditBlog(c *gin.Context) {
	id, _ := strconv.Atoi(c.PostForm("id"))
	data := repository.GetOne(id)
	title := c.PostForm("title")
	data.Title = title
	body := c.PostForm("body")
	data.Body = body
	data.Update()
	c.JSON(200, data)
}

func ShowCheckDeleteBlog(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data := repository.GetOne(id)
	c.JSON(200, data)
}

func DeleteBlog(c *gin.Context) {
	id, _ := strconv.Atoi(c.PostForm("id"))
	fmt.Println("delete:", id)
	data := repository.GetOne(id)
	data.Delete()
	c.JSON(200, data)
}
