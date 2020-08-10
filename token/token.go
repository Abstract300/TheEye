package token

import (
	"encoding/json"
	"os"

	"github.com/pkg/errors"
)

type TokenParser interface {
	ParseToken(fileName string) ([]byte, error)
}

// Noop is a no allocation struct used just to satify TokenParser
type Noop struct{}

// parseToken reads the token file and returns the file data; fails with an error otherwise.
func (nop Noop) ParseToken(fileName string) ([]byte, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return []byte{}, errors.Wrap(err, "Could not open file to parse token.")
	}
	fileInfo, err := os.Stat(fileName)
	if err != nil {
		return []byte{}, errors.Wrap(err, "Could not stat the file.")
	}

	fileData := make([]byte, fileInfo.Size())

	_, err = file.Read(fileData)
	return fileData, errors.Wrap(err, "Could not read content for the file.")

}

// NewToken takes a token.json file and returns a token; fails with an error otherwise.
func NewToken(fileName string, parser TokenParser) (string, error) {
	type Auth struct {
		Token string `json:"token"`
	}

	cfg := Auth{}

	fileData, err := parser.ParseToken(fileName)
	if err != nil {
		return "", errors.Wrap(err, "Error reading "+fileName+" data.")
	}

	err = json.Unmarshal(fileData, &cfg)
	if err != nil {
		return "", errors.Wrap(err, "Errors unmarshalling.")
	}

	return cfg.Token, nil
}
