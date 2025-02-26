package server

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/jamiegyoung/runemarkers/internal/hashing"
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
				// Sleep for a very short time, this fixes any files
				// missing issues
				time.Sleep(200 * time.Millisecond)

				buildError = action(event.Name)
				if buildError != nil {
					debug(fmt.Sprintf("A build error occured %v", buildError))
					return
				}

				// add the file back into the watchlist
				err = watcher.Add(event.Name)
				if err != nil {
					debug(fmt.Sprintf("error adding %v to watcher", event.Name))
					return
				}
			case err, ok := <-watcher.Errors:
				debug(fmt.Sprintf("error: %v", err))
				if !ok {
					return
				}
			}
		}
	}()

	for _, item := range watchlist {
		debug(fmt.Sprintf("initializing watcher for %v", item))
		err := watcher.Add(item)
		debug(fmt.Sprintf("watching %v", item))

		if err != nil {
			panic(err)
		}
	}

	<-make(chan struct{})
	debug("Watcher ended, this is probably an issue")
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
	ignores := []string{
		"README.md",
		"public/.*",
		"LICENSE",
		"\\.git/.*",
		"\\.github/.*",
		"tmp/.*",
		"\\.gitignore",
		"Dockerfile",
		"compose.yml",
		"lastmod.db",
	}
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

		hash, err := hashing.HashFile(filepath)
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
