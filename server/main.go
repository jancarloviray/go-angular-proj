package main

import (
	"./models"
	"./routes"
)

func main() {
	routes.Init()
	models.Setup()
}
