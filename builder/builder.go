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

const output_path = "public"

var log = logger.New("build")

func Build(skip_thumbnails bool) {
	if _, err := os.Stat(output_path); os.IsNotExist(err) {
		err := os.Mkdir(output_path, 0755)
		if err != nil {
			panic(err)
		}
	}

	log("Reading entities")
	found_entities, err := entities.ReadAllEntities()
	if err != nil {
		panic(err)
	}

	api.Generate(found_entities)
	thumbnails.Collect(found_entities, output_path, skip_thumbnails)
	pages.GeneratePages(output_path, found_entities)
	entitypages.GeneratePages(output_path, found_entities)
}
