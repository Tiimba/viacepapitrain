package cepai_testes

import (
	cepai "apicep/src"
	"os"
	"testing"
)

func TestGet(t *testing.T) {
	viacepTests := []struct {
		name     string
		url      string
		expected int
	}{
		{
			name:     "Should Return 200",
			url:      os.Getenv("URL") + "13099160/json/",
			expected: 200,
		},
		{
			name:     "Should return 400",
			url:      os.Getenv("URL") + "131099160/json/",
			expected: 400,
		},
	}

	for _, tt := range viacepTests {
		t.Run(tt.name, func(t *testing.T) {
			if _, response_code := cepai.GetRequest(tt.url); response_code != tt.expected {
				t.Errorf("GetRequest(%s) = %d; want %d", tt.url, response_code, tt.expected)
			} else {
				t.Logf("GetRequest(%s) = %d; want %d", tt.url, response_code, tt.expected)
			}
		})
	}
}
