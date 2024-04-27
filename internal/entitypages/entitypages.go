package entitypages

import (
	"sync"

	"github.com/jamiegyoung/runemarkers-go/internal/entities"
	"github.com/jamiegyoung/runemarkers-go/internal/logger"
	"github.com/jamiegyoung/runemarkers-go/internal/pageio"
)

var log = logger.New("entitypages")

type EntityPage struct {
	Entity *entities.Entity
}

func (p *EntityPage) Data() map[string]interface{} {
	return map[string]interface{}{
		"Entity": p.Entity,
	}
}

func GeneratePages(destination string, foundEntities []*entities.Entity) {
	log("Generating entity pages")

	path := "templates/entitypages/entity.tmpl"

	page, err := pageio.ReadPageString(path)
	if err != nil {
		panic(err)
	}

	var wg sync.WaitGroup

	// for each of the entities, generate an entity page
	for _, entity := range foundEntities {
		wg.Add(1)
		go func(entity *entities.Entity) {
			defer wg.Done()
			log("Rendering " + entity.Name + " to " + destination + "/" + entity.Uri + ".html")
			output := pageio.CreateOutFile(destination, entity.Uri+".html")
			defer output.Close()

			data := EntityPage{
				Entity: entity,
			}

			pageio.RenderPage(entity.Name, page, output, &data)
		}(entity)
	}

	wg.Wait()
	log("Done!")
}
