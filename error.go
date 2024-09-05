package paypal

import (
	"fmt"
	"net/http"
)

const (
	ErrTypeValidation         = "VALIDATION_ERROR"
	ErrTypeInstrumentDeclined = "INSTRUMENT_DECLINED"
	ErrTypePaymentAlreadyDone = "PAYMENT_ALREADY_DONE"
)

type ResponseError struct {
	Response        *http.Response `json:"-"`
	Name            string         `json:"name"`
	Message         string         `json:"message"`
	InformationLink string         `json:"information_link"`
	DebugId         string         `json:"debug_id"`
	Details         []ErrorDetail  `json:"details"`
}

func (err *ResponseError) Error() string {
	return fmt.Sprintf("[%s]%d %s [Name]%s [Message]%s", err.Response.Request.Method, err.Response.StatusCode, err.Response.Request.URL, err.Name, err.Message)
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

func (err *IdentityError) Error() string {
	return fmt.Sprintf("[%s]%d %s [Error]%s [Description]%s", err.Response.Request.Method, err.Response.StatusCode, err.Response.Request.URL, err.Name, err.ErrorDescription)
}
