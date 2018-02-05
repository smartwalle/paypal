package paypal

import "time"

type Token struct {
	Scope       string    `json:"scope"`
	Nonce       string    `json:"nonce"`
	AccessToken string    `json:"access_token"`
	TokenType   string    `json:"token_type"`
	AppId       string    `json:"app_id"`
	ExpiresIn   int       `json:"expires_in"`
	ExpiresAt   time.Time `json:"expires_at"`
}
