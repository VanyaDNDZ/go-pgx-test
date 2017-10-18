package web

import (
	"github.com/gin-gonic/gin"
	"cc-fileshare/web/controllers"
	)

func CreateApiV1(engine *gin.Engine) {
	fileController := new(controllers.FilesController)
	fs := engine.Group("/fileshare")
	fs.POST("/files", fileController.AddFile)
	fs.PUT("/files/:fileid/info", fileController.UpdateAttributes)
	fs.GET("/files/:fileid/info", fileController.GetFileInfo)
	fs.GET("/files/:fileid", fileController.GetFile)
	fs.POST("/files_search", fileController.SearchFile)
}
