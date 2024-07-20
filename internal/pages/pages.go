package pages

import (
	"fmt"
	"github.com/jamiegyoung/runemarkers-go/internal/entities"
	"github.com/jamiegyoung/runemarkers-go/internal/logger"
	"github.com/jamiegyoung/runemarkers-go/internal/pageio"
	"path/filepath"
	"sync"
)

var log = logger.New("pages")

type Page struct {
	data map[string]interface{}
}

func (p Page) Data() map[string]interface{} {
	return p.data
}

func NewPage(data map[string]interface{}) Page {
	return Page{
		data: data,
	}
}

func GeneratePages(destination string, entities []*entities.Entity) error {
	paths, err := filepath.Glob("templates/pages/*.tmpl")
	if err != nil {
		return err
	}

	log("Found " + fmt.Sprint(len(paths)) + " page(s)")

	data := NewPage(map[string]interface{}{"Entities": entities})

	var wg sync.WaitGroup
	errc := make(chan error, len(paths))

	for _, path := range paths {
		wg.Add(1)

		go func(path string) {
			defer wg.Done()

			page, err := pageio.ReadPageString(path)
			if err != nil {
				errc <- err
				return
			}

			// create the output directory if it doesn't exist
			output, err := pageio.CreateOutFile(destination, path)
			defer output.Close()
			if err != nil {
				errc <- err
				return
			}

			name := filepath.Base(path)

			err = pageio.RenderHtml(name, page, output, &data)
			errc <- err
		}(path)
	}

	wg.Wait()
	close(errc)

	for err := range errc {
		if err != nil {
			return err
		}
	}

	return nil
}
