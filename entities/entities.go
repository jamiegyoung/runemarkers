package entities

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/jamiegyoung/runemarkers-go/logger"
)

// Tile colour can be multiple types and we don't really care which one it is
// so we'll just use the empty interface here
type Tile struct {
	RegionId int         `json:"regionId"`
	RegionX  int         `json:"regionX"`
	RegionY  int         `json:"regionY"`
	Z        int         `json:"z"`
	Color    interface{} `json:"color"`
	Label    string      `json:"label,omitempty"`
}

type Source struct {
	Link     string `json:"link"`
	Name     string `json:"name"`
	Modified string `json:"modified,omitempty"`
}

type Entity struct {
	SafeURI                 string   `json:"safeURI"`
	Name                    string   `json:"name"`
	Subcategory             string   `json:"subcategory,omitempty"`
	AltName                 string   `json:"altName,omitempty"`
	Tags                    []string `json:"tags"`
	Tiles                   []Tile   `json:"tiles"`
	Thumbnail               string   `json:"thumbnail"`
	Wiki                    string   `json:"wiki"`
	Source                  *Source  `json:"source,omitempty"`
	RecommendedGuideVideoId string   `json:"recommendedGuideVideoId,omitempty"`
	FullName                string   `json:"fullName"`
	FullAltName             string   `json:"fullAltName"`
}

var log = logger.Logger("entity")

func urlEncode(s string) string {
	lowered := strings.ToLower(s)
	unspacedAndLowered := strings.ReplaceAll(lowered, " ", "-")
	return url.QueryEscape(unspacedAndLowered)
}

func parseName(file string) string {
	if !strings.HasSuffix(file, ".json") {
		return filepath.Base(file)
	}
	return filepath.Base(file[:len(file)-len(filepath.Ext(file))])
}

func ReadAllEntities() ([]*Entity, error) {
	files, err := filepath.Glob("entities/*.json")
	if err != nil {
		return nil, err
	}

	log("Found " + fmt.Sprint(len(files)) + " entity file(s)")

	var wg sync.WaitGroup
	entities := make([]*Entity, len(files))

	for i, file := range files {
		wg.Add(1)
		go func(i int, file_path string) {
			defer wg.Done()
			entity_name := parseName(file_path)
			entity, err := ReadEntityAndParse(entity_name)
			if err != nil {
				panic(err)
			}
			log("Read " + file_path)
			entities[i] = entity
		}(i, file)
	}

	wg.Wait()

	return entities, nil
}

func CollectThumbnails(entities []*Entity, output_path string) error {
	for index, entity := range entities {
		log("(" + fmt.Sprint(index+1) + "/" + fmt.Sprint(len(entities)) + ") Collecting thumbnail for " + entity.Name)

		thumbnail_url := entity.Thumbnail

		response, err := http.Get(thumbnail_url)
		if err != nil {
			log("Error getting thumbnail: " + err.Error())
			return err
		}
		defer response.Body.Close()

		// get thumbnail file type from the end of the thumbnail_url
		thumbnail_file_type := filepath.Ext(thumbnail_url)

		// create directory if it doesn't exist
		thumbnail_output_path := output_path + "/thumbnails"

		err = os.MkdirAll(thumbnail_output_path, 0755)
		if err != nil {
			return err
		}

		file, err := os.Create(thumbnail_output_path + "/" + entity.SafeURI + thumbnail_file_type)
		if err != nil {
			return err
		}
		defer file.Close()

		_, err = io.Copy(file, response.Body)
		if err != nil {
			return err
		}

		log("Collected thumbnail for " + entity.Name)

		if index < len(entities)-1 {
			// sleep if not the last entity to prevent spamming the server
			time.Sleep(time.Millisecond * 200)
		}
	}

	return nil
}

func ReadEntityAndParse(name string) (*Entity, error) {
	// only add json if it doesn't exist
	data, err := os.ReadFile("entities/" + name + ".json")
	if err != nil {
		return nil, err
	}

	return parseEntity(data)
}

func transformEntity(entity *Entity) {
	if entity.Subcategory == "" {
		entity.SafeURI = urlEncode(entity.Name)
		return
	}

	entity.SafeURI = urlEncode(
		fmt.Sprintf("%s (%s)", entity.Name, entity.Subcategory),
	)
}

func parseEntity(data []byte) (*Entity, error) {
	var target *Entity
	err := json.Unmarshal(data, &target)

	transformEntity(target)

	return target, err
}
