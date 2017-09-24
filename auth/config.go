package auth

import "net/http"

type Config struct {
	CookieSecret []byte

	BaseURL      string
	Issuer       string
	ClientID     string
	ClientSecret string

	Get func(info UserInfo) (*UserIdentifier, error)

	SuccessHandler    http.Handler
	FailureHandler    http.Handler
	PostLogoutHandler http.Handler
}
