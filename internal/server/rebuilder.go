package server

import (
	"fmt"

	"github.com/jamiegyoung/runemarkers-go/internal/builder"
)

func rebuild(path string) {
	hash, err := NewHash(path)
	if err != nil {
		panic(err)
	}

	if hashes[path] != hash {
		debug(fmt.Sprintf("modified file: %v, rebuilding", path))
		hashes[path] = hash
		builder.Build(true)
	}
}
