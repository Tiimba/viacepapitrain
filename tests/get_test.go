package cepai_testes

import (
	cepai "apicep/src"
	"os"
	"testing"
)

type viacepTests struct {
	name     string
	url      string
	expected int
}

func TestGet(t *testing.T) {
	t.Run("Should Return 200", func(t *testing.T) {
		tt := viacepTests{
			url:      os.Getenv("URL") + "13099160/json/",
			expected: 200,
		}

		if _, response_code := cepai.GetRequest(tt.url); response_code != tt.expected {
			t.Errorf("GetRequest(%s) = %d; want %d", tt.url, response_code, tt.expected)
		} else {
			t.Logf("GetRequest(%s) = %d; want %d", tt.url, response_code, tt.expected)
		}
	})

	t.Run("Should Return Empty", func(t *testing.T) {
		tt := viacepTests{
			url:      os.Getenv("URL") + "131099160/json/",
			expected: 400,
		}

		if _, response_code := cepai.GetRequest(tt.url); response_code != tt.expected {
			t.Errorf("GetRequest(%s) = %d; want %d", tt.url, response_code, tt.expected)
		} else {
			t.Logf("GetRequest(%s) = %d; want %d", tt.url, response_code, tt.expected)
		}

	})
}
