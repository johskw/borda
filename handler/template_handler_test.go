package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func getGinRouter() *gin.Engine {
	gin.SetMode("test")
	r := gin.New()
	r.LoadHTMLGlob("../templates/*.html")
	return r
}

func TestShowTop(t *testing.T) {
	r := getGinRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	r.GET("/", ShowTop)
	r.ServeHTTP(w, req)
	if w.Code != 200 {
		t.Error("Status code is", w.Code)
	}
}

func TestShowSetting(t *testing.T) {
	r := getGinRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/setting", nil)
	r.GET("/setting", ShowSetting)
	r.ServeHTTP(w, req)
	if w.Code != 200 {
		t.Error("Status code is", w.Code)
	}
}
