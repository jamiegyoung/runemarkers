package thumbnails

import (
	"github.com/jamiegyoung/runemarkers-go/entities"
	"github.com/jamiegyoung/runemarkers-go/logger"
)

var log = logger.New("thumbnails")

func Collect(found_entities []*entities.Entity, output_path string, skip bool) {
	errs := make(chan error, 1)

	// check if --skip-thumbs is passed
	if !skip {
		go func() {
			log("Collecting thumbnails, use --skip-thumbs to skip collect thumbnails")
			errs <- entities.CollectThumbnails(found_entities, output_path)
		}()
	} else {
		log("Skipping thumbnail collection, Updating thumbnail urls to safe url file names, assuming all are png")

		for _, entity := range found_entities {
			entity.Thumbnail = "thumbnails/" + entity.Uri + ".png"
		}

		close(errs)
	}

	if err := <-errs; err != nil {
		panic(err)
	}
}
