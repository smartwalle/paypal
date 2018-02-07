package paypal

import (
	"fmt"
	"testing"
)

func TestPayPal_StoreCreditCard(t *testing.T) {
	//var p = &CreditCard{}
	//p.Number = "4417119669820331"
	//p.Type = "visa"
	//p.ExpireMonth = "11"
	//p.ExpireYear = "2019"
	//p.FirstName = "Joe"
	//p.LastName = "Shopper"
	//p.BillingAddress = &BillingAddress{}
	//p.BillingAddress.Line1 = "52 N Main St."
	//p.BillingAddress.City = "Johnstown"
	//p.BillingAddress.CountryCode = "US"
	//p.BillingAddress.PostalCode = "43210"
	//p.BillingAddress.State = "OH"
	//p.BillingAddress.Phone = "408-334-8890"
	//
	//var results, err = paypal.StoreCreditCard(p)
	//if err != nil {
	//	t.Fatal(err)
	//}
	//fmt.Println(results.Id)
}

func TestPayPal_GetCreditCardList(t *testing.T) {
	//var p = &CreditCardListParam{}
	//var results, err = paypal.GetCreditCardList(p)
	//if err != nil {
	//	t.Fatal(err)
	//}
	//for _, item := range results.Items {
	//	fmt.Println(item.Id, item.Number)
	//}
}

func TestPayPal_DeleteCreditCard(t *testing.T) {
	//var err = paypal.DeleteCreditCard("CARD-9DS52679D6859942PLJ5GNGY")
	//if err != nil {
	//	t.Fatal(err)
	//}
}

func TestPayPal_GetCreditCardDetails(t *testing.T) {
	//var results, err = paypal.GetCreditCardDetails("CARD-74L5398732203540KLJ5GMTQ")
	//if err != nil {
	//	t.Fatal(err)
	//}
	//fmt.Println(results.Id, results.Number)
}
