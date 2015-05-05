package config

import (
	"encoding/json"
	"io/ioutil"
)

func Read(filename string) (config map[string]interface{}, err error) {
	fileContents, err := ioutil.ReadFile(filename)

	if err == nil {
		json.Unmarshal(fileContents, &config)
	}

	return
}
