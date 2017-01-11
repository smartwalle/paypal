package paypal

import (
	"testing"
	"fmt"
)

func TestPayment(te *testing.T) {
	var c = New("AT2V6Y2Kh7mFN5tE_c-BdeAyqS4HBcGF4Kl8seWPvA-jhY2CW6MMSr-t-mvf9F6GNFkobgtp6L2GDPxI", "EDjPl5bFyIrydUS__Nsd34l8t4O7aPEsd8Z_xY7cI_5_jTl1jt-Wn6QnzaB6-J3coWtrNoIc31DrblC3", PAY_PAL_SANDBOX_API_URL)

	var p = &Payment{}
	p.Intent = K_PAYPAL_INTENT_SALE
	p.Payer = &Payer{}
	p.Payer.PaymentMethod = "paypal"
	p.RedirectURLs = &RedirectURLs{}
	p.RedirectURLs.CancelURL = "http://www.baidu.com"
	p.RedirectURLs.ReturnURL = "http://www.qq.com"
	var t = &Transaction{}
	p.Transactions = []*Transaction{t}

	t.Amount = &Amount{}
	t.Amount.Total = "30.11"
	t.Amount.Currency = "USD"
	t.Amount.Details = &AmountDetails{}
	t.Amount.Details.Subtotal = "30.00"
	t.Amount.Details.Tax = "0.07"
	t.Amount.Details.Shipping = "0.03"
	t.Amount.Details.HandlingFee = "1.00"
	t.Amount.Details.ShippingDiscount = "-1.00"
	t.Amount.Details.Insurance = "0.01"

	t.Description = "This is the payment transaction description."
	t.Custom = "EBAY_EMS_90048630024435"
	t.InvoiceNumber = "48787589673"

	t.PaymentOptions = &PaymentOptions{}
	t.PaymentOptions.AllowedPaymentMethod = "INSTANT_FUNDING_SOURCE"
	t.SoftDescriptor = "ECHI5786786"

	t.ItemList = &ItemList{}
	t.ItemList.ShippingAddress = &ShippingAddress{}
	t.ItemList.ShippingAddress.RecipientName = "Hello World"
	t.ItemList.ShippingAddress.Line1 = "4thFloor"
	t.ItemList.ShippingAddress.Line2 = "unit#34"
	t.ItemList.ShippingAddress.City = "SAn Jose"
	t.ItemList.ShippingAddress.CountryCode = "US"
	t.ItemList.ShippingAddress.PostalCode = "95131"
	t.ItemList.ShippingAddress.Phone = "011862212345678"
	t.ItemList.ShippingAddress.State = "CA"

	var i1, i2 = &Item{}, &Item{}
	t.ItemList.Items = []*Item{i1, i2}

	i1.Name = "hat"
	i1.Description = "Brown color hat"
	i1.Quantity = 5
	i1.Price = "3"
	i1.Tax = "0.01"
	i1.SKU = "1"
	i1.Currency = "USD"

	i2.Name = "handbag"
	i2.Description = "Black color hand bag"
	i2.Quantity = 1
	i2.Price = "15"
	i2.Tax = "0.02"
	i2.SKU = "product34"
	i2.Currency = "USD"

	p.NoteToPayer = "Contact us for any questions on your order."

	fmt.Println(c.CreatePayment(p))
}
