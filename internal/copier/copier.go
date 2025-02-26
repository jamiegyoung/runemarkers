package copier

import (
	"io"
	"os"
	"path/filepath"
	"sync"

	"github.com/jamiegyoung/runemarkers/internal/logger"
)

var log = logger.New("copier")

func Copy(input_glob string, output string) error {
	files, err := filepath.Glob(input_glob)
	if err != nil {
		return err
	}

	var wg sync.WaitGroup

	if os.IsNotExist(err) {
		err = os.MkdirAll(filepath.Dir(output), 0755)
		if err != nil {
			log("Error creating directory: " + output)
			return err
		}
	}

	errc := make(chan error)

	for _, path := range files {
		dest := output + "/" + filepath.Base(path)

		wg.Add(1)
		go func(path string, dest string) {
			defer wg.Done()

			log("Copying " + path + " to " + dest)

			// check if its a directory
			fileInfo, err := os.Stat(path)
			if err != nil {
				errc <- err
				return
			}

			if fileInfo.IsDir() {
				err = Copy(path+"/*", dest)
				if err != nil {
					errc <- err
				}
				return
			}

			err = os.MkdirAll(filepath.Dir(dest), 0755)
			if err != nil {
				log("Error creating directory: " + dest)
				errc <- err
				return
			}

			src, err := os.Open(path)
			if err != nil {
				log("Error opening file: " + dest)
				errc <- err
				return
			}
			defer src.Close()

			output, err := os.Create(dest)
			if err != nil {
				log("Error creating file: " + dest + err.Error())
				errc <- err
				return
			}
			defer output.Close()

			_, err = io.Copy(output, src)
			if err != nil {
				log("Error copying file: " + dest + err.Error())
				errc <- err
			}
		}(path, dest)
	}

	wg.Wait()
	close(errc)

	for err := range errc {
		return err
	}

	return nil

}
