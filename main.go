package main

import (
	"github.com/gin-gonic/gin"
	"cc-fileshare/web"
)

func main() {
	engine := gin.Default()
	web.CreateApiV1(engine)
	engine.Run() // listen and serve on 0.0.0.0:8080
}