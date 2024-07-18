package server

import (
	"fmt"

	"github.com/jamiegyoung/runemarkers-go/internal/builder"
	"github.com/jamiegyoung/runemarkers-go/internal/hashing"
)

func rebuild(path string) error {
	hash, err := hashing.HashFile(path)
	if err != nil {
		debug("Unable to generate new hash, file likely doesn't exist anymore")
		return err
	}

	if fileHashes[path] != hash {
		debug(fmt.Sprintf("modified file: %v, rebuilding", path))
		fileHashes[path] = hash
		err := builder.Build(true)
		return err
	}

	return nil
}
