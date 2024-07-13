package pageio

import (
	"os"
	"path/filepath"
	"strings"
)

func ReadPageString(path string) (string, error) {
	log("Reading " + path)
	bytes, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func CreateOutFile(destination string, pagePath string) (*os.File, error) {
	if _, err := os.Stat(destination); os.IsNotExist(err) {
		err := os.Mkdir(destination, 0755)
		if err != nil {
			return nil, err
		}
	}

	output, err := os.Create(
		destination + "/" + replaceTmplWithHtml(filepath.Base(pagePath)),
	)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func replaceTmplWithHtml(tmp string) string {
	return strings.Replace(tmp, ".tmpl", ".html", -1)
}
