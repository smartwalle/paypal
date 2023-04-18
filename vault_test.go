package paypal_test

import (
	"github.com/smartwalle/paypal"
	"testing"
)

func TestPayPal_StoreCreditCard(t *testing.T) {
	var p = &paypal.CreditCard{}
	p.Number = "4417119669820331"
	p.Type = "visa"
	p.ExpireMonth = "11"
	p.ExpireYear = "2024"
	p.FirstName = "Joe"
	p.LastName = "Shopper"
	p.BillingAddress = &paypal.BillingAddress{}
	p.BillingAddress.Line1 = "52 N Main St."
	p.BillingAddress.City = "Johnstown"
	p.BillingAddress.CountryCode = "US"
	p.BillingAddress.PostalCode = "43210"
	p.BillingAddress.State = "OH"
	p.BillingAddress.Phone = "408-334-8890"

	var result, err = client.StoreCreditCard(p)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result.Id)
}

func TestPayPal_GetCreditCardList(t *testing.T) {
	var p = &paypal.CreditCardListParam{}
	var result, err = client.GetCreditCardList(p)
	if err != nil {
		t.Fatal(err)
	}
	for _, item := range result.Items {
		t.Log(item.Id, item.Number)
	}
}

func TestPayPal_DeleteCreditCard(t *testing.T) {
	var err = client.DeleteCreditCard("CARD-9DS52679D6859942PLJ5GNGY")
	if err != nil {
		t.Fatal(err)
	}
}

func TestPayPal_GetCreditCardDetails(t *testing.T) {
	var result, err = client.GetCreditCardDetails("CARD-74L5398732203540KLJ5GMTQ")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result.Id, result.Number)
}
