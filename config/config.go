package config

import (
	"encoding/json"

	"github.com/pkg/errors"
)

// TokenParser reads a file from disk for Config.
type TokenParser interface {
	ReadFile(name string) ([]byte, error)
}

// Config holds the Token information to start the bot.
type Config struct {
	disk  TokenParser
	Token string `json:"token"`
}

// NewConfig parses token data from a json file and returns it bo
func NewConfig(fileName string, file TokenParser) (string, error) {
	cfg := Config{
		disk: file,
	}

	fileData, err := file.ReadFile(fileName)
	if err != nil {
		return "", errors.Wrap(err, "Error reading "+fileName+" data.")
	}

	err = json.Unmarshal(fileData, &cfg)
	if err != nil {
		return "", errors.Wrap(err, "Errors unmarshalling.")
	}

	return cfg.Token, nil
}
