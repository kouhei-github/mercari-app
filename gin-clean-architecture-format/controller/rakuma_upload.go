package controller

import (
	"github.com/gin-gonic/gin"
	"kouhei-github/sample-gin/usecase"
)

type MyStruct struct {
}

func (my *MyStruct) Error() string {
	return "ロギング"
}

func UploadToRakumaHandler(c *gin.Context) {
	// UseCase Interfaceを呼び出し
	var handle usecase.UseCase
	// Interfaceの実装
	handle = &usecase.UploadToRakumaUseCase{}
	err := handle.Handle()
	if err != nil {
		c.JSON(500, map[string]string{"message": err.Error()})
		return
	}
	c.JSON(200, map[string]string{"status": "Done"})
}
