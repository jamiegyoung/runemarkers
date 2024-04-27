package pages

import (
	"fmt"
	"github.com/jamiegyoung/runemarkers-go/internal/entities"
	"github.com/jamiegyoung/runemarkers-go/internal/libs"
	"github.com/jamiegyoung/runemarkers-go/internal/logger"
	"github.com/jamiegyoung/runemarkers-go/internal/pageio"
	"path/filepath"
	"sync"
)

var log = logger.New("pages")

type IndexPage struct {
	pageio.Page
	Entities []*entities.Entity
}

func (p *IndexPage) Data() map[string]interface{} {
	return map[string]interface{}{
		"Entities": p.Entities,
	}
}

func GeneratePages(destination string, entities []*entities.Entity) {
	paths, err := filepath.Glob("templates/pages/*.tmpl")
	if err != nil {
		panic(err)
	}

	log("Found " + fmt.Sprint(len(paths)) + " page(s)")

	data := IndexPage{
		Entities: entities,
	}

	var wg sync.WaitGroup

	for _, path := range paths {
		wg.Add(1)

		go func(path string) {
			defer wg.Done()

			page, err := pageio.ReadPageString(path)
			if err != nil {
				panic(err)
			}

			// create the output directory if it doesn't exist
			output := pageio.CreateOutFile(destination, path)
			defer output.Close()

			name := filepath.Base(path)

			pageio.RenderPage(name, page, output, &data)
		}(path)
	}

	wg.Wait()

	err = libs.CopyLibs(destination + "/js")
	if err != nil {
		panic(err)
	}

	log("Done!")
}
