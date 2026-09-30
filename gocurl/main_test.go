package main

import "testing"

func TestParseArgs(t *testing.T) {
	config, err := parseArgs([]string{"example.com"})

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if config.URL != "example.com" {
		t.Errorf("expected URL example.com, got %s", config.URL)
	}

	if config.Method != "GET" {
		t.Errorf("expected method GET, got %s", config.Method)
	}

	if config.ShowHeaders {
		t.Errorf("expected flag ShowHeaders is eqaul to flase, got %t", config.ShowHeaders)
	}

	if !config.ShowBody {
		t.Errorf("expected ShowBodyu is equal to true, got %t", config.ShowBody)
	}
}
