package templating

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"sync"

	"github.com/jamiegyoung/runemarkers-go/logger"
)

var log = logger.Logger("templating")

func TemplateWithComponents(name string, text string) (*template.Template, error) {
	log("Generating template with components for " + name)
	templ, err := template.New(name).Parse(text)
	if err != nil {
		return nil, err
	}

	components, err := readComponents()
	if err != nil {
		return nil, err
	}

	// Load all components and return the template
	for _, component := range components {
		templ, err = templ.Parse(component)
		if err != nil {
			return nil, err
		}
	}

	return templ, nil
}

func readComponents() ([]string, error) {
	files, err := filepath.Glob("components/*.tmpl")
	if err != nil {
		return nil, err
	}

	var wg sync.WaitGroup

	file_strings := make([]string, len(files))

	for _, file := range files {
		wg.Add(1)

		go func(file string) {
			defer wg.Done()
			file_bytes, err := os.ReadFile(file)
			if err != nil {
				panic(err)
			}

			// remove the directory from the file name extension
			file_name := filepath.Base(file[:len(file)-len(filepath.Ext(file))])

			log("Collected component " + file_name)

			file_strings = append(file_strings, fmt.Sprintf("{{ define \"%s\" }}%s{{ end }}", file_name, string(file_bytes)))
		}(file)
	}

	wg.Wait()

	return file_strings, nil
}
