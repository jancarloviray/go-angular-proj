package routes

import (
	"../middlewares"
	"../models/task"
	"fmt"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
	"net/http"
)

func Init() {
	static()
	api()

	goji.Serve()
}

func static() {
	goji.Get("/", Root)
	goji.NotFound(NotFound)
}

func api() {
	api := web.New()
	api.Use(middlewares.Secure)
	api.Get("/api/tasks", task.GetTasks)
	api.Get("/api/tasks/:id", task.GetTask)
	api.Post("/api/tasks", task.CreateTask)

	// If last character is an asterisk,
	// the path is treated as a prefix
	// and can be used to implement
	// sub-routes which can have custom
	// middlewares.
	goji.Handle("/api/*", api)
}

// HANDLERS

func NotFound(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Custom Not Found Handler", 404)
}
