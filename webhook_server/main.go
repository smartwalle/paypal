package main

import (
	"fmt"
	"github.com/smartwalle/paypal"
	"net/http"
)

func main() {
	var pp = paypal.New("AS8XSa9JrOJ3rf0kxVqCgRLIlMpgaKhLTShpYxISysR1VpnN6AMLfrvj-upOMuNkXdb9bTIzsFH4umB5", "ECA3_usif2DUgGxgcBTddOKgg2rbjUT7J3B3-Ud9z9y54AK9mYTDDFyadmMLSo1QOiO2rci99FSq1PbZ", false)
	http.HandleFunc("/paypal", func(w http.ResponseWriter, req *http.Request) {
		var event, err = pp.GetWebhookEvent("6WJ221414R474672F", req)
		if err != nil {
			fmt.Println(err)
			return
		}

		if event == nil {
			return
		}
		switch event.EventType {
		case paypal.K_EVENT_TYPE_PAYMENT_SALE_COMPLETED:
			var sale = event.Sale()
			fmt.Println("支付成功", event.Id, event.ResourceType, event.EventType, sale.State, sale.ParentPayment, sale.InvoiceNumber)
		}
	})

	http.ListenAndServe(":6565", nil)
}
