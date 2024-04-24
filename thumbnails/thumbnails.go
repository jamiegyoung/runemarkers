package thumbnails

import (
	"github.com/jamiegyoung/runemarkers-go/entities"
	"github.com/jamiegyoung/runemarkers-go/logger"
)

var log = logger.New("thumbnails")

func Collect(ents []*entities.Entity, destination string, skip bool) {
	errs := make(chan error, 1)

	// check if --skip-thumbs is passed
	if !skip {
		go func() {
			log("Collecting thumbnails, use --skip-thumbs to skip collect thumbnails")
			errs <- entities.CollectThumbnails(ents, destination)
		}()
	} else {
		log("Skipping thumbnail collection, Updating thumbnail urls to safe url file names, assuming all are png")

		for _, entity := range ents {
			entity.Thumbnail = "thumbnails/" + entity.Uri + ".png"
		}

		close(errs)
	}

	if err := <-errs; err != nil {
		panic(err)
	}
}
