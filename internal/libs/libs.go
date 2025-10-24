package libs

import (
	"github.com/jamiegyoung/runemarkers/internal/copier"
	"github.com/jamiegyoung/runemarkers/internal/logger"
	"path/filepath"
)

var log = logger.New("libs")

func Copy(output string) error {
	log("Copying libraries")

	err := copier.Copy("libs/*.js", output)
	if err != nil {
		return err
	}

	err = copier.Copy("libs/*.css", filepath.Dir(output)+"/css")
	if err != nil {
		return err
	}

	return copier.Copy("libs/images", output+"/images")
}
