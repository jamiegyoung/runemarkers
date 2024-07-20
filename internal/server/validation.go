package server

import (
	"os"
	"path/filepath"
	"slices"
	"strings"
)

// I would rather store this in ram than read the directory
// every time a request is made

func validateEntity(unsafeName string) bool {
	debug("Validating entity:" + unsafeName)
	allowedEntites, allowedEntitesErr := existingEntities()
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

	// append sitemap
	names = append(names, "sitemap.xml")

	debug("Allowed entities found: " + strings.Join(names, ", "))
	return names, nil
}
