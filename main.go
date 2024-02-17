package main

import (
	"fmt"
	"github.com/jamiegyoung/runemarkers-go/entities"
	"github.com/jamiegyoung/runemarkers-go/libs"
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

const output_path = "public"
const pages_glob = "pages/*.tmpl"

func main() {

	pages_paths, err := filepath.Glob(pages_glob)

	if err != nil {
		panic(err)
	}

	log("Found " + fmt.Sprint(len(pages_paths)) + " page(s)")

	var wg sync.WaitGroup

	log("Reading entities")
	found_entities, err := entities.ReadAllEntities()
	if err != nil {
		panic(err)
	}

  errs := make(chan error, 1)
	go func() {
    errs <- entities.CollectThumbnails(found_entities, output_path)
  }()

	pageData := Page{
		Entities: found_entities,
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

	err = libs.CopyLibs(output_path + "/js")
	if err != nil {
		panic(err)
	}

  // listen for errors from the thumbnail collection
  if err := <-errs; err != nil {
    panic(err)
  }

	log("Done!")

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
