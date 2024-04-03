package entities

import (
	"crypto/md5"
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
	Uri                     string
	SafeUri                 string
	ApiUri                  string
	SafeApiUri              string
	Name                    string   `json:"name"`
	Subcategory             string   `json:"subcategory,omitempty"`
	AltName                 string   `json:"altName,omitempty"`
	Tags                    []string `json:"tags"`
	Tiles                   []Tile   `json:"tiles"`
  TilesString             string
	Thumbnail               string   `json:"thumbnail"`
	Wiki                    string   `json:"wiki"`
	Source                  *Source  `json:"source,omitempty"`
	RecommendedGuideVideoId string   `json:"recommendedGuideVideoId,omitempty"`
	FullName                string
	FullAltName             string
}

var log = logger.New("entity")

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
			log("Reading " + file_path)

			entity_name := parseName(file_path)
			entity, err := ReadEntityAndParse(entity_name)

			if err != nil {
				panic(err)
			}
			entities[i] = entity
		}(i, file)
	}

	wg.Wait()

	return entities, nil
}

func CollectThumbnails(entities []*Entity, output_path string) error {
	// create directory if it doesn't exist
	thumbnail_output_path := output_path + "/thumbnails"

	// remove previous files in directory
	files, err := filepath.Glob(thumbnail_output_path + "/*")
	if err != nil {
		return err
	}

	log("Removing previous thumbnails")
	for _, file := range files {
		err = os.Remove(file)
		if err != nil {
			return err
		}
	}

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

		err = os.MkdirAll(thumbnail_output_path, 0755)
		if err != nil {
			return err
		}

		unescaped, err := url.QueryUnescape(entity.Uri)
		if err != nil {
			return err
		}

		file, err := os.Create(thumbnail_output_path + "/" + unescaped + thumbnail_file_type)
		if err != nil {
			return err
		}
		defer file.Close()

		_, err = io.Copy(file, response.Body)
		if err != nil {
			return err
		}

		if index < len(entities)-1 {
			// sleep if not the last entity to prevent spamming the server
			time.Sleep(time.Millisecond * 200)
		}

		// update the thumbnail to the new path
		entities[index].Thumbnail = "thumbnails/" + entity.Uri + thumbnail_file_type
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

func transformToUrl(s string, tilesString string) string {
	lowered := strings.ToLower(s)
	return strings.ReplaceAll(lowered, " ", "-")
}

func parseName(file string) string {
	if !strings.HasSuffix(file, ".json") {
		return filepath.Base(file)
	}
	return filepath.Base(file[:len(file)-len(filepath.Ext(file))])
}

func getEntityUri(entity Entity) string {
	if entity.Subcategory == "" {
		return transformToUrl(entity.Name, entity.TilesString)
	}

	return transformToUrl(
		fmt.Sprintf("%s (%s)", entity.Name, entity.Subcategory),
		entity.TilesString,
	)
}

func entityTilesHash(tilesString string) string {
	// generate a hash based on the entity tiles
	hash := md5.New()
	hash.Write([]byte(tilesString))
	// truncate hash to 8 characters
	return fmt.Sprintf("%x", hash.Sum(nil))[:8]
}

func transformEntity(entity *Entity) {
	tilesString, err := json.Marshal(entity.Tiles)
	if err != nil {
		panic(err)
	}

	entity.TilesString = string(tilesString)

	entity.FullName = fmt.Sprintf("%s %s", entity.Name, entity.Subcategory)
	entity.FullAltName = fmt.Sprintf("%s %s", entity.AltName, entity.Subcategory)

	hash := entityTilesHash(entity.TilesString)

	entity.Uri = getEntityUri(*entity)
	entity.SafeUri = url.QueryEscape(entity.Uri)

	entity.ApiUri = entity.Uri + "-" + hash
	entity.SafeApiUri = url.QueryEscape(entity.ApiUri)
}

func parseEntity(data []byte) (*Entity, error) {
	var target *Entity
	err := json.Unmarshal(data, &target)

	transformEntity(target)

	return target, err
}
