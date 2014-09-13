package main

import (
	"./routes"
	"github.com/zenazn/goji"
)

const password = "pass"

func main() {
	routes.Init()
	goji.Serve()
}
