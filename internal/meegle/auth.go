package meegle

// AuthManager stores credentials needed to authenticate with Meegle.
type AuthManager struct {
	PluginID     string
	PluginSecret string
	UserKey      string
}

// NewAuthManager builds an AuthManager from plugin and user credentials.
func NewAuthManager(pluginID, pluginSecret, userKey string) *AuthManager {
	return &AuthManager{
		PluginID:     pluginID,
		PluginSecret: pluginSecret,
		UserKey:      userKey,
	}
}
