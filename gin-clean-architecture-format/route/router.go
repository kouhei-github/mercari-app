package route

import (
	"github.com/gin-gonic/gin"
	"kouhei-github/sample-gin/controller"
)

func GetRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/", controller.ShowHelloWorld)
	// ブログのテーブル
	r.POST("/api/v1/create-blog", controller.InsertBlog)

	// 認証のテーブル
	r.POST("/api/v1/auth", controller.InsertAuthenticateHandler)
	r.PUT("/api/v1/auth", controller.UpdateAuthenticateHandler)

	// カテゴリーのテーブル
	r.POST("/api/v1/category", controller.InsertCategoryHandler)
	r.POST("/api/v1/category-batch", controller.InsertBatchCategoriesHandler)

	// デレバリー項目の追加
	r.POST("/api/v1/delivery", controller.InsertDeliveryHandler)

	// 商品情報の追加
	r.POST("/api/v1/merchandise", controller.CreateMerchandiseHandler)
	r.POST("/api/v1/merchandise-batch", controller.BulkInsertMerchandiseHandler)
	r.PUT("/api/v1/merchandise-batch", controller.BulkUpdateMerchandiseHandler)

	// ラクマ UPLOAD
	r.POST("/api/v1/upload-to-rakuma", controller.UploadToRakumaHandler)
	//r.GET("/all", controller.ShowAllBlog)
	//r.GET("/show/:id", controller.ShowOneBlog)
	//r.GET("/create", controller.ShowCreateBlog)
	//r.POST("/create", controller.CreateBlog)
	//r.GET("/edit/:id", controller.ShowEditBlog)
	//r.POST("/edit", controller.EditBlog)
	//r.GET("/delete/:id", controller.ShowCheckDeleteBlog)
	//r.POST("/delete", controller.DeleteBlog)
	return r
}
