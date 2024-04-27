package api

import (
	"os"

	"github.com/jamiegyoung/runemarkers-go/internal/entities"
	"github.com/jamiegyoung/runemarkers-go/internal/logger"
	"github.com/jamiegyoung/runemarkers-go/internal/pageio"
)

var log = logger.New("api")

func Generate(ents []*entities.Entity) {
	GenerateButtons(ents)
}

type ButtonPage struct {
	pageio.Page
	TilesString string
}

func (p ButtonPage) Data() map[string]interface{} {
	return map[string]interface{}{
		"Page":        &p.Page,
		"TilesString": p.TilesString,
	}
}

func GenerateButtons(ents []*entities.Entity) {
	log("Generating buttons api")

	button, err := pageio.ReadPageString("templates/api/button.tmpl")
	if err != nil {
		panic(err)
	}

	// create the api folder if it doesn't exist
	if _, err := os.Stat("public/api"); os.IsNotExist(err) {
		mkdirerr := os.Mkdir("public/api", 0755)
		if mkdirerr != nil {
			panic(mkdirerr)
		}
	}

	for _, entity := range ents {
		output := pageio.CreateOutFile(
			"public/api",
			"api/button_"+entity.ApiUri+".html",
		)
		defer output.Close()

		data := ButtonPage{
			TilesString: entity.TilesString,
			Page: pageio.Page{
				ShowInfoButton: false,
			},
		}

		pageio.RenderPage(
			"button_"+entity.ApiUri,
			button,
			output,
			data,
		)
	}
}
