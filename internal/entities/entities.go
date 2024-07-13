package entities

import (
	"crypto/md5"
	"encoding/base64"
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

	"github.com/jamiegyoung/runemarkers-go/internal/logger"
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
	MapLink                 string
	Thumbnail               string  `json:"thumbnail"`
	Wiki                    string  `json:"wiki"`
	Source                  *Source `json:"source,omitempty"`
	RecommendedGuideVideoId string  `json:"recommendedGuideVideoId,omitempty"`
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
	ents := make([]*Entity, len(files))
	errc := make(chan error)

	for i, file := range files {
		wg.Add(1)
		go func(i int, path string) {
			defer wg.Done()
			log("Reading " + path)

			name := parseName(path)
			entity, err := ReadEntityAndParse(name)
			if err != nil {
				errc <- err
			}

			ents[i] = entity
		}(i, file)
	}

	wg.Wait()
	close(errc)

	err = <-errc

	return ents, err
}

func CollectThumbnails(ents []*Entity, outputPath string) error {
	// create directory if it doesn't exist
	destination := outputPath + "/thumbnails"

	// remove previous files in directory
	files, err := filepath.Glob(destination + "/*")
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

	for index, entity := range ents {
		log("(" + fmt.Sprint(index+1) + "/" + fmt.Sprint(len(ents)) + ") Collecting thumbnail for " + entity.Name)

		response, err := http.Get(entity.Thumbnail)
		if err != nil {
			log("Error getting thumbnail: " + err.Error())
			return err
		}
		defer response.Body.Close()

		// get thumbnail file type from the end of the thumbnail url
		filetype := filepath.Ext(entity.Thumbnail)

		err = os.MkdirAll(destination, 0755)
		if err != nil {
			return err
		}

		unescaped, err := url.QueryUnescape(entity.Uri)
		if err != nil {
			return err
		}

		file, err := os.Create(destination + "/" + unescaped + filetype)
		if err != nil {
			return err
		}
		defer file.Close()

		_, err = io.Copy(file, response.Body)
		if err != nil {
			return err
		}

		if index < len(ents)-1 {
			// sleep if not the last entity to prevent spamming the server
			time.Sleep(time.Millisecond * 200)
		}

		// update the thumbnail to the new path
		ents[index].Thumbnail = "thumbnails/" + entity.Uri + filetype
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

// Tiles must be provided as a string
func entityTilesHash(str string) string {
	// generate a hash based on the entity tiles
	hash := md5.New()
	hash.Write([]byte(str))
	// truncate hash to 8 characters
	return fmt.Sprintf("%x", hash.Sum(nil))[:8]
}

func mapLink(d []byte) string {
	return "https://runelite.net/tile/show/" + strings.ReplaceAll(
		base64.StdEncoding.EncodeToString(d),
		"=",
		"",
	)
}

func transformEntity(entity *Entity) error {
	tilesString, err := json.Marshal(entity.Tiles)
	if err != nil {
		return err
	}

	entity.TilesString = string(tilesString)

	entity.MapLink = mapLink(tilesString)

	entity.FullName = fmt.Sprintf("%s %s", entity.Name, entity.Subcategory)
	entity.FullAltName = fmt.Sprintf("%s %s", entity.AltName, entity.Subcategory)

	hash := entityTilesHash(entity.TilesString)

	entity.Uri = getEntityUri(*entity)
	entity.SafeUri = url.QueryEscape(entity.Uri)

	entity.ApiUri = entity.Uri + "-" + hash
	entity.SafeApiUri = url.QueryEscape(entity.ApiUri)

	return nil
}

func parseEntity(data []byte) (*Entity, error) {
	var target *Entity
	err := json.Unmarshal(data, &target)
	if err != nil {
		return target, err
	}

	err = transformEntity(target)
	return target, err
}
