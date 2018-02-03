package main

import (
	"fmt"
	"github.com/smartwalle/paypal"
	"net/http"
)

func main() {
	var pp = paypal.New("AS8XSa9JrOJ3rf0kxVqCgRLIlMpgaKhLTShpYxISysR1VpnN6AMLfrvj-upOMuNkXdb9bTIzsFH4umB5", "ECA3_usif2DUgGxgcBTddOKgg2rbjUT7J3B3-Ud9z9y54AK9mYTDDFyadmMLSo1QOiO2rci99FSq1PbZ", false)
	http.HandleFunc("/paypal", func(w http.ResponseWriter, req *http.Request) {
		var event, err = pp.GetWebhookEvent("7WR65199RD094230L", req)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(event.Id, event.ResourceType)
	})

	http.ListenAndServe(":6565", nil)
}
