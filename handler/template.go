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
	if !event.Finished {
		c.HTML(http.StatusOK, "event.html", gin.H{
			"event": event,
		})
		return
	}
	result, err := service.GetResult(event)
	if err != nil {
		log.Print(err)
	}
	c.HTML(http.StatusOK, "event.html", gin.H{
		"event":  event,
		"result": result,
	})
}

func ShowEventEdit(c *gin.Context) {
	event, err := service.GetEventFromStrID(c.Param("id"))
	if err != nil {
		log.Print(err)
	}
	ok := event.CheckPassword(c.PostForm("password"))
	if !ok {
		c.Redirect(http.StatusMovedPermanently, "/event/"+event.StrID())
		return
	}
	c.HTML(http.StatusOK, "edit.html", gin.H{
		"event": event,
	})
}

func ShowVote(c *gin.Context) {
	event, err := service.GetEventFromStrID(c.Param("event_id"))
	if err != nil {
		log.Print(err)
	}
	c.HTML(http.StatusOK, "vote.html", gin.H{
		"event": event,
	})
}
