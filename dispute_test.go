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
