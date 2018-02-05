package paypal

import (
	"fmt"
	"net/http"
)

const (
	K_ERR_TYPE_VALIDATION_ERROR     = "VALIDATION_ERROR"
	K_ERR_TYPE_INSTRUMENT_DECLINED  = "INSTRUMENT_DECLINED"
	K_ERR_TYPE_PAYMENT_ALREADY_DONE = "PAYMENT_ALREADY_DONE"
)

type ResponseError struct {
	Response        *http.Response `json:"-"`
	Name            string         `json:"name"`
	Message         string         `json:"message"`
	InformationLink string         `json:"information_link"`
	DebugId         string         `json:"debug_id"`
	Details         []ErrorDetail  `json:"details"`
}

func (this *ResponseError) Error() string {
	return fmt.Sprintf("[%s]%d %s [Name]%s [Message]%s", this.Response.Request.Method, this.Response.StatusCode, this.Response.Request.URL, this.Name, this.Message)
}

type ErrorDetail struct {
	Field string `json:"field"`
	Issue string `json:"issue"`
}

type IdentityError struct {
	Response         *http.Response `json:"-"`
	Name             string         `json:"error"`
	ErrorDescription string         `json:"error_description"`
}

func (this *IdentityError) Error() string {
	return fmt.Sprintf("[%s]%d %s [Error]%s [Description]%s", this.Response.Request.Method, this.Response.StatusCode, this.Response.Request.URL, this.Name, this.ErrorDescription)
}
