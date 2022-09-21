package main

import (
	"cepai"
	"os"
	"testing"
)

type viacepTests struct {
	name     string
	url      string
	expected string
}

func TestGetCorrect(t *testing.T) {

	tt := viacepTests{
		name:     "Get Correct info value: '13099160'",
		url:      os.Getenv("URL") + "13099160/json/",
		expected: "Campinas",
	}

	if got := cepai.GetRequest(tt.url); got.Localidade != tt.expected {
		t.Errorf("GetRequest(%s) = %s; want %s", tt.url, got.Localidade, tt.expected)
	} else {
		t.Logf("GetRequest(%s) = %s; want %s", tt.url, got.Localidade, tt.expected)
	}
}

func TestGetWrong(t *testing.T) {

	tt := viacepTests{
		name:     "Get 89099160 Wrong/Nil info",
		url:      "https://viacep.com.br/ws/89099160/json/",
		expected: "",
	}

	if got := cepai.GetRequest(tt.url); got.Localidade != tt.expected {
		t.Errorf("GetRequest(%s) = %s; want %s", tt.url, got.Localidade, tt.expected)
	}
}
