package server

import (
	"os"
	"path/filepath"
	"slices"
	"strings"
)

// I would rather store this in ram than read the directory
// every time a request is made
var allowedEntites, allowedEntitesErr = existingEntities()

func validateEntitiy(unsafeName string) bool {
	if allowedEntitesErr != nil {
		return false
	}

	if slices.Contains(allowedEntites, unsafeName) {
		return true
	}

	return false
}

// Subtracts b from a
func subtractStrings(a string, b string) string {
	return a[:len(a)-len(b)]
}

func existingEntities() ([]string, error) {
	debug("looking for allowed entities")
	var files, err = os.ReadDir("public")
	if err != nil {
		return nil, err
	}

	var names []string

	for _, file := range files {
		filename := file.Name()
		ext := filepath.Ext(filename)
		name := subtractStrings(filename, ext)
		names = append(names, name)
	}

	debug("Allowed entities found: " + strings.Join(names, ", "))
	return names, nil
}
