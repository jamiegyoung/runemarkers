package server

import (
	"hash/crc32"
	"os"
)

type FileHashes map[string]uint32

var fileHashes FileHashes = make(FileHashes)

func NewHash(path string) (uint32, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return 0, err
	}

	return crc32.ChecksumIEEE(content), nil
}
