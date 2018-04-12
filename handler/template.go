package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowTop(c *gin.Context) {
	c.HTML(http.StatusOK, "top.html", nil)
}

func ShowSetting(c *gin.Context) {
	c.HTML(http.StatusOK, "setting.html", nil)
}
