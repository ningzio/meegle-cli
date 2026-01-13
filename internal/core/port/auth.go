package port

import "context"

// TokenStore defines the interface for storing and retrieving authentication tokens.
type TokenStore interface {
	SaveToken(accessToken string, refreshToken string) error
	GetAccessToken() (string, error)
	// Add other methods as needed (e.g., GetRefreshToken, ClearToken)
}

// Authenticator defines the interface for the authentication flow.
type Authenticator interface {
	// Login initiates the OAuth2 flow and returns the access token.
	Login(ctx context.Context) (string, error)
}
