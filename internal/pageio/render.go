package pageio

import (
	"os"

	"github.com/jamiegyoung/runemarkers-go/internal/templating"
)

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
