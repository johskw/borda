package main

import "github.com/johskw/borda/config"

func main() {
	r := config.GetRouter()
	r.LoadHTMLGlob("app/templates/*.html")
	r.Static("/src", "app/templates/src")
	r.Run(":8080")
}
