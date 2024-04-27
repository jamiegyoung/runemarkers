package server

import (
	"fmt"
	"net/http"
)

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
	if validateEntitiy(entity) {
		debug("serving ./public/" + entity + ".html")
		http.ServeFile(w, r, "./public/"+entity+".html")
		return
	}

	debug("unsafe entity in " + r.URL.Path + ", serving 404")
	w.WriteHeader(404)
	fmt.Fprintf(w, "404 Entity not found")
}
