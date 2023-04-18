package paypal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	kSandboxURL    = "https://api.sandbox.paypal.com"
	kProductionURL = "https://api.paypal.com"
)

const (
	kGetAccessTokenAPI = "/v1/oauth2/token"
)

type Client struct {
	clientId     string
	secret       string
	apiDomain    string
	isProduction bool
	Token        *Token
	Client       *http.Client
}

func New(clientId, secret string, isProduction bool) (client *Client) {
	client = &Client{}
	client.Client = http.DefaultClient
	client.clientId = clientId
	client.secret = secret
	client.isProduction = isProduction
	if isProduction {
		client.apiDomain = kProductionURL
	} else {
		client.apiDomain = kSandboxURL
	}
	return client
}

func (this *Client) BuildAPI(paths ...string) string {
	var path = this.apiDomain
	for _, p := range paths {
		p = strings.TrimSpace(p)
		if len(p) > 0 {
			if strings.HasSuffix(path, "/") {
				path = path + p
			} else {
				if strings.HasPrefix(p, "/") {
					path = path + p
				} else {
					path = path + "/" + p
				}
			}
		}
	}
	return path
}

func (this *Client) doRequestWithAuth(method, url string, param, result interface{}) (err error) {
	if this.Token == nil || this.Token.ExpiresAt.Before(time.Now()) {
		this.Token, err = this.GetAccessToken()
		if err != nil {
			return err
		}
	}

	var req *http.Request
	req, err = this.request(method, url, param)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bearer "+this.Token.AccessToken)
	return this.doRequest(req, result)
}

func (this *Client) GetAccessToken() (token *Token, err error) {
	var api = this.BuildAPI(kGetAccessTokenAPI)

	var param = url.Values{}
	param.Add("grant_type", "client_credentials")

	req, err := http.NewRequest("POST", api, strings.NewReader(param.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.SetBasicAuth(this.clientId, this.secret)

	err = this.doRequest(req, &token)
	if err != nil {
		return nil, err
	}
	if token != nil {
		token.ExpiresAt = time.Now().Add(time.Duration(token.ExpiresIn) * time.Second)
	}
	return token, err
}

func (this *Client) request(method, url string, param interface{}) (*http.Request, error) {
	var body io.Reader
	if param != nil {
		data, err := json.Marshal(param)
		if err != nil {
			return nil, err
		}
		body = bytes.NewBuffer(data)
	}
	return http.NewRequest(method, url, body)
}

func (this *Client) doRequest(req *http.Request, result interface{}) error {
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Accept-Language", "en_US")

	if req.Header.Get("Content-Type") == "" {
		req.Header.Add("Content-Type", "application/json")
	}

	var (
		err  error
		rsp  *http.Response
		data []byte
	)

	rsp, err = this.Client.Do(req)
	if err != nil {
		return err
	}
	defer rsp.Body.Close()

	data, err = io.ReadAll(rsp.Body)
	if err != nil {
		return err
	}

	if req.URL.Path != kGetAccessTokenAPI {
		var buf = &bytes.Buffer{}
		buf.WriteString("\n=========== Begin ============")
		buf.WriteString("\n【请求信息】")
		buf.WriteString(fmt.Sprintf("\n%s %d %s", req.Method, rsp.StatusCode, req.URL.String()))
		for key := range req.Header {
			buf.WriteString(fmt.Sprintf("\n%s: %s", key, req.Header.Get(key)))
		}
		buf.WriteString("\n【返回信息】")
		for key := range rsp.Header {
			buf.WriteString(fmt.Sprintf("\n%s: %s", key, rsp.Header.Get(key)))
		}
		buf.WriteString(fmt.Sprintf("\n%s", string(data)))
		buf.WriteString("\n===========  End  ============")

		logger.Println(buf.String())
	}

	switch rsp.StatusCode {
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
		e.Response = rsp
		if len(data) > 0 {
			if err = json.Unmarshal(data, e); err != nil {
				return err
			}
		}
		return e
	case http.StatusNoContent:
		if req.Method == http.MethodDelete {
			return nil
		}
		fallthrough
	default:
		var e = &ResponseError{}
		e.Response = rsp
		if len(data) > 0 {
			if err = json.Unmarshal(data, e); err != nil {
				return err
			}
		}
		return e
	}

	return err
}
