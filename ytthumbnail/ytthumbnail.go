package ytthumbnail

import (
	"errors"
	"net/http"

	"github.com/jamiegyoung/runemarkers-go/internal/logger"
)

var log = logger.New("ytthumbnail")

func Get(videoId string) (string, error) {
	versions := []string{"sddefault", "hqdefault", "0", "mqdefault", "maxresdefault", "default"}
	for _, version := range versions {
		url := "https://i.ytimg.com/vi_webp/" + videoId + "/" + version + ".webp"
		log("Checking " + url)

		if exists(url) {
			return url, nil
		}
	}
	return "", errors.New("No thumbnail found")
}

func exists(url string) bool {
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != 200 {
		return false
	}
	return true
}
