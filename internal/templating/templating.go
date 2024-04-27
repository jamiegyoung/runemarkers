package templating

import (
	"fmt"
	"github.com/jamiegyoung/runemarkers-go/internal/logger"
	"html/template"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

var log = logger.New("templating")

var componentCache []string = nil
var styleCache string = ""

func ClearCache() {
	componentCache = nil
	styleCache = ""
}

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

	files, err := filepath.Glob("templates/shared/*.css")
	if err != nil {
		log("Error reading component styles")
		return "", err
	}

	var wg sync.WaitGroup

	fileStrings := make([]string, len(files))

	for i, file := range files {
		wg.Add(1)
		go func(filePath string, index int) {
			defer wg.Done()
			log("Collecting component style " + filePath)
			fileBytes, err := os.ReadFile(filePath)
			if err != nil {
				log("Error reading file")
				panic(err)
			}

			fileStrings[index] = string(fileBytes)
		}(file, i)
	}

	wg.Wait()

	log("Creating styles component")
	styles := strings.Join(fileStrings, "\n")
	component := fmt.Sprintf("{{define \"styles\"}}<style>%s</style>{{ end }}", styles)

	styleCache = component

	return component, nil
}

func readComponents() ([]string, error) {
	if componentCache != nil {
		log("Using cached components")
		return componentCache, nil
	}

	files, err := filepath.Glob("templates/shared/*.tmpl")
	if err != nil {
		log("Error reading components")
		return nil, err
	}

	var wg sync.WaitGroup

	fileStrings := make([]string, len(files))

	for _, file := range files {
		wg.Add(1)

		go func(path string) {
			defer wg.Done()
			log("Collecting component " + path)

			bytes, err := os.ReadFile(path)
			if err != nil {
				log("Error reading file")
				panic(err)
			}

			// remove the directory from the file name extension
			name := filepath.Base(path[:len(path)-len(filepath.Ext(path))])

			fileStrings = append(
				fileStrings,
				fmt.Sprintf("{{ define \"%s\" }}%s{{ end }}",
					name, string(bytes)),
			)
		}(file)
	}

	wg.Wait()

	componentCache = fileStrings

	return fileStrings, nil
}
