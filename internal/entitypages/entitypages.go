package entitypages

import (
	"github.com/jamiegyoung/runemarkers-go/internal/entities"
	"github.com/jamiegyoung/runemarkers-go/internal/logger"
	"github.com/jamiegyoung/runemarkers-go/internal/pageio"
	"github.com/jamiegyoung/runemarkers-go/internal/pages"
)

var log = logger.New("entitypages")

func GeneratePages(destination string, foundEntities []*entities.Entity) error {
	log("Generating entity pages")

	path := "templates/entitypages/entity.tmpl"

	page, err := pageio.ReadPageString(path)
	if err != nil {
		return err
	}

	// for each of the entities, generate an entity page
	for _, entity := range foundEntities {
		log("Rendering " + entity.Name + " to " + destination + "/" + entity.Uri + ".html")
		output, err := pageio.CreateOutFile(destination, entity.Uri+".html")
		defer output.Close()

		if err != nil {
			return err
		}

		data := pages.NewPage(map[string]interface{}{"Entity": entity})

		err = pageio.RenderHtml(entity.Name, page, output, &data)
		if err != nil {
			return err
		}
	}
	return nil
}
