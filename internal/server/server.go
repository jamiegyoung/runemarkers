package server

import (
	"log"
	"net/http"
)

func addHandler(mux *http.ServeMux, pattern string, handler http.Handler) {
	mux.Handle(pattern, errorMiddleware(handler))
}

func addHandlerFunc(mux *http.ServeMux, pattern string, next func(w http.ResponseWriter, r *http.Request)) {
	addHandler(mux, pattern, http.HandlerFunc(next))
}

func Start() {
	go func() {
		err := watch(rebuild)
		if err != nil {
			debug("A watcher error occured")
		}
	}()

	mux := http.NewServeMux()

	handleStatic(mux)

	debug("registering GET / handler")
	addHandlerFunc(mux, "GET /", handleIndex)

	debug("registering GET /{entity} handler")
	addHandlerFunc(mux, "GET /{entity}", handleEntity)

	s := &http.Server{
		Handler: mux,
		Addr:    ":8080",
	}

	debug("starting server")

	log.Fatal(s.ListenAndServe())

	debug("server listening at port 8080")
}
