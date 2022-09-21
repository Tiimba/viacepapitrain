package main

import (
	"testing"
)

type viacepTests struct {
	name     string
	url      string
	expected string
}

func TestGetWrong(t *testing.T) {

	tt := viacepTests{
		name:     "Get 89099160 Wrong/Nil info",
		url:      "https://viacep.com.br/ws/89099160/json/",
		expected: "",
	}

	if got := GetRequest(tt.url); got.Localidade != tt.expected {
		t.Errorf("GetRequest(%s) = %s; want %s", tt.url, got.Localidade, tt.expected)
	}
}
