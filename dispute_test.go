package paypal

import (
	"fmt"
	"testing"
)

func TestPayPal_GetDisputeList(t *testing.T) {
	var param = &DisputeListParam{}
	param.PageSize = 10
	var results, err = paypal.GetDisputeList(param)
	if err != nil {
		t.Fatal(err)
	}
	for _, dispute := range results.Items {
		fmt.Println(dispute.DisputeId, dispute.Status)
	}
}

func TestPayPal_GetDisputeDetails(t *testing.T) {
	//var results, err = paypal.GetDisputeDetails("PP-000-042-621-836")
	//if err != nil {
	//	t.Fatal(err)
	//}
	//fmt.Println(results.DisputeId, results.DisputedTransactions[0].InvoiceNumber)
}
