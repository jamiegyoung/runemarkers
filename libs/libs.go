package libs

import (
	"github.com/jamiegyoung/runemarkers-go/logger"
	"io"
	"os"
	"path/filepath"
	"sync"
)

var log = logger.New("libs")

func CopyLibs(output string) error {
	libs, err := filepath.Glob("libs/*.js")
	if err != nil {
		return err
	}

	err = os.MkdirAll(output, 0755)
	if err != nil {
		return err
	}

	var wg sync.WaitGroup

	for _, path := range libs {
		dest := output + "/" + filepath.Base(path)
		wg.Add(1)
		go func(path string, dest string) {
			defer wg.Done()

			src, err := os.Open(path)
			if err != nil {
				panic(err)
			}
			defer src.Close()

			output, err := os.Create(dest)
			if err != nil {
				panic(err)
			}
			defer output.Close()

			log("copying " + path + " to " + dest)
			_, err = io.Copy(output, src)
			if err != nil {
				panic(err)
			}
		}(path, dest)
	}

	wg.Wait()
	return nil

}
