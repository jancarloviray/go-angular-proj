package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	g "github.com/zenazn/goji"
	"github.com/zenazn/goji/web"

	"../config"
	"../middlewares"
	"../models/task"
)

func Setup() {
	static()
	api()
	err()
}

func static() {
	g.Get("/", http.FileServer(http.Dir(config.C.PublicPath)))

	static := web.New()
	static.Get("/styles/*", http.StripPrefix("/styles/", http.FileServer(http.Dir(config.C.PublicPath+"/styles"))))
	static.Get("/scripts/*", http.StripPrefix("/scripts/", http.FileServer(http.Dir(config.C.PublicPath+"/scripts"))))
	static.Get("/images/*", http.FileServer(http.Dir(config.C.PublicPath+"/images")))
	static.Get("/robots.txt", http.FileServer(http.Dir(config.C.PublicPath)))

	g.Handle("/scripts/*", static)
	g.Handle("/styles/*", static)
	g.Handle("/images/*", static)
}

func err() {
	// custom 404
	g.NotFound(notFound)
}

func api() {
	api := web.New()
	api.Use(middlewares.Secure)
	api.Use(middlewares.JSON)
	api.Get("/api/tasks", listTasks)
	api.Get("/api/tasks/:id", getTask)
	api.Post("/api/tasks", createTask)

	// If last character is an asterisk,
	// the path is treated as a prefix
	// and can be used to implement
	// sub-routes which can have custom
	// middlewares.
	g.Handle("/api/*", api)
}

// HANDLERS

func notFound(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Custom Not Found Handler", 404)
}

func listTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := task.List()
	if err != nil {
		panic(err)
	}
	res, err := json.Marshal(tasks)
	fmt.Fprint(w, string(res))
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
