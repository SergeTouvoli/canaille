package compose

import (
	"os"

	"gopkg.in/yaml.v3"
)

func ParseFile(path string) (*ComposeFile, error) {
	data, err := os.ReadFile(path)

	if err != nil {
		return nil, err
	}

	var composeFile ComposeFile

	err = yaml.Unmarshal(data, &composeFile)

	if err != nil {
		return nil, err
	}

	return &composeFile, nil

}
