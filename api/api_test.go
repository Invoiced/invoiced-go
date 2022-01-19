package api

import "testing"

func TestApi(t *testing.T) {
	api := New("test", false)
	if api.Api.Sandbox {
		t.Fatal("Sandbox value did not match")
	}
	if api.Api.Key != "test" {
		t.Fatal("API key value did not match")
	}
}
