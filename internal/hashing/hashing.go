package hashing

import (
	"hash/crc32"
	"os"
)

func HashFile(path string) (uint32, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return 0, err
	}

	return Hash(content), nil
}

func HashString(data string) uint32 {
	return Hash([]byte(data))
}

func Hash(data []byte) uint32 {
	return crc32.ChecksumIEEE(data)
}
