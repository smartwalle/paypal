package paypal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	PAY_PAL_SANDBOX_API_URL = "https://api.sandbox.paypal.com"
	PAY_PAY_LVIE_API_URL    = "https://api.paypal.com"
)

const (
	k_GET_ACCESS_TOKEN_API = "/v1/oauth2/token"
)

type PayPal struct {
	clientId string
	secret   string
	APIBase  string
	Token    *Token
}

func New(clientId, secret, API string) (client *PayPal) {
	client = &PayPal{}
	client.clientId = clientId
	client.secret = secret
	client.APIBase = API
	return client
}

func (this *PayPal) request(method, url string, payload interface{}) (*http.Request, error) {
	var buf io.Reader
	if payload != nil {
		b, err := json.Marshal(payload)
		if err != nil {
			return nil, err
		}
		buf = bytes.NewBuffer(b)
		fmt.Println(string(b))
	}
	return http.NewRequest(method, url, buf)
}

func (this *PayPal) doRequest(req *http.Request, result interface{}) error {
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Accept-Language", "en_US")

	if req.Header.Get("Content-Type") == "" {
		req.Header.Add("Content-Type", "application/json")
	}

	var (
		err  error
		rep  *http.Response
		data []byte
	)

	rep, err = http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer rep.Body.Close()

	data, err = ioutil.ReadAll(rep.Body)
	if err != nil {
		return err
	}

	if req.URL.Path != k_GET_ACCESS_TOKEN_API {
		fmt.Println("=========== Begin ============")
		fmt.Println(req.Method, rep.StatusCode, req.URL.String())
		fmt.Println(req.Header)
		fmt.Println(string(data))
		fmt.Println("===========  End  ============")
	}

	switch rep.StatusCode {
	case http.StatusOK, http.StatusCreated:
		if result != nil {
			if err = json.Unmarshal(data, result); err != nil {
				if err.Error() == "json: cannot unmarshal number into Go value of type string" {
					return nil
				}
				return err
			}
		}
		return nil
	case http.StatusUnauthorized:
		var e = &IdentityError{}
		e.Response = rep
		if len(data) > 0 {
			if err = json.Unmarshal(data, e); err != nil {
				return err
			}
		}
		return e
	default:
		var e = &ResponseError{}
		e.Response = rep
		if len(data) > 0 {
			if err = json.Unmarshal(data, e); err != nil {
				return err
			}
		}
		return e
	}

	return err
}

func (this *PayPal) doRequestWithAuth(req *http.Request, result interface{}) (err error) {
	if this.Token == nil || this.Token.ExpiresAt.Before(time.Now()) {
		this.Token, err = this.GetAccessToken()
		if err != nil {
			return err
		}
	}
	req.Header.Set("Authorization", "Bearer " + this.Token.AccessToken)
	return this.doRequest(req, result)
}

func (this *PayPal) GetAccessToken() (token *Token, err error) {
	var api = fmt.Sprintf("%s%s", this.APIBase, k_GET_ACCESS_TOKEN_API)

	var p = url.Values{}
	p.Add("grant_type", "client_credentials")

	req, err := http.NewRequest("POST", api, strings.NewReader(p.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.SetBasicAuth(this.clientId, this.secret)

	err = this.doRequest(req, &token)
	if err != nil {
		return nil, err
	}
	if token != nil {
		token.ExpiresAt = time.Now().Add(time.Duration(token.ExpiresIn/2) * time.Second)
		this.Token = token
	}
	return token, err
}
