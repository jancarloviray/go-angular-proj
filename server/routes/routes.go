package routes

import (
	m "../middlewares"
	"fmt"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
	"net/http"
)

func Init() {
	static()
	api()
}

func static() {
	goji.Get("/", handler)
	goji.NotFound(NotFound)
}

func api() {
	api := web.New()
	api.Use(m.Secure)
	api.Get("/api/tasks", GetTasks)
	api.Get("/api/tasks/:id", GetTask)
	api.Post("/api/tasks", CreateTask)

	// If last character is an asterisk, the path is treated as a prefix
	// and can be used to implement sub-routes which can have custom middlewares.
	goji.Handle("/api/*", api)
}

// HANDLERS

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "static!", r.URL)
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Custom Not Found Handler", 404)
}

func Root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!")
}

func GetTasks(w http.ResponseWriter, r *http.Request) {
	//task.GetTasks()
	fmt.Fprint(w, "list task")
}

func GetTask(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "task with id: %s", c.URLParams["id"])
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "task created")
}
