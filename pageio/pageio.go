package pageio

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/jamiegyoung/runemarkers-go/logger"
	"github.com/jamiegyoung/runemarkers-go/templating"
)

var log = logger.New("pageio")

type Page struct {
	ShowInfoButton bool
	CardHidden     bool
}

type renderable interface {
	Data() map[string]interface{}
}

func RenderPage[T renderable](
	name string,
	page string,
	output *os.File,
	data T) {
	templ, err := templating.TemplateWithComponents(name, page)
	if err != nil {
		panic(err)
	}

	// create an interface containing both the page and the data
	log("Rendering " + name + " to " + output.Name())
	err = templ.ExecuteTemplate(output, name, data.Data())
	if err != nil {
		panic(err)
	}
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
