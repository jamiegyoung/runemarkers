package main

import (
	"fmt"
	"github.com/jamiegyoung/runemarkers-go/entities"
	"github.com/jamiegyoung/runemarkers-go/logger"
	"github.com/jamiegyoung/runemarkers-go/templating"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

type Page struct {
	Entities []*entities.Entity
}

var log = logger.Logger("main")

func main() {
	output_path := "public"
	pages_paths, err := filepath.Glob("pages/*.tmpl")
	if err != nil {
		panic(err)
	}

	log("Found " + fmt.Sprint(len(pages_paths)) + " page(s)")

	var wg sync.WaitGroup

	log("Reading entities")
	entities, err := entities.ReadAllEntities()

	if err != nil {
		panic(err)
	}

	pageData := Page{
		Entities: entities,
	}

	for _, page_path := range pages_paths {
		wg.Add(1)

		go func(page_path string) {
			defer wg.Done()

			page_string, err := readPageString(page_path)
			if err != nil {
				panic(err)
			}

			// create the output directory if it doesn't exist
			out_file := createOutFile(output_path, page_path)
			defer out_file.Close()

			page_name := filepath.Base(page_path)

			renderPage(page_name, page_string, out_file, pageData)
		}(page_path)
	}

	wg.Wait()
	log("Done")
}

func renderPage(page_name string, page_string string, out_file *os.File, pageData Page) {
	templ, err := templating.TemplateWithComponents(page_name, page_string)
	if err != nil {
		panic(err)
	}

	log("Rendering " + page_name + " to " + out_file.Name())
	err = templ.ExecuteTemplate(out_file, page_name, pageData)
	if err != nil {
		panic(err)
	}
}

func readPageString(page_path string) (string, error) {
	log("Reading " + page_path)
	page_bytes, err := os.ReadFile(page_path)
	if err != nil {
		return "", err
	}

	return string(page_bytes), nil
}

func createOutFile(output_path string, page_path string) *os.File {
	if _, err := os.Stat(output_path); os.IsNotExist(err) {
		err := os.Mkdir(output_path, 0755)
		if err != nil {
			panic(err)
		}
	}

	out_file, err := os.Create(output_path + "/" + replaceTmplWithHtml(filepath.Base(page_path)))
	if err != nil {
		panic(err)
	}
	return out_file
}

func replaceTmplWithHtml(tmp string) string {
	return strings.Replace(tmp, ".tmpl", ".html", -1)
}
