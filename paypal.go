package paypal

import (
	"fmt"
	"time"
	"net/url"
	"net/http"
	"github.com/smartwalle/going/http2"
)

const (
	PAY_PAL_SANDBOX_API_URL = "https://api.sandbox.paypal.com"
	PAY_PAY_LVIE_API_URL    = "https://api.paypal.com"
)

const (
	k_GET_ACCESS_TOKEN_API = "/v1/oauth2/token"
)

type Client struct {
	ClientId string
	Secret   string
	APIBase  string
	Token    *Token
}

func NewClient(clientId, secret, API string) (client *Client) {
	client = &Client{}
	client.ClientId = clientId
	client.Secret = secret
	client.APIBase = API
	return client
}

func (this *Client) GetAccessToken() (token *Token, err error) {
	var api = fmt.Sprintf("%s%s", PAY_PAL_SANDBOX_API_URL, k_GET_ACCESS_TOKEN_API)
	fmt.Println(api)

	var p = url.Values{}
	p.Add("grant_type", "client_credentials")

	req, err := http2.NewRequest("POST", api, p)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-type", "application/x-www-form-urlencoded")
	req.Header.Add("Accept-Language", "en_US")
	req.SetBasicAuth(this.ClientId, this.Secret)

	rep, err := http2.DoJSONRequest(http.DefaultClient, req, &token)
	fmt.Println(rep.Status)
	if err == nil {
		token.ExpiresAt = time.Now().Add(time.Duration(token.ExpiresIn / 2) * time.Second)
		this.Token = token
	}
	return token, err
}
