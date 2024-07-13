package entitypages

import (
	"sync"

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

	var wg sync.WaitGroup
	errc := make(chan error)

	// for each of the entities, generate an entity page
	for _, entity := range foundEntities {
		wg.Add(1)
		go func(entity *entities.Entity) {
			defer wg.Done()
			log("Rendering " + entity.Name + " to " + destination + "/" + entity.Uri + ".html")
			output, err := pageio.CreateOutFile(destination, entity.Uri+".html")
			defer output.Close()
			if err != nil {
				errc <- err
				return
			}

			data := pages.NewPage(map[string]interface{}{"Entity": entity})

			err = pageio.RenderPage(entity.Name, page, output, &data)
			if err != nil {
				errc <- err
			}
		}(entity)
	}

	wg.Wait()
	close(errc)

	return <-errc
}
