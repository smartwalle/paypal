package paypal

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
	"log"
	"fmt"
)

const (
	k_PAY_PAL_SANDBOX_API_URL    = "https://api.sandbox.paypal.com"
	k_PAY_PAL_PRODUCTION_API_URL = "https://api.paypal.com"
)

const (
	k_GET_ACCESS_TOKEN_API = "/v1/oauth2/token"
)

type PayPal struct {
	clientId     string
	secret       string
	apiDomain    string
	isProduction bool
	Token        *Token
	logger       *log.Logger
}

func New(clientId, secret string, isProduction bool) (client *PayPal) {
	client = &PayPal{}
	client.clientId = clientId
	client.secret = secret
	client.isProduction = isProduction
	if isProduction {
		client.apiDomain = k_PAY_PAL_PRODUCTION_API_URL
	} else {
		client.apiDomain = k_PAY_PAL_SANDBOX_API_URL
	}
	return client
}

func (this *PayPal) BuildAPI(paths ...string) string {
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

func (this *PayPal) SetLogWriter(w io.Writer) {
	if w == nil {
		this.logger = nil
		return
	}
	if this.logger != nil {
		this.logger.SetOutput(w)
		return
	}
	this.logger = log.New(w, "[PayPal]", log.Ldate|log.Ltime)
}

func (this *PayPal) log(args ...interface{}) {
	if this.logger != nil {
		this.logger.Println(args...)
	}
}

func (this *PayPal) doRequestWithAuth(method, url string, param, result interface{}) (err error) {
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

func (this *PayPal) GetAccessToken() (token *Token, err error) {
	var api = this.BuildAPI(k_GET_ACCESS_TOKEN_API)

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
		token.ExpiresAt = time.Now().Add(time.Duration(token.ExpiresIn) * time.Second)
	}
	return token, err
}

func (this *PayPal) request(method, url string, payload interface{}) (*http.Request, error) {
	var buf io.Reader
	if payload != nil {
		b, err := json.Marshal(payload)
		if err != nil {
			return nil, err
		}
		buf = bytes.NewBuffer(b)
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
		rsp  *http.Response
		data []byte
	)

	rsp, err = http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer rsp.Body.Close()

	data, err = ioutil.ReadAll(rsp.Body)
	if err != nil {
		return err
	}

	if req.URL.Path != k_GET_ACCESS_TOKEN_API {
		if this.logger != nil {
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

			this.log(buf.String())
		}
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
		if req.Method == "DELETE" {
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
