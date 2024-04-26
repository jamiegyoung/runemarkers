package server

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/jamiegyoung/runemarkers-go/logger"
)

var debug = logger.New("server")

// Subtracts b from a
func subtractStrings(a string, b string) string {
	return a[:len(a)-len(b)]
}

func existingEntities() ([]string, error) {
	debug("looking for allowed entities")
	var files, err = os.ReadDir("public")
	if err != nil {
		return nil, err
	}

	var names []string

	for _, file := range files {
		filename := file.Name()
		ext := filepath.Ext(filename)
		name := subtractStrings(filename, ext)
		names = append(names, name)
	}

	debug("Allowed entities found: " + strings.Join(names, ", "))
	return names, nil
}

// I would rather store this in ram than read the directory
// every time a request is made
var allowedEntites, allowedEntitesErr = existingEntities()

func isSafeEntity(unsafeName string) bool {
	if allowedEntitesErr != nil {
		panic(allowedEntitesErr)
	}

	if slices.Contains(allowedEntites, unsafeName) {
		return true
	}

	return false
}

func Start() {
	go func() {
		err := startWatching()
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
