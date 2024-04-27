package builder

import (
	"os"

	"github.com/jamiegyoung/runemarkers-go/internal/api"
	"github.com/jamiegyoung/runemarkers-go/internal/entities"
	"github.com/jamiegyoung/runemarkers-go/internal/entitypages"
	"github.com/jamiegyoung/runemarkers-go/internal/logger"
	"github.com/jamiegyoung/runemarkers-go/internal/pages"
	"github.com/jamiegyoung/runemarkers-go/internal/thumbnails"
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
