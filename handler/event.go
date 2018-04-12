package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/johskw/borda/model"
)

func CreateEvent(c *gin.Context) {
	var event model.Event
	err := c.Bind(&event)
	if err != nil {
		log.Print(err)
	}
	event, err = event.Create()
	if err != nil {
		log.Print(err)
	}
	for _, choiceName := range event.ChoiceNames {
		choice := model.Choice{
			Name:    choiceName,
			EventID: event.ID,
		}
		err = choice.Create()
		if err != nil {
			log.Print(err)
		}
	}
	c.Redirect(http.StatusMovedPermanently, "/process/"+event.StrID())
}
