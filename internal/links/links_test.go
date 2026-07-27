package links

import (
	"encoding/json"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/jamiegyoung/runemarkers/internal/entities"
)

const (
	userAgent      = "runemarkers-link-check (+https://github.com/jamiegyoung/runemarkers)"
	oembedEndpoint = "https://www.youtube.com/oembed"
)

var client = &http.Client{Timeout: 15 * time.Second}

func TestEntityLinks(t *testing.T) {
	paths, err := filepath.Glob("../../entities/*.json")
	if err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}

	if len(paths) == 0 {
		t.Fatal("Expected entity files, got none")
	}

	for _, path := range paths {
		t.Run(filepath.Base(path), func(t *testing.T) {
			data, err := os.ReadFile(path)
			if err != nil {
				t.Fatalf("Expected nil, got %v", err)
			}

			var ent entities.Entity
			if err := json.Unmarshal(data, &ent); err != nil {
				t.Fatalf("Expected nil, got %v", err)
			}

			checkLink(t, "thumbnail", ent.Thumbnail)
			checkLink(t, "wiki", ent.Wiki)
			checkLink(t, "source", ent.Source.Link)

			if ent.RecommendedGuideVideoId != "" {
				checkUrl(t, "recommendedGuideVideoId", oembed(ent.RecommendedGuideVideoId))
			}
		})
	}
}

// A youtube watch page answers 200 even when the video is gone, oembed doesn't
func checkLink(t *testing.T, field string, link string) {
	if link == "" {
		return
	}

	if id, isVideo := videoId(link); isVideo {
		checkUrl(t, field, oembed(id))
		return
	}

	checkUrl(t, field, link)
}

func checkUrl(t *testing.T, field string, link string) {
	// sleep between requests to prevent spamming the server
	defer time.Sleep(time.Millisecond * 200)

	req, err := http.NewRequest(http.MethodGet, link, nil)
	if err != nil {
		t.Errorf("%s %q is not a valid url: %v", field, link, err)
		return
	}

	req.Header.Set("User-Agent", userAgent)

	resp, err := client.Do(req)
	if err != nil {
		t.Errorf("%s %q could not be reached: %v", field, link, err)
		return
	}
	defer resp.Body.Close()

	// Youtube uses 403 for private videos, other hosts use it to block our ip
	if resp.StatusCode == http.StatusForbidden && !strings.HasPrefix(link, oembedEndpoint) {
		t.Logf("%s %q could not be checked, returned %s", field, link, resp.Status)
		return
	}

	if resp.StatusCode >= http.StatusBadRequest {
		t.Errorf("%s %q returned %s", field, link, resp.Status)
	}
}

func oembed(videoId string) string {
	query := url.Values{
		"url":    {"https://www.youtube.com/watch?v=" + videoId},
		"format": {"json"},
	}
	return oembedEndpoint + "?" + query.Encode()
}

func videoId(rawUrl string) (string, bool) {
	parsed, err := url.Parse(rawUrl)
	if err != nil {
		return "", false
	}

	host := strings.TrimPrefix(strings.ToLower(parsed.Hostname()), "www.")

	if host == "youtu.be" {
		id := strings.Trim(parsed.Path, "/")
		return id, id != ""
	}

	if host != "youtube.com" {
		return "", false
	}

	id := parsed.Query().Get("v")
	return id, id != ""
}
