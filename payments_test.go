package paypal

import (
	"testing"
)

//func TestPayPal_ExpressCreatePayment(t *testing.T) {
//	fmt.Println(getPayPal().ExpressCreatePayment("10", "USD", "http://www.baidu.com", "http://192.168.192.250:3000/paypal"))
//}
//
func TestPayPal_CreatePayment(t *testing.T) {
	//var p = &Payment{}
	//p.Intent = K_PAYMENT_INTENT_SALE
	//p.Payer = &Payer{}
	//p.Payer.PaymentMethod = "paypal"
	//p.RedirectURLs = &RedirectURLs{}
	//p.RedirectURLs.CancelURL = "http://www.baidu.com"
	//p.RedirectURLs.ReturnURL = "http://127.0.0.1:9001/paypal"
	//
	//var transaction = &Transaction{}
	//p.Transactions = []*Transaction{transaction}
	//
	//transaction.Amount = &Amount{}
	//transaction.Amount.Total = "30.11"
	//transaction.Amount.Currency = "USD"
	//transaction.Amount.Details = &AmountDetails{}
	//transaction.Amount.Details.Subtotal = "30.00"
	//transaction.Amount.Details.Tax = "0.07"
	//transaction.Amount.Details.Shipping = "0.03"
	//transaction.Amount.Details.HandlingFee = "1.00"
	//transaction.Amount.Details.ShippingDiscount = "-1.00"
	//transaction.Amount.Details.Insurance = "0.01"
	//
	//transaction.Description = "This is the payment transaction description."
	//transaction.Custom = "EBAY_EMS_90048630024435"
	//transaction.InvoiceNumber = uuid.New()
	//
	//transaction.PaymentOptions = &PaymentOptions{}
	//transaction.PaymentOptions.AllowedPaymentMethod = "INSTANT_FUNDING_SOURCE"
	//transaction.SoftDescriptor = "ECHI5786786"
	//
	//transaction.ItemList = &ItemList{}
	//transaction.ItemList.ShippingAddress = &ShippingAddress{}
	//transaction.ItemList.ShippingAddress.RecipientName = "Hello World"
	//transaction.ItemList.ShippingAddress.Line1 = "4thFloor"
	//transaction.ItemList.ShippingAddress.Line2 = "unit#34"
	//transaction.ItemList.ShippingAddress.City = "SAn Jose"
	//transaction.ItemList.ShippingAddress.CountryCode = "US"
	//transaction.ItemList.ShippingAddress.PostalCode = "95131"
	//transaction.ItemList.ShippingAddress.Phone = "011862212345678"
	//transaction.ItemList.ShippingAddress.State = "CA"
	//
	//var i1, i2 = &Item{}, &Item{}
	//transaction.ItemList.Items = []*Item{i1, i2}
	//
	//i1.Name = "hat"
	//i1.Description = "Brown color hat"
	//i1.Quantity = "5"
	//i1.Price = "3"
	//i1.Tax = "0.01"
	//i1.SKU = "1"
	//i1.Currency = "USD"
	//
	//i2.Name = "handbag"
	//i2.Description = "Black color hand bag"
	//i2.Quantity = "1"
	//i2.Price = "15"
	//i2.Tax = "0.02"
	//i2.SKU = "product34"
	//i2.Currency = "USD"
	//
	//p.NoteToPayer = "Contact us for any questions on your order."
	//
	//var payment, err = paypal.CreatePayment(p)
	//if err != nil {
	//	t.Fatal(err)
	//}
	//
	//if payment != nil {
	//	fmt.Println("CreatePayment", payment.Id)
	//}
}

//func TestPayPal_GetPaymentList(t *testing.T) {
//	var p = &PaymentListParam{}
//	fmt.Println(getPayPal().GetPaymentList(p))
//}
//
//func TestPayPal_GetPaymentDetails(t *testing.T) {
//	fmt.Println(getPayPal().GetPaymentDetails("PAY-1SJ16214TY566804MLB26S5I"))
//}
//
//func TestPayPal_ExecuteApprovedPayment(t *testing.T) {
//	fmt.Println(getPayPal().ExecuteApprovedPayment("PAY-0E809244MY2080201LB26LWI", "XV9HF9K25FB38"))
//}
//
//func TestPayPal_GetSaleDetails(t *testing.T) {
//	fmt.Println(getPayPal().GetSaleDetails("84E33686SW065691F"))
//}

//func TestPayPal_RefundSale(t *testing.T) {
//	fmt.Println(getPayPal().RefundSale("5SW33389HH3038001", "48787589677", "30.11", "USD"))
//}
//
//func TestPayPal_GetRefundDetails(t *testing.T) {
//	fmt.Println(getPayPal().GetRefundDetails("7AG891841V2503230"))
//}
