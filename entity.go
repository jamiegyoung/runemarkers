package main

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"strings"
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

func urlEncode(s string) string {
	lowered := strings.ToLower(s)
	unspacedAndLowered := strings.ReplaceAll(lowered, " ", "-")
	return url.QueryEscape(unspacedAndLowered)
}

func ReadEntityAndParse(name string) (*Entity, error) {
	data, err := os.ReadFile("entities/" + name + ".json")
	if err != nil {
		return nil, err
	}

	return parseEntity(data)
}

func transformEntity(entity *Entity) {
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
