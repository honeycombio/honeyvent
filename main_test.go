package main

import (
	"testing"
)

func TestAPIHostParsing(t *testing.T) {
	u, err := parseAPIHost("api.honeycomb.io")
	if err != nil {
		t.Error(err)
	}
	if u.Scheme != "https" {
		t.Errorf("Expected https, got: %s", u.Scheme)
	}
	u, err = parseAPIHost("https://api.honeycomb.io")
	if err != nil {
		t.Error(err)
	}
	if u.Scheme != "https" {
		t.Errorf("Expected https, got: %s", u.Scheme)
	}
}
