package server

import (
	"fmt"
	"path/filepath"
	"slices"

	"github.com/fsnotify/fsnotify"
	"github.com/jamiegyoung/runemarkers-go/builder"
)

func watcher(watchlist []string) error {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}
	defer watcher.Close()

	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Has(fsnotify.Write) {
					path := event.Name
					hash, err := NewHash(path)
					if err != nil {
						panic(err)
					}

					if fileHashes[path] != hash {
						debug(fmt.Sprintf("modified file: %v, rebuilding", path))
						fileHashes[path] = hash
						builder.Build(true)
					}

				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				debug(fmt.Sprintf("error: %v", err))
			}
		}
	}()

	for _, item := range watchlist {
		err := watcher.Add(item)
		if err != nil {
			panic(err)
		}
	}

	<-make(chan struct{})
	return nil
}

func devFiles() ([]string, error) {
	debug("getting dev files")

	allFilepaths, err := filepath.Glob("**/*")
	if err != nil {
		return nil, err
	}

	ignoreGlobs := []string{"README.md", "public/*", "LICENSE", ".git/*", ".github/*", "tmp/*"}

	var ignores []string

	for _, ignoreGlob := range ignoreGlobs {
		paths, err := filepath.Glob(ignoreGlob)
		if err != nil {
			return nil, err
		}
		ignores = append(ignores, paths...)
	}

	var filepaths []string

	for _, filepath := range allFilepaths {
		if slices.Contains(ignores, filepath) {
			continue
		}
		filepaths = append(filepaths, filepath)
	}

	return filepaths, err
}

func startWatching() error {
	filepaths, err := devFiles()
	if err != nil {
		return err
	}

	var watchlist []string

	for _, filepath := range filepaths {
		watchlist = append(watchlist, filepath)

		hash, err := NewHash(filepath)
		if err != nil {
			return err
		}

		fileHashes[filepath] = hash
	}

	debug("starting watcher")
	err = watcher(watchlist)
	if err != nil {
		return err
	}

	return nil
}
