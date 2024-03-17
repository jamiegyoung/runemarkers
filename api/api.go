package api

import (
	"os"

	"github.com/jamiegyoung/runemarkers-go/entities"
	"github.com/jamiegyoung/runemarkers-go/logger"
	"github.com/jamiegyoung/runemarkers-go/pageio"
)

var log = logger.New("api")

func Generate(found_entities []*entities.Entity) {
	GenerateButtons(found_entities)
}

func GenerateButtons(found_entities []*entities.Entity) {
	log("Generating buttons api")

	button_string, err := pageio.ReadPageString("./api/button.tmpl")
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

	for _, entity := range found_entities {
		out_file := pageio.CreateOutFile(
			"public/api",
			"api/button_"+entity.URI+".html",
		)

		pageio.RenderPage(
			"button_"+entity.URI,
			button_string,
			out_file,
			entity.TilesString,
		)
	}
}
