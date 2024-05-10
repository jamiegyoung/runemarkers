package server

import (
	"fmt"
	"net/http"
)

// modified by the watcher on rebuild
var buildError error = nil

func errorMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if buildError != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "500 Internal Server Error: %s", buildError)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func handleStatic(mux *http.ServeMux) {
	fs := http.FileServer(http.Dir("public"))

	debug("registering GET /api/ handler")
	mux.Handle("GET /api/", fs)
	debug("registering GET /css/ handler")
	mux.Handle("GET /css/", fs)
	debug("registering GET /js/ handler")
	mux.Handle("GET /js/", fs)
	debug("registering GET /thumbnails/ handler")
	mux.Handle("GET /thumbnails/", fs)
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	debug("serving public/index.html")
	http.ServeFile(w, r, "public/index.html")
}

func handleEntity(w http.ResponseWriter, r *http.Request) {
	entity := r.PathValue("entity")
	// assume invalid if errors
	valid := validateEntitiy(entity)

	if valid {
		debug("serving ./public/" + entity + ".html")
		http.ServeFile(w, r, "./public/"+entity+".html")
		return
	}

	debug("unsafe entity in " + r.URL.Path + ", serving 404")
	w.WriteHeader(404)
	fmt.Fprintf(w, "404 Entity not found")
	return
}
