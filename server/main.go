package main

import (
	"./models"
	"./routes"

	"github.com/zenazn/goji"
)

func main() {
	routes.Init()
	models.Setup()
	server()
}

func server() {
	goji.Serve()
}
