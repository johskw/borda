package config

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"

	"github.com/gin-gonic/gin"
	"github.com/johskw/borda/handler"
)

func GetRouter() *gin.Engine {
	r := gin.Default()
	r.Use(RequestLogger())

	r.GET("/", handler.ShowTop)
	r.GET("/setting", handler.ShowSetting)

	r.POST("/createevent", handler.CreateEvent)
	r.GET("/event/:id", handler.ShowEvent)
	r.POST("/event/:id/edit", handler.ShowEventEdit)
	r.POST("/updateevent/:id", handler.UpdateEvent)
	r.POST("/finishevent/:id", handler.FinishEvent)

	r.GET("/vote/:event_id", handler.ShowVote)
	r.POST("/vote/:event_id", handler.CreateVote)

	return r
}

func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		buf, _ := ioutil.ReadAll(c.Request.Body)
		rdr1 := ioutil.NopCloser(bytes.NewBuffer(buf))
		rdr2 := ioutil.NopCloser(bytes.NewBuffer(buf))
		fmt.Println(readBody(rdr1))
		c.Request.Body = rdr2
		c.Next()
	}
}

func readBody(reader io.Reader) string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(reader)
	s := buf.String()
	return s
}
