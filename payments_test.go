package paypal_test

import (
	"fmt"
	"github.com/smartwalle/paypal"
	"testing"
	"time"
)

func TestPayPal_ExpressCreatePayment(t *testing.T) {
	var p, err = client.ExpressCreatePayment("test_invoice_number", "10", "USD", "http://www.baidu.com", "http://192.168.192.250:3000/paypal")
	if err != nil {
		t.Fatal(err)
	}
	for _, link := range p.Links {
		t.Log(link.Method, link.Rel, link.Href)
	}
}

func TestPayPal_CreatePayment(t *testing.T) {
	var p = &paypal.Payment{}
	p.Intent = paypal.PaymentIntentSale
	p.Payer = &paypal.Payer{}
	p.Payer.PaymentMethod = "paypal"
	p.RedirectURLs = &paypal.RedirectURLs{}
	p.RedirectURLs.CancelURL = "http://www.baidu.com"
	p.RedirectURLs.ReturnURL = "http://127.0.0.1:9001/paypal"

	var transaction = &paypal.Transaction{}
	p.Transactions = []*paypal.Transaction{transaction}

	transaction.Amount = &paypal.Amount{}
	transaction.Amount.Total = "30.11"
	transaction.Amount.Currency = "USD"
	transaction.Amount.Details = &paypal.AmountDetails{}
	transaction.Amount.Details.Subtotal = "30.00"
	transaction.Amount.Details.Tax = "0.07"
	transaction.Amount.Details.Shipping = "0.03"
	transaction.Amount.Details.HandlingFee = "1.00"
	transaction.Amount.Details.ShippingDiscount = "-1.00"
	transaction.Amount.Details.Insurance = "0.01"

	transaction.Description = "This is the payment transaction description."
	transaction.Custom = "EBAY_EMS_90048630024435"
	transaction.InvoiceNumber = fmt.Sprintf("%d", time.Now().UnixMicro())

	transaction.PaymentOptions = &paypal.PaymentOptions{}
	transaction.PaymentOptions.AllowedPaymentMethod = "INSTANT_FUNDING_SOURCE"
	transaction.SoftDescriptor = "ECHI5786786"

	transaction.ItemList = &paypal.ItemList{}
	transaction.ItemList.ShippingAddress = &paypal.ShippingAddress{}
	transaction.ItemList.ShippingAddress.RecipientName = "Hello World"
	transaction.ItemList.ShippingAddress.Line1 = "4thFloor"
	transaction.ItemList.ShippingAddress.Line2 = "unit#34"
	transaction.ItemList.ShippingAddress.City = "SAn Jose"
	transaction.ItemList.ShippingAddress.CountryCode = "US"
	transaction.ItemList.ShippingAddress.PostalCode = "95131"
	transaction.ItemList.ShippingAddress.Phone = "011862212345678"
	transaction.ItemList.ShippingAddress.State = "CA"

	var i1, i2 = &paypal.Item{}, &paypal.Item{}
	transaction.ItemList.Items = []*paypal.Item{i1, i2}

	i1.Name = "hat"
	i1.Description = "Brown color hat"
	i1.Quantity = "5"
	i1.Price = "3"
	i1.Tax = "0.01"
	i1.SKU = "1"
	i1.Currency = "USD"

	i2.Name = "handbag"
	i2.Description = "Black color hand bag"
	i2.Quantity = "1"
	i2.Price = "15"
	i2.Tax = "0.02"
	i2.SKU = "product34"
	i2.Currency = "USD"

	p.NoteToPayer = "Contact us for any questions on your order."

	var payment, err = client.CreatePayment(p)
	if err != nil {
		t.Fatal(err)
	}

	if payment != nil {
		t.Log("CreatePayment", payment.Id)
	}

	for _, link := range payment.Links {
		t.Log(link.Method, link.Rel, link.Href)
	}
}

func TestPayPal_GetPaymentList(t *testing.T) {
	var p = &paypal.PaymentListParam{}
	t.Log(client.GetPaymentList(p))
}

func TestPayPal_GetPaymentDetails(t *testing.T) {
	t.Log(client.GetPaymentDetails("PAY-1SJ16214TY566804MLB26S5I"))
}

func TestPayPal_ExecuteApprovedPayment(t *testing.T) {
	t.Log(client.ExecuteApprovedPayment("PAY-0E809244MY2080201LB26LWI", "XV9HF9K25FB38"))
}

func TestPayPal_GetSaleDetails(t *testing.T) {
	t.Log(client.GetSaleDetails("84E33686SW065691F"))
}

func TestPayPal_RefundSale(t *testing.T) {
	var p = &paypal.RefundSaleParam{}
	p.InvoiceNumber = "48787589677"
	p.Amount = &paypal.Amount{}
	p.Amount.Total = "30.11"
	p.Amount.Currency = "USD"
	p.Amount.Total = "30.11"
	t.Log(client.RefundSale("5SW33389HH3038001", p))
}

func TestPayPal_GetRefundDetails(t *testing.T) {
	t.Log(client.GetRefundDetails("7AG891841V2503230"))
}
