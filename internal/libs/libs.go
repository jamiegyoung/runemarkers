package libs

import (
	"github.com/jamiegyoung/runemarkers-go/internal/copier"
	"github.com/jamiegyoung/runemarkers-go/internal/logger"
)

var log = logger.New("libs")

func Copy(output string) error {
	log("Copying libraries")
	return copier.Copy("libs/*.js", output)
}
