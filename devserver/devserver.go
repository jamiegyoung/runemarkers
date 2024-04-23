package devserver

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"slices"

	"github.com/jamiegyoung/runemarkers-go/logger"
)

var debug = logger.New("devserver")

// Subtracts b from a
func subtractStrings(a string, b string) string {
	return a[:len(a)-len(b)]
}

func allowedEntites() ([]string, error) {
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

	return names, nil
}

// I would rather store this in ram than read the directory
// every time a request is made
var allowed_entites, allowed_entites_err = allowedEntites()

func isSafeEntity(unsafe_name string) bool {
	if allowed_entites_err != nil {
		panic(allowed_entites_err)
	}

	if slices.Contains(allowed_entites, unsafe_name) {
		return true
	}

	return false
}

func Start() {

	mux := http.NewServeMux()

	_, err := allowedEntites()
	if err != nil {
		panic(err)
	}

	fs := http.FileServer(http.Dir("public"))

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "public/index.html")
	})

	mux.HandleFunc("GET /{entity}", func(w http.ResponseWriter, r *http.Request) {
		entity := r.PathValue("entity")
		if isSafeEntity(entity) {
			http.ServeFile(w, r, "./public/"+entity+".html")
			return
		}

		w.WriteHeader(404)
		fmt.Fprintf(w, "404 Entity not found")
	})

	mux.Handle("GET /api/", fs)
	mux.Handle("GET /css/", fs)
	mux.Handle("GET /js/", fs)
	mux.Handle("GET /thumbnails/", fs)

	s := &http.Server{
		Handler: mux,
		Addr:    ":8080",
	}

	log.Fatal(s.ListenAndServe())
}
