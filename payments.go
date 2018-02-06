package paypal

const (
	k_PAYMENT_API = "/v1/payments/payment"
	k_SALE_API    = "/v1/payments/sale"
	k_REFUND_API  = "/v1/payments/refund"
)

// CreatePayment https://developer.paypal.com/docs/api/payments/#payment
// 因为接口返回的 payment 数据只比提交的 payment 数据多了几个字段，所以本接口的参数和返回结果共用同一数据结构。
func (this *PayPal) CreatePayment(payment *Payment) (results *Payment, err error) {
	var api = this.BuildAPI(k_PAYMENT_API)
	err = this.doRequestWithAuth("POST", api, payment, &results)
	return results, err
}

func (this *PayPal) ExpressCreatePayment(total, currency, cancelURL, returnURL string) (results *Payment, err error) {
	var p = &Payment{}
	p.Intent = K_PAYMENT_INTENT_SALE
	p.Payer = &Payer{}
	p.Payer.PaymentMethod = "paypal"
	p.RedirectURLs = &RedirectURLs{}
	p.RedirectURLs.CancelURL = cancelURL
	p.RedirectURLs.ReturnURL = returnURL

	var transaction = &Transaction{}
	p.Transactions = []*Transaction{transaction}

	transaction.Amount = &Amount{}
	transaction.Amount.Total = total
	transaction.Amount.Currency = currency

	results, err = this.CreatePayment(p)
	return results, err
}

// GetPaymentList https://developer.paypal.com/docs/api/payments/#payment_list
func (this *PayPal) GetPaymentList(param *PaymentListParam) (results *PaymentList, err error) {
	var api = this.BuildAPI(k_PAYMENT_API) + param.QueryString()
	err = this.doRequestWithAuth("GET", api, nil, &results)
	return results, err
}

// GetPaymentDetails https://developer.paypal.com/docs/api/payments/#payment_get
func (this *PayPal) GetPaymentDetails(paymentId string) (results *Payment, err error) {
	var api = this.BuildAPI(k_PAYMENT_API, paymentId)
	err = this.doRequestWithAuth("GET", api, nil, &results)
	return results, err
}

// ExecuteApprovedPayment https://developer.paypal.com/docs/api/payments/#payment_execute
// 从回调 URL 中获取 PayerId
func (this *PayPal) ExecuteApprovedPayment(paymentId, payerId string) (results *Payment, err error) {
	var p = map[string]interface{}{}
	p["payer_id"] = payerId

	var api = this.BuildAPI(k_PAYMENT_API, paymentId, "execute")
	err = this.doRequestWithAuth("POST", api, p, &results)
	return results, err
}

// GetSaleDetails https://developer.paypal.com/docs/api/payments/#sale_get
func (this *PayPal) GetSaleDetails(saleId string) (results *Sale, err error) {
	var api = this.BuildAPI(k_SALE_API, saleId)
	err = this.doRequestWithAuth("GET", api, nil, &results)
	return results, err
}

// RefundSale https://developer.paypal.com/docs/api/payments/#sale_refund
func (this *PayPal) RefundSale(saleId, invoiceNumber, total, currency string) (results *Refund, err error) {
	var p = &refundSaleParam{}
	p.Amount.Total = total
	p.Amount.Currency = currency
	p.InvoiceNumber = invoiceNumber

	var api = this.BuildAPI(k_SALE_API, saleId, "/refund")
	err = this.doRequestWithAuth("POST", api, p, &results)
	return results, err
}

// GetRefundDetails https://developer.paypal.com/docs/api/payments/#refund_get
func (this *PayPal) GetRefundDetails(refundId string) (results *Refund, err error) {
	var api = this.BuildAPI(k_REFUND_API, refundId)
	err = this.doRequestWithAuth("GET", api, nil, &results)
	return results, err
}
