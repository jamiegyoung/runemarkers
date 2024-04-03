package pages

import (
	"fmt"
	"github.com/jamiegyoung/runemarkers-go/entities"
	"github.com/jamiegyoung/runemarkers-go/libs"
	"github.com/jamiegyoung/runemarkers-go/logger"
	"github.com/jamiegyoung/runemarkers-go/pageio"
	"path/filepath"
	"sync"
)

var log = logger.New("pages")

const pages_glob = "pages/*.tmpl"

type IndexPage struct {
	pageio.Page
	Entities []*entities.Entity
}

func (p *IndexPage) Data() map[string]interface{} {
	return map[string]interface{}{
		"Entities": p.Entities,
	}
}

func GeneratePages(output_path string, found_entities []*entities.Entity) {
	pages_paths, err := filepath.Glob(pages_glob)
	if err != nil {
		panic(err)
	}

	log("Found " + fmt.Sprint(len(pages_paths)) + " page(s)")

	page_data := IndexPage{
		Entities: found_entities,
	}

	var wg sync.WaitGroup

	for _, page_path := range pages_paths {
		wg.Add(1)

		go func(page_path string) {
			defer wg.Done()

			page_string, err := pageio.ReadPageString(page_path)
			if err != nil {
				panic(err)
			}

			// create the output directory if it doesn't exist
			out_file := pageio.CreateOutFile(output_path, page_path)
			defer out_file.Close()

			page_name := filepath.Base(page_path)

			pageio.RenderPage(page_name, page_string, out_file, &page_data)
		}(page_path)
	}

	wg.Wait()

	err = libs.CopyLibs(output_path + "/js")
	if err != nil {
		panic(err)
	}

	log("Done!")
}
