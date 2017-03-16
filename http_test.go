package goutil

import (
	"net/http"
	"testing"
)

func TestQueryValue(t *testing.T) {
	req, err := http.NewRequest("GET", "/?first=a&2nd=b&5nd=e&last=z", nil)
	if err != nil {
		t.Fatal(err)
	}

	var tests = []struct {
		param    string
		expected string
	}{
		{"first", "a"},
		{"2nd", "b"},
		{"5nd", "e"},
		{"last", "z"},
		{"notexists", ""},
	}

	for _, test := range tests {
		if got := QueryValue(test.param, req); got != test.expected {
			t.Errorf("QueryValue(%q) = %v, want %v", test.param, got, test.expected)
		}
	}

}
