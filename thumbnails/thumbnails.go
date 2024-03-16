package thumbnails

import (
	"os"

	"github.com/jamiegyoung/runemarkers-go/args"
	"github.com/jamiegyoung/runemarkers-go/entities"
	"github.com/jamiegyoung/runemarkers-go/logger"
)

var log = logger.New("thumbnails")

func Collect(found_entities []*entities.Entity, output_path string) {
	argsWithoutProg := os.Args[1:]

	errs := make(chan error, 1)

	// check if --skip-thumbs is passed
	if !args.HasArg(argsWithoutProg, "--skip-thumbs") {
		go func() {
			log("Collecting thumbnails, use --skip-thumbs to collect thumbnails")
			errs <- entities.CollectThumbnails(found_entities, output_path)
		}()
	} else {
		log("Skipping thumbnail collection, Updating thumbnail urls to safe url file names, assuming all are png")

		for _, entity := range found_entities {
			entity.Thumbnail = "thumbnails/" + entity.URI + ".png"
		}

		close(errs)
	}

	if err := <-errs; err != nil {
		panic(err)
	}
}
