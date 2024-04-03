package entitypages

import (
	"sync"

	"github.com/jamiegyoung/runemarkers-go/entities"
	"github.com/jamiegyoung/runemarkers-go/logger"
	"github.com/jamiegyoung/runemarkers-go/pageio"
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

func GeneratePages(output_path string, found_entities []*entities.Entity) {
	log("Generating entity pages")

	page_path := "entitypages/entity.tmpl"

	page_string, err := pageio.ReadPageString(page_path)
	if err != nil {
		panic(err)
	}

	var wg sync.WaitGroup

	// for each of the entities, generate an entity page
	for _, entity := range found_entities {
		wg.Add(1)
		go func(entity *entities.Entity) {
			defer wg.Done()
			log("Rendering " + entity.Name + " to " + output_path + "/" + entity.Uri + ".html")
			out_file := pageio.CreateOutFile(output_path, entity.Uri+".html")
			defer out_file.Close()

			page_data := EntityPage{
				Entity: entity,
			}

			pageio.RenderPage(entity.Name, page_string, out_file, &page_data)
		}(entity)
	}

	wg.Wait()

	// err = libs.CopyLibs(output_path + "/js")
	// if err != nil {
	// 	panic(err)
	// }

	log("Done!")
}
