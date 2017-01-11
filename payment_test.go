package paypal

import (
	"fmt"
	"testing"
)

var paypal *PayPal

func getPayPal() *PayPal {
	if paypal == nil {
		paypal = New("AT2V6Y2Kh7mFN5tE_c-BdeAyqS4HBcGF4Kl8seWPvA-jhY2CW6MMSr-t-mvf9F6GNFkobgtp6L2GDPxI", "EDjPl5bFyIrydUS__Nsd34l8t4O7aPEsd8Z_xY7cI_5_jTl1jt-Wn6QnzaB6-J3coWtrNoIc31DrblC3", PAY_PAL_SANDBOX_API_URL)
	}
	return paypal
}

func TestCreatePayment(t *testing.T) {
	var p = &Payment{}
	p.Intent = K_PAYPAL_INTENT_SALE
	p.Payer = &Payer{}
	p.Payer.PaymentMethod = "paypal"
	p.RedirectURLs = &RedirectURLs{}
	p.RedirectURLs.CancelURL = "http://www.baidu.com"
	p.RedirectURLs.ReturnURL = "http://www.qq.com"
	var ti = &Transaction{}
	p.Transactions = []*Transaction{ti}

	ti.Amount = &Amount{}
	ti.Amount.Total = "30.11"
	ti.Amount.Currency = "USD"
	ti.Amount.Details = &AmountDetails{}
	ti.Amount.Details.Subtotal = "30.00"
	ti.Amount.Details.Tax = "0.07"
	ti.Amount.Details.Shipping = "0.03"
	ti.Amount.Details.HandlingFee = "1.00"
	ti.Amount.Details.ShippingDiscount = "-1.00"
	ti.Amount.Details.Insurance = "0.01"

	ti.Description = "This is the payment transaction description."
	ti.Custom = "EBAY_EMS_90048630024435"
	ti.InvoiceNumber = "48787589673"

	ti.PaymentOptions = &PaymentOptions{}
	ti.PaymentOptions.AllowedPaymentMethod = "INSTANT_FUNDING_SOURCE"
	ti.SoftDescriptor = "ECHI5786786"

	ti.ItemList = &ItemList{}
	ti.ItemList.ShippingAddress = &ShippingAddress{}
	ti.ItemList.ShippingAddress.RecipientName = "Hello World"
	ti.ItemList.ShippingAddress.Line1 = "4thFloor"
	ti.ItemList.ShippingAddress.Line2 = "unit#34"
	ti.ItemList.ShippingAddress.City = "SAn Jose"
	ti.ItemList.ShippingAddress.CountryCode = "US"
	ti.ItemList.ShippingAddress.PostalCode = "95131"
	ti.ItemList.ShippingAddress.Phone = "011862212345678"
	ti.ItemList.ShippingAddress.State = "CA"

	var i1, i2 = &Item{}, &Item{}
	ti.ItemList.Items = []*Item{i1, i2}

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

	fmt.Println(getPayPal().CreatePayment(p))
}

//func TestGetPaymentList(t *testing.T) {
//	var p = &PaymentListParam{}
//	fmt.Println(getPayPal().GetPaymentList(p))
//}

//func TestGetPaymentDetails(t *testing.T) {
//	fmt.Println(getPayPal().GetPaymentDetails("PAY-9MR71248PB553763WLB24PAQ"))
//}

//func TestExecutePayment(t *testing.T) {
//	fmt.Println(getPayPal().ExecutePayment("PAY-6B21970242729815FLB243DA", "XV9HF9K25FB38"))
//}