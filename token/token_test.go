package token

import (
	"testing"
)

type NoopTest struct{}

func (f NoopTest) ReadToken(name string) ([]byte, error) {
	return []byte(`{"token": "test token data"}`), nil
}

func TestNewConfig(t *testing.T) {
	tokenFile := NoopTest{}

	got, err := NewToken("fileName", tokenFile)
	if err != nil {
		t.Errorf("[Error NewConfig]: %v", err)
	}

	testCase := "test token data"

	if got != testCase {
		t.Errorf("Got -> %v || Wanted -> %v", got, testCase)
	}
}
