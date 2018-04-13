package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/johskw/borda/model"
	"github.com/johskw/borda/service"
)

func CreateVote(c *gin.Context) {
	intID, _ := strconv.Atoi(c.Param("event_id"))
	vote := model.Vote{
		EventID: uint(intID),
	}
	err := c.Bind(&vote)
	if err != nil {
		log.Print(err)
	}
	_, err = service.CreateVoteAndScores(vote)
	if err != nil {
		log.Print(err)
	}
	c.Redirect(http.StatusMovedPermanently, "/event/"+c.Param("event_id"))
}
