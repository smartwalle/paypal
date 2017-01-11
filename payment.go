package paypal

import (
	"fmt"
	"net/http"
	"net/url"
)

const (
	k_PAYMENT_API = "/v1/payments/payment"
	k_SALE_API    = "/v1/payments/sale"
)

// CreatePayment https://developer.paypal.com/docs/api/payments/#payment
// 因为接口返回的 payment 数据只比提交的 payment 数据多了几个字段，所以本接口的参数和返回结果共用同一数据结构。
func (this *PayPal) CreatePayment(payment *Payment) (results *Payment, err error) {
	var api = this.API(k_PAYMENT_API)
	var req *http.Request
	req, err = this.request("POST", api, payment)
	if err != nil {
		return nil, err
	}
	err = this.doRequestWithAuth(req, &results)
	return results, err
}

type PaymentListParam struct {
	Count      int
	StartId    string
	StartIndex int
	StartTime  string
	EndTime    string
	SortBy     string
	SortOrder  string
}

func (this *PaymentListParam) QueryString() string {
	var p = url.Values{}
	if len(this.StartId) > 0 {
		p.Set("start_id", this.StartId)
	}
	if len(this.StartTime) > 0 {
		p.Set("start_time", this.StartTime)
	}
	if len(this.EndTime) > 0 {
		p.Set("end_time", this.EndTime)
	}
	if this.StartIndex > 0 {
		p.Set("start_index", fmt.Sprintf("%d", this.StartIndex))
	}
	if this.Count > 0 {
		p.Set("count", fmt.Sprintf("%f", this.Count))
	}
	if len(this.SortBy) > 0 {
		p.Set("sort_by", this.SortBy)
	}
	if len(this.SortOrder) > 0 {
		p.Set("sort_order", this.SortOrder)
	}
	return "?" + p.Encode()
}

type PaymentListResp struct {
	Payments []*Payment `json:"payments"`
	Count    int        `json:"count"`
	NextId   string     `json:"next_id"`
}

// GetPaymentList https://developer.paypal.com/docs/api/payments/#payment_list
func (this *PayPal) GetPaymentList(param *PaymentListParam) (results *PaymentListResp, err error) {
	var api = this.API(k_PAYMENT_API) + param.QueryString()
	var req *http.Request
	req, err = this.request("GET", api, nil)
	if err != nil {
		return nil, err
	}
	err = this.doRequestWithAuth(req, &results)
	return results, err
}

// GetPaymentDetails https://developer.paypal.com/docs/api/payments/#payment_get
func (this *PayPal) GetPaymentDetails(paymentId string) (results *Payment, err error) {
	var api = this.API(k_PAYMENT_API, paymentId)
	var req *http.Request
	req, err = this.request("GET", api, nil)
	if err != nil {
		return nil, err
	}
	err = this.doRequestWithAuth(req, &results)
	return results, err
}

// ExecuteApprovedPayment https://developer.paypal.com/docs/api/payments/#payment_execute
// 从回调 URL 中获取 PayerId
func (this *PayPal) ExecuteApprovedPayment(paymentId, payerId string) (results *Payment, err error) {
	var p = map[string]interface{}{}
	p["payer_id"] = payerId

	var api = this.API(k_PAYMENT_API, paymentId, "execute")
	var req *http.Request
	req, err = this.request("POST", api, p)
	if err != nil {
		return nil, err
	}
	err = this.doRequestWithAuth(req, &results)
	return results, err
}

// GetSaleDetails https://developer.paypal.com/docs/api/payments/#sale_get
func (this *PayPal) GetSaleDetails(saleId string) (results *Sale, err error) {
	var api = this.API(k_SALE_API, saleId)
	var req *http.Request
	req, err = this.request("GET", api, nil)
	if err != nil {
		return nil, err
	}
	err = this.doRequestWithAuth(req, &results)

	return results, err
}
