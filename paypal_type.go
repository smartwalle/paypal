package paypal

import "time"

type jsonString string

func (this jsonString) MarshalJSON() ([]byte, error) {
	return []byte(string(this)), nil
}

type Token struct {
	Scope       string    `json:"scope"`
	Nonce       string    `json:"nonce"`
	AccessToken string    `json:"access_token"`
	TokenType   string    `json:"token_type"`
	AppId       string    `json:"app_id"`
	ExpiresIn   int       `json:"expires_in"`
	ExpiresAt   time.Time `json:"expires_at"`
}

type RelatedResources struct {
	Sale   *Sale   `json:"sale,omitempty"`
	Refund *Refund `json:"refund,omitempty"`
}

type Currency struct {
	Currency string `json:"currency,omitempty"`
	Value    string `json:"value,omitempty"`
}

type Link struct {
	Href    string `json:"href,omitempty"`
	Rel     string `json:"rel,omitempty"`
	Method  string `json:"method,omitempty"`
	EncType string `json:"encType,omitempty"`
}