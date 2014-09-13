package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"

	"github.com/zenazn/goji"     // framework
	"github.com/zenazn/goji/web" // toolkit
)

const password = "pass"

func main() {
	goji.Get("/", Root)

	goji.NotFound(NotFound)

	api := web.New()
	api.Use(Secure)
	api.Get("/api/tasks", ListTasks)
	api.Get("/api/tasks/:id", GetTask)
	api.Post("/api/tasks", CreateTask)

	// If last character is an asterisk, the path is treated as a prefix
	// and can be used to implement sub-routes which can have custom middlewares.
	goji.Handle("/api/*", api)

	goji.Serve()
}

// HANDLERS

func Root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!")
}

func ListTasks(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "list task")
}

func GetTask(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "task with id: %s", c.URLParams["id"])
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "task created")
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Custom Not Found Handler", 404)
}

// MIDDLEWARES

func PlainText(h http.Handler) http.Handler {
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(handler)
}

func Secure(h http.Handler) http.Handler {
	handler := func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization") // Basic dXJAKLRYHA
		if !strings.HasPrefix(auth, "Basic") || len(auth) == 0 {
			forceAuth(w)
			return
		}
		if len(auth) > 0 {
			pass, err := base64.StdEncoding.DecodeString(auth[6:])
			if err != nil || !strings.HasSuffix(string(pass), password) {
				forceAuth(w)
				return
			}
			h.ServeHTTP(w, r)
		}
	}
	return http.HandlerFunc(handler)
}

// HELPERS

func forceAuth(w http.ResponseWriter) {
	w.Header().Set("WWW-Authenticate", `Basic realm="TaskMgmt"`)
	w.WriteHeader(http.StatusUnauthorized)
	w.Write([]byte("Restricted!"))
}
