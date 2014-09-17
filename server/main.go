package main

import (
	"./models"
	"./routes"

	g "github.com/zenazn/goji"
)

func main() {
	routes.Setup()
	models.Setup()
	server()
}

func server() {
	g.Serve()
}
