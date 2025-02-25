package server

import (
	"log"
	"net/http"
)

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

	debug("starting server at port 8080")
	log.Fatal(s.ListenAndServe())
	debug("server stopped")
}

func addHandlerFunc(mux *http.ServeMux, pattern string, next func(w http.ResponseWriter, r *http.Request)) {
	mux.Handle(pattern, errorMiddleware(http.HandlerFunc(next)))
}
