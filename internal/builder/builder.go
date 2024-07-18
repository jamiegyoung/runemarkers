package builder

import (
	"os"

	"github.com/jamiegyoung/runemarkers-go/internal/api"
	"github.com/jamiegyoung/runemarkers-go/internal/assets"
	"github.com/jamiegyoung/runemarkers-go/internal/entities"
	"github.com/jamiegyoung/runemarkers-go/internal/entitypages"
	"github.com/jamiegyoung/runemarkers-go/internal/libs"
	"github.com/jamiegyoung/runemarkers-go/internal/logger"
	"github.com/jamiegyoung/runemarkers-go/internal/pages"
	"github.com/jamiegyoung/runemarkers-go/internal/robots"
	"github.com/jamiegyoung/runemarkers-go/internal/templating"
	"github.com/jamiegyoung/runemarkers-go/internal/thumbnails"
)

const destination = "public"

var log = logger.New("build")

func Build(skipThumbs bool) error {
	templating.ClearCache()

	if _, err := os.Stat(destination); os.IsNotExist(err) {
		err := os.Mkdir(destination, 0755)
		if err != nil {
			return err
		}
	}
	log("Reading entities")

	ents, err := entities.ReadAllEntities()
	if err != nil {
		return err
	}

	log("Generating entities")
	err = api.Generate(ents)
	if err != nil {
		return err
	}

	log("Collecting thumbnails")
	err = thumbnails.Collect(ents, destination, skipThumbs)
	if err != nil {
		return err
	}

	log("Generating pages")
	err = pages.GeneratePages(destination, ents)
	if err != nil {
		return err
	}

	err = entitypages.GeneratePages(destination, ents)
	if err != nil {
		return err
	}

	err = libs.Copy(destination + "/js")
	if err != nil {
		return err
	}

	err = assets.Copy(destination + "/assets")
	if err != nil {
		return err
	}

	err = robots.Copy(destination)
	if err != nil {
		return err
	}

	return nil
}
