package main

import (
	"io/ioutil"

	"github.com/pkg/errors"
)

type Token struct {
	value string
}

func (tkn *Token) ReadFile(name string) ([]byte, error) {
	file, err := ioutil.ReadFile(name)
	if err != nil {
		return []byte{}, errors.Wrap(err, "Error reading config file: ")
	}
	return file, nil
}
