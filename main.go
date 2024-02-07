package main

import (
	"os"
	"path/filepath"
	"sync"
)

type Page struct {
	Title    string
	Body     string
	Entities []*Entity
}

func main() {
	output_path := "public"
	pages_paths, err := filepath.Glob("pages/*.html")

	if err != nil {
		panic(err)
	}

	var wg sync.WaitGroup

	for _, page_path := range pages_paths {
		wg.Add(1)

		go func(page_path string) {
			defer wg.Done()

			page_bytes, err := os.ReadFile(page_path)
			if err != nil {
				panic(err)
			}

			page_string := string(page_bytes)

			entities, err := ReadAllEntities()

			if err != nil {
				panic(err)
			}

			pageData := Page{
				Title:  "Test",
				Body:   "Test body",
				Entities: entities,
			}

			// create the output directory if it doesn't exist
			if _, err := os.Stat(output_path); os.IsNotExist(err) {
				err := os.Mkdir(output_path, 0755)
				if err != nil {
					panic(err)
				}
			}

			out_file, err := os.Create(output_path + "/" + filepath.Base(page_path))
			if err != nil {
				panic(err)
			}
			defer out_file.Close()

			page_name := filepath.Base(page_path)

			templ, err := TemplateWithComponents(page_name, page_string)
			if err != nil {
				panic(err)
			}

			err = templ.ExecuteTemplate(out_file, page_name, pageData)
			if err != nil {
				panic(err)
			}
		}(page_path)
	}

	wg.Wait()
}
