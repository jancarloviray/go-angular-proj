package routes

import (
	"../config"
	"../middlewares"
	"../models/task"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
	"net/http"
)

func Init() {
	Static()
	API()

	goji.Serve()
}

func Static() {
	goji.Get("/", http.FileServer(http.Dir(config.C.PublicPath)))
	goji.Get("/robots.txt", http.FileServer(http.Dir(config.C.PublicPath)))
	goji.Get("/favicon.ico", http.FileServer(http.Dir(config.C.PublicPath+"/images")))

	// custom 404
	goji.NotFound(notFound)
}

func API() {
	api := web.New()
	api.Use(middlewares.Secure)
	api.Get("/api/tasks", listTasks)
	api.Get("/api/tasks/:id", getTask)
	api.Post("/api/tasks", createTask)

	// If last character is an asterisk,
	// the path is treated as a prefix
	// and can be used to implement
	// sub-routes which can have custom
	// middlewares.
	goji.Handle("/api/*", api)
}

// HANDLERS

func notFound(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Custom Not Found Handler", 404)
}

func listTasks(w http.ResponseWriter, r *http.Request) {
	task.List()
}

func getTask(c web.C, w http.ResponseWriter, r *http.Request) {
	id := c.URLParams["id"]
	task.Get(id)
}

func createTask(c web.C, w http.ResponseWriter, r *http.Request) {
	t := task.Task{}
	task.Create(t)
}

func updateTask(c web.C, w http.ResponseWriter, r *http.Request) {
	id := c.URLParams["id"]
	t := task.Task{}
	task.Update(id, t)
}
