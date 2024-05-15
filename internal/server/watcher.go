package server

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/fsnotify/fsnotify"
)

var fileHashes map[string]uint32 = make(map[string]uint32)

func watcher(watchlist []string, action func(string) error) error {
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
					buildError = action(event.Name)
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

// Takes a dir and provide all paths, ignoring any matched regex
func allFilesRecursive(dir string, ignores []string) ([]string, error) {
	files := []string{}

	err := filepath.WalkDir(dir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		for _, ignore := range ignores {
			match, err := regexp.MatchString(fmt.Sprintf("^%v$", ignore), path)
			if err != nil {
				return err
			}

			if match {
				return nil
			}
		}

		if d.IsDir() {
			return nil
		}

		debug("found file " + path)
		files = append(files, path)
		return nil
	})

	return files, err
}

func devFiles() ([]string, error) {
	ignores := []string{"README.md", "public/.*", "LICENSE", "\\.git/.*", "\\.github/.*", "tmp/.*", "\\.gitignore"}
	filepaths, err := allFilesRecursive("./", ignores)

	if err != nil {
		return nil, err
	}

	return filepaths, err
}

func watch(action func(string) error) error {
	filepaths, err := devFiles()
	if err != nil {
		return err
	}

	var watchlist []string

	for _, filepath := range filepaths {
		watchlist = append(watchlist, filepath)

		hash, err := newHash(filepath)
		if err != nil {
			return err
		}

		fileHashes[filepath] = hash
	}

	debug(fmt.Sprintf("watching %v", strings.Join(watchlist, ", ")))

	debug("starting watcher")
	err = watcher(watchlist, action)
	if err != nil {
		return err
	}

	return nil
}
