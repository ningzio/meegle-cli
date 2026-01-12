package meegle

type AuthManager struct {
	PluginID     string
	PluginSecret string
	UserKey      string
}

func NewAuthManager(pluginID, pluginSecret, userKey string) *AuthManager {
	return &AuthManager{
		PluginID:     pluginID,
		PluginSecret: pluginSecret,
		UserKey:      userKey,
	}
}
