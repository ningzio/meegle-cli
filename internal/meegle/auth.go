package meegle

import "time"

type authCache struct {
	token     string
	expiresAt time.Time
}

func (a *authCache) valid() bool {
	return a.token != "" && time.Now().Before(a.expiresAt)
}
