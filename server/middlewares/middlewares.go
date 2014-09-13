package middlewares

import (
	"../config"
	"encoding/base64"
	"net/http"
	"strings"
)

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
			if err != nil || !strings.HasSuffix(string(pass), config.C.Pass) {
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
