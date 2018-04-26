package app

import (
	"net/http"

	"github.com/johskw/borda/config"
)

func init() {
	r := config.GetRouter()
	r.LoadHTMLGlob("templates/*.html")
	r.Static("/src", "templates/src")
	http.Handle("/", r)
}
