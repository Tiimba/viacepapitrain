package cepai_testes

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

func TestGet(t *testing.T) {
	t.Run("Should Return Campinas", func(t *testing.T) {
		tt := viacepTests{
			url:      os.Getenv("URL") + "13099160/json/",
			expected: "Campinas",
		}

		if got := cepai.GetRequest(tt.url); got.Localidade != tt.expected {
			t.Errorf("GetRequest(%s) = %s; want %s", tt.url, got.Localidade, tt.expected)
		} else {
			t.Logf("GetRequest(%s) = %s; want %s", tt.url, got.Localidade, tt.expected)
		}
	})

	t.Run("Should Return Empty", func(t *testing.T) {
		tt := viacepTests{
			url:      os.Getenv("URL") + "131099160/json/",
			expected: "",
		}

		if got := cepai.GetRequest(tt.url); got.Localidade != tt.expected {
			t.Errorf("GetRequest(%s) = %s; want %s", tt.url, got.Localidade, tt.expected)
		} else {
			t.Logf("GetRequest(%s) = %s; want %s", tt.url, got.Localidade, tt.expected)
		}

	})
}
