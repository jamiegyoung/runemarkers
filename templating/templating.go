package templating

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/jamiegyoung/runemarkers-go/logger"
)

var log = logger.New("templating")

func TemplateWithComponents(name string, text string) (*template.Template, error) {
	log("Generating template with components for " + name)
	templ, err := template.New(name).Parse(text)
	if err != nil {
		return nil, err
	}

	style, err := readComponentStyles()
	if err != nil {
		return nil, err
	}

	log("Parsing colleted styles")
	templ, err = templ.Parse(style)
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

func readComponentStyles() (string, error) {
	files, err := filepath.Glob("components/*.css")
	if err != nil {
		return "", err
	}

	var wg sync.WaitGroup

	file_strings := make([]string, len(files))

	for _, file := range files {
		wg.Add(1)
		go func(file_path string) {
			defer wg.Done()
			log("Collecting component style " + file_path)
			file_bytes, err := os.ReadFile(file_path)
			if err != nil {
				panic(err)
			}

			file_strings = append(
				file_strings,
				string(file_bytes),
			)
		}(file)
	}

	wg.Wait()

	log("Creating styles component")
	joined_styles := strings.Join(file_strings, "\n")
	styles_component := fmt.Sprintf("{{define \"styles\"}}<style>%s</style>{{ end }}", joined_styles)

	log("Styles component: " + styles_component)

	return styles_component, nil
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

		go func(file_path string) {
			defer wg.Done()
			log("Collecting component " + file_path)

			file_bytes, err := os.ReadFile(file_path)
			if err != nil {
				panic(err)
			}

			// remove the directory from the file name extension
			file_name := filepath.Base(file_path[:len(file_path)-len(filepath.Ext(file_path))])

			file_strings = append(
				file_strings,
				fmt.Sprintf("{{ define \"%s\" }}%s{{ end }}",
					file_name, string(file_bytes)),
			)
		}(file)
	}

	wg.Wait()

	return file_strings, nil
}
