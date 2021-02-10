package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type SystemReader struct{}

func (s SystemReader) Read(path string) (*Config, error) {
	data, err := ioutil.ReadFile(path)

	if err != nil {
		return nil, err
	}

	var cnf Config
	err = yaml.Unmarshal(data, &cnf)

	if err != nil {
		return nil, err
	}

	return &cnf, nil
}

func NewReader() Reader {
	return &SystemReader{}
}
