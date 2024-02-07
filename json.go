package main

import (
	"encoding/json"
	"os"
)

func ParseJSON(jsonData []byte, target interface{}) error {
	return json.Unmarshal(jsonData, &target)
}


func ParseJSONFile(filename string, target interface{}) error {
	jsonData, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	return ParseJSON(jsonData, target)
}
