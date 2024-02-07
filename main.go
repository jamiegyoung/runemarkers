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
	Title       string
	Body        string
	Entities    []*entities.Entity
	ButtonClass string
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
		Title:       "Test",
		Body:        "Test body",
		Entities:    entities,
		ButtonClass: "btn-primary",
	}

	for _, page_path := range pages_paths {
		wg.Add(1)

		go func(page_path string) {
			defer wg.Done()

			log("Reading " + page_path)
			page_bytes, err := os.ReadFile(page_path)
			if err != nil {
				panic(err)
			}

			page_string := string(page_bytes)

			// create the output directory if it doesn't exist
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
			defer out_file.Close()

			page_name := filepath.Base(page_path)

			templ, err := templating.TemplateWithComponents(page_name, page_string)
			if err != nil {
				panic(err)
			}

			log("Rendering " + page_name + " to " + out_file.Name())
			err = templ.ExecuteTemplate(out_file, page_name, pageData)
			if err != nil {
				panic(err)
			}
		}(page_path)
	}

	wg.Wait()
  log("Done")
}

func replaceTmplWithHtml(tmp string) string {
	return strings.Replace(tmp, ".tmpl", ".html", -1)
}
