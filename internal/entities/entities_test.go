package entities

import (
	"fmt"
	"reflect"
	"testing"
)

func TestParseEntity(t *testing.T) {
	mockData := []byte("{\"name\":\"a test creature & thing\",\"altName\":\"A massive creature\",\"subcategory\":\"what!\",\"thumbnail\":\"thumbnailurl\",\"wiki\":\"wikiurl\",\"tags\":[\"cox\",\"chambers of xeric\",\"raids\"],\"recommendedGuideVideoId\":\"dQw4w9WgXcQ\",\"source\":{\"name\":\"sourcename\",\"link\":\"sourcelink\",\"modified\":\"amodification\"},\"tiles\":[{\"regionId\":12889,\"regionX\":37,\"regionY\":43,\"z\":0,\"color\":{\"value\":-1,\"falpha\":0}},{\"regionId\":12889,\"regionX\":37,\"regionY\":45,\"z\":0,\"color\":{\"value\":-1,\"falpha\":0}}]}")

	target, err := parseEntity(mockData)
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}

	correctData := Entity{
		Uri:                     "a-test-creature-&-thing-(what!)",
		Name:                    "a test creature & thing",
		AltName:                 "A massive creature",
		Subcategory:             "what!",
		Thumbnail:               "thumbnailurl",
		RecommendedGuideVideoId: "dQw4w9WgXcQ",
		Wiki:                    "wikiurl",
		Tags:                    []string{"cox", "chambers of xeric", "raids"},
		Source: &Source{
			Name:     "sourcename",
			Link:     "sourcelink",
			Modified: "amodification",
		},
		Tiles: []Tile{
			{
				RegionId: 12889,
				RegionX:  37,
				RegionY:  43,
				Z:        0,
				Color: interface{}(map[string]interface{}{
					"value":  -1,
					"falpha": 0,
				}),
			},
			{
				RegionId: 12889,
				RegionX:  37,
				RegionY:  45,
				Z:        0,
				Color: interface{}(map[string]interface{}{
					"value":  -1,
					"falpha": 0,
				}),
			},
		},
	}

	// compare the data
	if !reflect.DeepEqual(*target, correctData) {
		if target.Uri != correctData.Uri {
			t.Errorf("Expected %v, got %v", correctData.Uri, target.Uri)
		}
		if target.Name != correctData.Name {
			t.Errorf("Expected %v, got %v", correctData.Name, target.Name)
		}
		if target.Thumbnail != correctData.Thumbnail {
			t.Errorf("Expected %v, got %v", correctData.Thumbnail, target.Thumbnail)
		}
		if target.Wiki != correctData.Wiki {
			t.Errorf("Expected %v, got %v", correctData.Wiki, target.Wiki)
		}
		if !reflect.DeepEqual(target.Tags, correctData.Tags) {
			t.Errorf("Expected %v, got %v", correctData.Tags, target.Tags)
		}
		if !reflect.DeepEqual(target.Source, correctData.Source) {
			t.Errorf("Expected %v, got %v", correctData.Source, target.Source)
		}
		// compare strings of tiles
		if !reflect.DeepEqual(target.Tiles, correctData.Tiles) {
			for i, tile := range target.Tiles {
				tileAsString := fmt.Sprintf("%v", tile)
				correctTileAsString := fmt.Sprintf("%v", correctData.Tiles[i])
				if tileAsString != correctTileAsString {
					t.Errorf("Expected %v, got %v", correctTileAsString, tileAsString)
				}
			}
		}
	}
}

func TestUrlEncode(t *testing.T) {
	safeString := transformToUrl("Test (Entity) another")

	if safeString != "test-(entity)-another" {
		t.Errorf("Expected %v, got %v", "test-(entity)-another", safeString)
	}
}

func TestGetEntityFail(t *testing.T) {
	_, err := parseEntity([]byte("{\"name\":2}"))
	if err == nil {
		t.Errorf("Expected err")
	}
}
