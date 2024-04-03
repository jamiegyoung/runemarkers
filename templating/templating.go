package templating

import (
	"fmt"
	"github.com/jamiegyoung/runemarkers-go/logger"
	"html/template"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

var log = logger.New("templating")

var componentCache []string = nil
var styleCache string = ""

func TemplateWithComponents(name string, text string) (*template.Template, error) {
	log("Generating template with components for " + name)
	templ, err := template.New(name).Parse(text)
	if err != nil {
		log("Error parsing text")
		return nil, err
	}

	style, err := readComponentStyles()
	if err != nil {
		log("Error reading styles")
		return nil, err
	}

	log("Parsing colleted styles")
	templ, err = templ.Parse(style)
	if err != nil {
		log("Error parsing styles")
		return nil, err
	}

	components, err := readComponents()
	if err != nil {
		log("Error reading components")
		return nil, err
	}

	// Load all components and return the template
	for _, component := range components {
		templ, err = templ.Parse(component)
		if err != nil {
			log("Error parsing component string: \n" + component)
			return nil, err
		}
	}
	return templ, nil
}

func readComponentStyles() (string, error) {
	if styleCache != "" {
		log("Using cached styles")
		return styleCache, nil
	}

	files, err := filepath.Glob("components/*.css")
	if err != nil {
		log("Error reading component styles")
		return "", err
	}

	var wg sync.WaitGroup

	file_strings := make([]string, len(files))

	for i, file := range files {
		wg.Add(1)
		go func(file_path string, index int) {
			defer wg.Done()
			log("Collecting component style " + file_path)
			file_bytes, err := os.ReadFile(file_path)
			if err != nil {
				log("Error reading file")
				panic(err)
			}

			file_strings[index] = string(file_bytes);
		}(file, i)
	}

	wg.Wait()

	log("Creating styles component")
	joined_styles := strings.Join(file_strings, "\n")
	styles_component := fmt.Sprintf("{{define \"styles\"}}<style>%s</style>{{ end }}", joined_styles)

	styleCache = styles_component

	return styles_component, nil
}

func readComponents() ([]string, error) {
	if componentCache != nil {
		log("Using cached components")
		return componentCache, nil
	}

	files, err := filepath.Glob("components/*.tmpl")
	if err != nil {
		log("Error reading components")
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
				log("Error reading file")
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

	componentCache = file_strings

	return file_strings, nil
}
