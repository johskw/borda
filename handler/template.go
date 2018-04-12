package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/johskw/borda/service"
)

func ShowTop(c *gin.Context) {
	c.HTML(http.StatusOK, "top.html", nil)
}

func ShowSetting(c *gin.Context) {
	c.HTML(http.StatusOK, "setting.html", nil)
}

func ShowEvent(c *gin.Context) {
	event, err := service.GetEventFromStrID(c.Param("id"))
	if err != nil {
		log.Print(err)
	}
	c.HTML(http.StatusOK, "event.html", gin.H{
		"event": event,
	})
}
