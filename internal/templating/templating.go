package templating

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	textTemplate "text/template"

	"github.com/jamiegyoung/runemarkers/internal/logger"
)

var log = logger.New("templating")

var funcMap = textTemplate.FuncMap{
	"countNewlines":      countNewlines,
	"truncateLines":      truncateLines,
	"truncateCharacters": truncateCharacters,
}

func readComponents() ([]string, error) {
	files, err := filepath.Glob("templates/shared/*.tmpl")
	if err != nil {
		log("Error reading components")
		return nil, err
	}

	components := make([]string, len(files))
	for i, path := range files {
		log("Collecting component " + path)

		bytes, err := os.ReadFile(path)
		if err != nil {
			return nil, err
		}

		// remove the directory from the file name extension
		name := filepath.Base(path[:len(path)-len(filepath.Ext(path))])

		components[i] = createComponent(bytes, name)
	}
	return components, nil
}

func readComponentStyles() (string, error) {
	files, err := filepath.Glob("templates/shared/*.css")
	if err != nil {
		log("Error reading component styles")
		return "", err
	}

	parts := make([]string, len(files))

	for i, path := range files {
		log("Collecting component style " + path)
		bytes, err := os.ReadFile(path)
		if err != nil {
			return "", err
		}
		parts[i] = string(bytes)
	}

	return createStyleComponent(parts), nil
}

func createComponent(data []byte, name string) string {
	return fmt.Sprintf("{{ define \"%s\" }}%s{{ end }}", name, string(data))
}

func createStyleComponent(parts []string) string {
	styles := strings.Join(parts, "\n")
	component := fmt.Sprintf("{{define \"styles\"}}<style>%s</style>{{ end }}", styles)

	return component
}

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
