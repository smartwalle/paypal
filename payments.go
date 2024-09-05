package paypal

import "net/http"

const (
	kPayment = "/v1/payments/payment"
	kSale    = "/v1/payments/sale"
	kRefund  = "/v1/payments/refund"
)

// CreatePayment https://developer.paypal.com/docs/api/payments/#payment
// 因为接口返回的 payment 数据只比提交的 payment 数据多了几个字段，所以本接口的参数和返回结果共用同一数据结构。
func (c *Client) CreatePayment(payment *Payment) (result *Payment, err error) {
	var api = c.BuildAPI(kPayment)
	err = c.doRequestWithAuth(http.MethodPost, api, payment, &result)
	return result, err
}

func (c *Client) ExpressCreatePayment(invoiceNumber, total, currency, cancelURL, returnURL string) (result *Payment, err error) {
	var p = &Payment{}
	p.Intent = PaymentIntentSale
	p.Payer = &Payer{}
	p.Payer.PaymentMethod = PaymentMethodPayPal
	p.RedirectURLs = &RedirectURLs{}
	p.RedirectURLs.CancelURL = cancelURL
	p.RedirectURLs.ReturnURL = returnURL

	var transaction = &Transaction{}
	transaction.InvoiceNumber = invoiceNumber
	p.Transactions = []*Transaction{transaction}

	transaction.Amount = &Amount{}
	transaction.Amount.Total = total
	transaction.Amount.Currency = currency

	result, err = c.CreatePayment(p)
	return result, err
}

// GetPaymentList https://developer.paypal.com/docs/api/payments/#payment_list
func (c *Client) GetPaymentList(param *PaymentListParam) (result *PaymentList, err error) {
	var api = c.BuildAPI(kPayment) + param.QueryString()
	err = c.doRequestWithAuth(http.MethodGet, api, nil, &result)
	return result, err
}

// GetPaymentDetails https://developer.paypal.com/docs/api/payments/#payment_get
func (c *Client) GetPaymentDetails(paymentId string) (result *Payment, err error) {
	var api = c.BuildAPI(kPayment, paymentId)
	err = c.doRequestWithAuth(http.MethodGet, api, nil, &result)
	return result, err
}

// ExecuteApprovedPayment https://developer.paypal.com/docs/api/payments/#payment_execute
// 从回调 URL 中获取 PayerId
func (c *Client) ExecuteApprovedPayment(paymentId, payerId string) (result *Payment, err error) {
	var p = map[string]interface{}{}
	p["payer_id"] = payerId

	var api = c.BuildAPI(kPayment, paymentId, "execute")
	err = c.doRequestWithAuth(http.MethodPost, api, p, &result)
	return result, err
}

// GetSaleDetails https://developer.paypal.com/docs/api/payments/#sale_get
func (c *Client) GetSaleDetails(saleId string) (result *Sale, err error) {
	var api = c.BuildAPI(kSale, saleId)
	err = c.doRequestWithAuth(http.MethodGet, api, nil, &result)
	return result, err
}

// RefundSale https://developer.paypal.com/docs/api/payments/#sale_refund
func (c *Client) RefundSale(saleId string, param *RefundSaleParam) (result *Refund, err error) {
	var api = c.BuildAPI(kSale, saleId, "/refund")
	err = c.doRequestWithAuth(http.MethodPost, api, param, &result)
	return result, err
}

// GetRefundDetails https://developer.paypal.com/docs/api/payments/#refund_get
func (c *Client) GetRefundDetails(refundId string) (result *Refund, err error) {
	var api = c.BuildAPI(kRefund, refundId)
	err = c.doRequestWithAuth(http.MethodGet, api, nil, &result)
	return result, err
}
