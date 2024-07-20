package pageio

import (
	"os"

	"github.com/jamiegyoung/runemarkers-go/internal/templating"
)

type renderable interface {
	Data() map[string]interface{}
}

func RenderHtml[T renderable](
	name string,
	page string,
	output *os.File,
	data T) error {
	templ, err := templating.HtmlTemplateWithComponents(name, page)
	if err != nil {
		return err
	}

	// create an interface containing both the page and the data
	log("Rendering " + name + " to " + output.Name())
	err = templ.ExecuteTemplate(output, name, data.Data())
	if err != nil {
		return err
	}

	return nil
}

func RenderText[T renderable](
	name string,
	page string,
	output *os.File,
	data T) error {
	templ, err := templating.TextTemplateWithComponents(name, page)
	if err != nil {
		return err
	}

	// create an interface containing both the page and the data
	log("Rendering " + name + " to " + output.Name())
	err = templ.ExecuteTemplate(output, name, data.Data())
	if err != nil {
		return err
	}

	return nil
}
