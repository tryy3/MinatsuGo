package main

import (
	"encoding/json"
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
	f, err := os.Open(file)

	if !os.IsExist(err) {
		f, err = os.Create(file)

		if err != nil {
			return err
		}
	}

	encoder := json.NewEncoder(f)
	err = encoder.Encode(v)

	return err
}
