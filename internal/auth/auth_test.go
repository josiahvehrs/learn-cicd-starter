package auth_test

import (
	"net/http"
	"testing"

	"github.com/bootdotdev/learn-cicd-starter/internal/auth"
)

func TestGetApiKey(t *testing.T) {
	header := http.Header{}
	header.Add("Authorization", "ApiKey test123")

	key, err := auth.GetAPIKey(header)
	if err != nil {
		t.Error("expected err to be nil")
	}
	if key != "test123" {
		t.Errorf("apikey does not match expected %s. got%s", "test123", key)
	}
}
