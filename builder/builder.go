package builder

import (
	"os"

	"github.com/jamiegyoung/runemarkers-go/api"
	"github.com/jamiegyoung/runemarkers-go/entities"
	"github.com/jamiegyoung/runemarkers-go/entitypages"
	"github.com/jamiegyoung/runemarkers-go/logger"
	"github.com/jamiegyoung/runemarkers-go/pages"
	"github.com/jamiegyoung/runemarkers-go/thumbnails"
)

const destination = "public"

var log = logger.New("build")

func Build(skipThumbs bool) {
	if _, err := os.Stat(destination); os.IsNotExist(err) {
		err := os.Mkdir(destination, 0755)
		if err != nil {
			panic(err)
		}
	}

	log("Reading entities")
	ents, err := entities.ReadAllEntities()
	if err != nil {
		panic(err)
	}

	api.Generate(ents)
	thumbnails.Collect(ents, destination, skipThumbs)
	pages.GeneratePages(destination, ents)
	entitypages.GeneratePages(destination, ents)
}
