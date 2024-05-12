package api

import (
	"os"

	"github.com/jamiegyoung/runemarkers-go/internal/entities"
	"github.com/jamiegyoung/runemarkers-go/internal/logger"
	"github.com/jamiegyoung/runemarkers-go/internal/pageio"
)

var log = logger.New("api")

func Generate(ents []*entities.Entity) error {
	return GenerateButtons(ents)
}

func GenerateButtons(ents []*entities.Entity) error {
	log("Generating buttons api")

	button, err := pageio.ReadPageString("templates/api/button.tmpl")
	if err != nil {
		return err
	}

	// create the api folder if it doesn't exist
	if _, err := os.Stat("public/api"); os.IsNotExist(err) {
		err = os.Mkdir("public/api", 0755)
		if err != nil {
			return err
		}
	}

	for _, entity := range ents {
		output, err := pageio.CreateOutFile(
			"public/api",
			"api/button_"+entity.ApiUri+".html",
		)
		if err != nil {
			return err
		}

		defer output.Close()

		data := ButtonPage{
			TilesString: entity.TilesString,
			Page: pageio.Page{
				ShowInfoButton: false,
			},
		}

		err = pageio.RenderPage(
			"button_"+entity.ApiUri,
			button,
			output,
			data,
		)
		if err != nil {
			return err
		}
	}
	return nil
}
