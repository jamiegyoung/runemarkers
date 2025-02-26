package robots

import (
	"github.com/jamiegyoung/runemarkers/internal/copier"
	"github.com/jamiegyoung/runemarkers/internal/logger"
)

var log = logger.New("robots.txt")

func Copy(output string) error {
	log("Copying libraries")
	return copier.Copy("./robots.txt", output)
}
