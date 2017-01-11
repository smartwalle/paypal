package paypal

import (
	"fmt"
	"net/http"
)

const (
	k_CREATE_PAYMENT_API = "/v1/payments/payment"
)

// CreatePayment https://developer.paypal.com/docs/api/payments/#payment
// 因为接口返回的 payment 数据只比提交的 payment 数据多了几个字段，所以本接口的参数和返回结果共用同一数据结构。
func (this *PayPal) CreatePayment(payment *Payment) (result *Payment, err error) {
	var api = fmt.Sprintf("%s%s", this.APIBase, k_CREATE_PAYMENT_API)
	var req *http.Request
	req, err = this.request("POST", api, payment)
	if err != nil {
		return nil, err
	}
	err = this.doRequestWithAuth(req, &result)
	return result, err
}