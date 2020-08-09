package config

import (
	"testing"
)

type TokenFile struct {
	Name string
	Data []byte
}

func (f *TokenFile) ReadFile(name string) ([]byte, error) {
	return []byte(`{"token": "test token data"}`), nil
}

func TestNewConfig(t *testing.T) {
	tokenFile := &TokenFile{}

	got, err := NewConfig("fileName", tokenFile)
	if err != nil {
		t.Errorf("[Error NewConfig]: %v", err)
	}

	testCase := "test token data"

	if got != testCase {
		t.Errorf("Got -> %v || Wanted -> %v", got, testCase)
	}
}
