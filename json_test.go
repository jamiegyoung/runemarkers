package main

import "testing"

func TestParseJSON(t *testing.T) {
	type TestJSONStruct struct {
		Test string `json:"test"`
	}

	var target TestJSONStruct

	jsonData := []byte(`{"test": "this is a test"}`)
	err := ParseJSON(jsonData, &target)
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}

	if target.Test != "this is a test" {
		t.Errorf("Expected %v, got %v", "this is a test", target.Test)
	}
}

func TestParseInvalidJSON(t *testing.T) {
	type TestJSONStruct struct {
		Test string `json:"test"`
	}

	var target TestJSONStruct

	jsonData := []byte(`test: "test"`)
	err := ParseJSON(jsonData, &target)
	if err == nil {
		t.Errorf("Expected err")
	}

	if target.Test != "" {
		t.Errorf("Expected empty string, got %v", target.Test)
	}
}

func TestParseInvalidJSONTypeMismatch(t *testing.T) {
	type TestJSONStruct struct {
		Test string `json:"test"`
	}
	var target TestJSONStruct
	jsonData := []byte(`{"test": 1}`)
	err := ParseJSON(jsonData, &target)
	if err == nil {
		t.Errorf("Expected err")
	}
}
