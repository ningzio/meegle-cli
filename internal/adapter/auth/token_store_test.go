package auth_test

import (
	"os"
	"path/filepath"
	"testing"

	"meegle-cli/internal/adapter/auth"
)

func TestFileTokenStore(t *testing.T) {
	// Setup temp home dir
	tempHome := t.TempDir()
	t.Setenv("HOME", tempHome)

	store, err := auth.NewFileTokenStore()
	if err != nil {
		t.Fatalf("Failed to create token store: %v", err)
	}

	// Test SaveToken
	err = store.SaveToken("test_access_token", "test_refresh_token")
	if err != nil {
		t.Errorf("SaveToken failed: %v", err)
	}

	// Verify file existence
	authPath := filepath.Join(tempHome, ".meegle", "auth.json")
	if _, err := os.Stat(authPath); os.IsNotExist(err) {
		t.Errorf("Auth file was not created at %s", authPath)
	}

	// Test GetAccessToken
	token, err := store.GetAccessToken()
	if err != nil {
		t.Errorf("GetAccessToken failed: %v", err)
	}
	if token != "test_access_token" {
		t.Errorf("Expected 'test_access_token', got '%s'", token)
	}

	// Test GetAccessToken with missing file
	os.Remove(authPath)
	_, err = store.GetAccessToken()
	if err == nil {
		t.Error("Expected error when auth file is missing, got nil")
	}
}
