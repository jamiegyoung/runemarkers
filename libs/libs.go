package libs

import (
	"github.com/jamiegyoung/runemarkers-go/logger"
	"io"
	"os"
	"path/filepath"
	"sync"
)

const libs_glob = "libs/*.js"

var log = logger.New("libs")

func CopyLibs(output_path string) error {
	file_paths, err := filepath.Glob(libs_glob)
	if err != nil {
		return err
	}

	err = os.MkdirAll(output_path, 0755)
	if err != nil {
		return err
	}

	var wg sync.WaitGroup

	for _, file_path := range file_paths {
		dest_path := output_path + "/" + filepath.Base(file_path)
		wg.Add(1)
		go func(file_path string, dest_path string) {
			defer wg.Done()

			src, err := os.Open(file_path)
			if err != nil {
				panic(err)
			}
			defer src.Close()

			dest_file, err := os.Create(dest_path)
			if err != nil {
				panic(err)
			}
			defer dest_file.Close()

			log("Copying " + file_path + " to " + dest_path)
			_, err = io.Copy(dest_file, src)
			if err != nil {
				panic(err)
			}
		}(file_path, dest_path)
	}

	wg.Wait()
	return nil

}
