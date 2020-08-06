package main

import (
	"encoding/json"
	"io/ioutil"

	"github.com/pkg/errors"
)

const (
	CONFIG_FILE = "config.json"
)

type Config struct {
	Token string `json:"token"`
}

func NewConfig(filename string) (Config, error) {
	fileData, err := ioutil.ReadFile(filename)
	if err != nil {
		return Config{}, errors.Wrap(err, "Error reading"+CONFIG_FILE+" data.")
	}

	return generateConfig(fileData)
}

func generateConfig(fileData []byte) (Config, error) {
	var cfg Config

	err := json.Unmarshal(fileData, &cfg)
	if err != nil {
		return Config{}, errors.Wrap(err, "Errors unmarshalling config.json.")
	}
	return cfg, nil
}
