package router

import (
	"github.com/gin-gonic/gin"
	"github.com/johskw/borda/handler"
)

func GetRouter() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*.html")
	r.GET("/", handler.ShowTop)
	r.GET("/setting", handler.ShowSetting)
	return r
}
