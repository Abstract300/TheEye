package main

import (
	"encoding/json"
	"testing"
)

func TestGenerateConfig(t *testing.T) {
	var cfgInstance Config

	rawConfigData := []byte(`{"token": "kekekekek"}`)
	err := json.Unmarshal(rawConfigData, &cfgInstance)
	if err != nil {
		t.Fatalf("%s %s", "[error] Error unmarshalling marshalledConfigData for testing.", err)
	}

	gotFromFunction, _ := generateConfig(rawConfigData)
	if gotFromFunction.Token != cfgInstance.Token {
		t.Error("Token wasn't being properly parsed... got ", gotFromFunction.Token, " but needed => ", cfgInstance.Token)
	}

}
