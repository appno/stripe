package cmd

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
)

func unmarshalInput(bytes []byte) (interface{}, error) {
	var result interface{}
	err := json.Unmarshal(bytes, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func readInput(args []string, index int, path string) ([]byte, error) {
	if len(args) > index {
		return []byte(args[index]), nil
	}
	return readFile(path)
}

func readFile(path string) ([]byte, error) {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}
	bytes, err := ioutil.ReadFile(absPath)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}
