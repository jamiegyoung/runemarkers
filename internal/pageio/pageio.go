package pageio

import (
	"os"
	"path/filepath"
	"strings"
)

type Page struct {
	ShowInfoButton bool
	CardHidden     bool
}

func ReadPageString(path string) (string, error) {
	log("Reading " + path)
	bytes, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func CreateOutFile(destination string, pagePath string) *os.File {
	if _, err := os.Stat(destination); os.IsNotExist(err) {
		err := os.Mkdir(destination, 0755)
		if err != nil {
			panic(err)
		}
	}

	output, err := os.Create(
		destination + "/" + replaceTmplWithHtml(filepath.Base(pagePath)),
	)
	if err != nil {
		panic(err)
	}
	return output
}

func replaceTmplWithHtml(tmp string) string {
	return strings.Replace(tmp, ".tmpl", ".html", -1)
}
