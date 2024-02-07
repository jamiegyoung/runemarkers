package main

import (
	// "html/template"
	"os"
	"path/filepath"
	"sync"
	"text/template"
)

type Page struct {
	Title string
	Body  string
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

			page_file, err := os.Open(page_path)
			if err != nil {
				panic(err)
			}

			defer page_file.Close()

			pageData := Page{
				Title: "Test",
				Body:  "Test body",
			}

			// create the output directory if it doesn't exist
			if _, err := os.Stat(output_path); os.IsNotExist(err) {
				os.Mkdir(output_path, 0755)
			}

			out_file, err := os.Create(output_path + "/" + filepath.Base(page_path))
			if err != nil {
				panic(err)
			}
			defer out_file.Close()

			t, err := template.New("page").ParseFiles(page_path)
			if err != nil {
				panic(err)
			}

			err = t.ExecuteTemplate(out_file, filepath.Base(page_path), pageData)
		}(page_path)
	}

	wg.Wait()
}
