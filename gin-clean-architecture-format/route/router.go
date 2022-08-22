package route

import (
	"github.com/gin-gonic/gin"
	"kouhei-github/sample-gin/controller"
)

func GetRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/", controller.ShowHelloWorld)
	r.GET("/all", controller.ShowAllBlog)
	r.GET("/show/:id", controller.ShowOneBlog)
	r.GET("/create", controller.ShowCreateBlog)
	r.POST("/create", controller.CreateBlog)
	r.GET("/edit/:id", controller.ShowEditBlog)
	r.POST("/edit", controller.EditBlog)
	r.GET("/delete/:id", controller.ShowCheckDeleteBlog)
	r.POST("/delete", controller.DeleteBlog)
	return r
}
