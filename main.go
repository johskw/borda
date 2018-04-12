package main

import "github.com/johskw/borda/config"

func main() {
	r := config.GetRouter()
	r.Run(":8080")
}
