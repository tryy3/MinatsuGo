package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func loadConfiguration(file string, v interface{}) (err error) {
	f, err := os.Open(file)

	if err != nil && !os.IsExist(err) {
		return saveConfiguration(file, v)
	}

	decoder := json.NewDecoder(f)
	return decoder.Decode(v)
}

func saveConfiguration(file string, v interface{}) (err error) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(file, b, os.FileMode(0644))
}
