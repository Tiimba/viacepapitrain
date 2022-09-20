package main

import (
	"testing"
)

type viacepTests struct {
	name     string
	url      string
	expected string
}

// test function
func TestGet1(t *testing.T) {

	tt := viacepTests{
		name:     "Get 13099160 info",
		url:      "https://viacep.com.br/ws/13099160/json/",
		expected: "Campinas",
	}

	if got := GetRequest(tt.url); got.Localidade != tt.expected {
		t.Errorf("GetRequest(%s) = %s; want %s", tt.url, got.Localidade, tt.expected)
	}
}
