package main

import (
	"github.com/johskw/borda/router"
)

func main() {
	r := router.GetRouter()
	r.Run(":8080")
}
