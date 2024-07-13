package templating

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/jamiegyoung/runemarkers-go/internal/logger"
)

var log = logger.New("templating")

var componentCache []string = nil
var styleCache string = ""

func countNewlines(s string) int {
	return strings.Count(s, "\n")
}

func truncateCharacters(s string, n int) string {
	if len(s) < n {
		return s
	}
	return s[:n-3] + "..."
}

func truncateLines(s string, n int) string {
	lines := strings.Split(s, "\n")
	if len(lines) < n {
		return s
	}
	return strings.Join(lines[:n-1], "\n") + "\n..."
}

func parseWithFuncs(tmpl *template.Template, text string) (*template.Template, error) {
	return tmpl.Funcs(template.FuncMap{
		"countNewlines":      countNewlines,
		"truncateLines":      truncateLines,
		"truncateCharacters": truncateCharacters,
	}).Parse(text)
}

func TemplateWithComponents(name string, text string) (*template.Template, error) {
	log("Generating template with components for " + name)
	templ, err := parseWithFuncs(template.New(name), text)

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
	templ, err = parseWithFuncs(templ, style)
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
		templ, err = parseWithFuncs(templ, component)
		if err != nil {
			log("Error parsing component string: \n" + component)
			return nil, err
		}
	}
	return templ, nil
}

func ClearCache() {
	componentCache = nil
	styleCache = ""
}

func createStyleComponent(fileStrings []string) string {
	log("Creating styles component")
	styles := strings.Join(fileStrings, "\n")
	component := fmt.Sprintf("{{define \"styles\"}}<style>%s</style>{{ end }}", styles)

	styleCache = component
	return component
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

	fileStrings := make([]string, len(files))

	var wg sync.WaitGroup

	errc := make(chan error, len(files))

	for i, file := range files {
		wg.Add(1)
		go func(filePath string, index int) {
			defer wg.Done()
			log("Collecting component style " + filePath)
			fileBytes, err := os.ReadFile(filePath)
			if err != nil {
				log(fmt.Sprintf("Error reading style file %v", filePath))
				errc <- err
			}
			fileStrings[index] = string(fileBytes)
		}(file, i)
	}

	wg.Wait()
	close(errc)

	for err := range errc {
		return "", err
	}

	return createStyleComponent(fileStrings), nil
}

func readComponents() ([]string, error) {
	// if componentCache != nil {
	// 	log("Using cached components")
	// 	return componentCache, nil
	// }

	files, err := filepath.Glob("templates/shared/*.tmpl")
	if err != nil {
		log("Error reading components")
		return nil, err
	}

	var wg sync.WaitGroup

	fileStrings := make([]string, len(files))

	errc := make(chan error, len(files))
	resc := make(chan string, len(files))

	for _, file := range files {
		wg.Add(1)

		go func(path string) {
			defer wg.Done()
			log("Collecting component " + path)

			bytes, err := os.ReadFile(path)
			if err != nil {
				log(fmt.Sprintf("Error reading file %v", path))
				errc <- err
				return
			}

			// remove the directory from the file name extension
			name := filepath.Base(path[:len(path)-len(filepath.Ext(path))])

			resc <- fmt.Sprintf(
				"{{ define \"%s\" }}%s{{ end }}",
				name,
				string(bytes),
			)
		}(file)
	}

	wg.Wait()
	close(errc)
	close(resc)

	for res := range resc {
		fileStrings = append(fileStrings, res)
	}

	for err := range errc {
		return fileStrings, err
	}

	// update cache
	componentCache = fileStrings

	return fileStrings, nil
}
