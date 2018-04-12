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
		log.Println(err)
	}
	c.Redirect(http.StatusMovedPermanently, "/event/"+event.StrID())
}
