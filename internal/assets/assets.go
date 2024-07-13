package assets

import (
	"github.com/jamiegyoung/runemarkers-go/internal/copier"
	"github.com/jamiegyoung/runemarkers-go/internal/logger"
)

var log = logger.New("assets")

func Copy(output string) error {
	log("Copying assets")
	err := copier.Copy("assets/**", output)
	// if err != nil {
	return err
	// }
	// return copier.Copy("assets/**", output)
}
