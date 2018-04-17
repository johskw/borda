package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/johskw/borda/model"
	"github.com/johskw/borda/service"
)

func CreateEvent(c *gin.Context) {
	var event model.Event
	err := c.Bind(&event)
	if err != nil {
		log.Print(err)
	}
	event, err = service.CreateEventAndChoices(event)
	if err != nil {
		log.Print(err)
	}
	c.Redirect(http.StatusMovedPermanently, "/event/"+event.StrID())
}

func UpdateEvent(c *gin.Context) {
	event, err := service.GetEventFromStrID(c.Param("id"))
	if err != nil {
		log.Print(err)
	}
	err = c.Bind(&event)
	if err != nil {
		log.Print(err)
	}
	err = event.Update()
	if err != nil {
		log.Print(err)
	}
	c.Redirect(http.StatusMovedPermanently, "/event/"+c.Param("id"))
}

func FinishEvent(c *gin.Context) {
	err := service.FinishEvent(c.Param("id"))
	if err != nil {
		log.Print(err)
	}
	c.Redirect(http.StatusMovedPermanently, "/event/"+c.Param("id"))
}
