package auth

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// FileTokenStore implements port.TokenStore using a local JSON file.
type FileTokenStore struct {
	filePath string
}

type tokenData struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// NewFileTokenStore creates a new FileTokenStore.
func NewFileTokenStore() (*FileTokenStore, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	// Define the path: ~/.meegle/auth.json
	path := filepath.Join(home, ".meegle", "auth.json")
	return &FileTokenStore{filePath: path}, nil
}

// SaveToken saves the tokens to the file.
func (s *FileTokenStore) SaveToken(accessToken, refreshToken string) error {
	data := tokenData{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	bytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal token data: %w", err)
	}

	// Ensure directory exists
	dir := filepath.Dir(s.filePath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create config directory: %w", err)
	}

	if err := os.WriteFile(s.filePath, bytes, 0600); err != nil {
		return fmt.Errorf("failed to write auth file: %w", err)
	}

	return nil
}

// GetAccessToken retrieves the access token from the file.
func (s *FileTokenStore) GetAccessToken() (string, error) {
	bytes, err := os.ReadFile(s.filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return "", fmt.Errorf("no token found, please login first")
		}
		return "", fmt.Errorf("failed to read auth file: %w", err)
	}

	var data tokenData
	if err := json.Unmarshal(bytes, &data); err != nil {
		return "", fmt.Errorf("failed to parse auth file: %w", err)
	}

	if data.AccessToken == "" {
		return "", fmt.Errorf("access token is empty, please login again")
	}

	return data.AccessToken, nil
}
