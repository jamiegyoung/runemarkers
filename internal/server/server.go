package server

import (
	"log"
	"net/http"
)

func Start() {
	go func() {
		err := watch(rebuild)
		if err != nil {
			panic(err)
		}
	}()

	mux := http.NewServeMux()

	handleStatic(mux)

	debug("registering GET / handler")
	mux.HandleFunc("GET /", handleIndex)

	debug("registering GET /{entity} handler")
	mux.HandleFunc("GET /{entity}", handleEntity)

	s := &http.Server{
		Handler: mux,
		Addr:    ":8080",
	}

	debug("starting server")

	log.Fatal(s.ListenAndServe())

	debug("server listening at port 8080")
}
