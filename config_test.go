package config

import (
	"fmt"
	"testing"
)

func TestReadKeyExists(t *testing.T) {
	config, err := Read("testconfig/simple.json")

	if err != nil {
		t.Fatal(err)
	}

	if config["key"] != "value" {
		t.Fatal("Unexpected config value")
	}
}

func TestReadKeyMissing(t *testing.T) {
	config, err := Read("testconfig/simple.json")

	if err != nil {
		t.Fatal(err)
	}

	if config["does not exist"] != nil {
		t.Fatal("Unexpected config value")
	}
}

func TestFileNotFound(t *testing.T) {
	_, err := Read("non existing filename")

	if err == nil {
		t.Fatal("err should not have been nil because file does not exist")
	}
}

func TestEmptyConfigFile(t *testing.T) {
	config, err := Read("testconfig/empty.json")

	if err != nil {
		t.Fatal("Returned nil, but shouldn't have")
	}

	if len(config) != 0 {
		t.Fatal("Expected empty map")
	}
}
