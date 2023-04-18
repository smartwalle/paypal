package paypal_test

import (
	"github.com/smartwalle/paypal"
	"testing"
)

func TestPayPal_GetDisputeList(t *testing.T) {
	var param = &paypal.DisputeListParam{}
	param.PageSize = 10
	var result, err = client.GetDisputeList(param)
	if err != nil {
		t.Fatal(err)
	}
	for _, dispute := range result.Items {
		t.Log(dispute.DisputeId, dispute.Status)
	}

	for _, link := range result.Links {
		t.Log(link.Method, link.Rel, link.Href, link.EncType)
	}
}

func TestPayPal_GetDisputeDetails(t *testing.T) {
	var result, err = client.GetDisputeDetails("PP-000-042-621-836")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result.DisputeId, result.DisputedTransactions[0].InvoiceNumber)
}
