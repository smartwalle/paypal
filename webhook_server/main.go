package main

import (
	"fmt"
	"github.com/smartwalle/paypal"
	"net/http"
)

func main() {
	var pp = paypal.New("AWm6HIhK8C8XaGSdPIJ8wb0PAuaXSL9qG5Yq_wlVlZJVGH9SuuAm8goBkoM7ZLWub6VqwKm2PbM_yk8r", "EJhZ1JgU4IbriMMtSayykQoupjfZ-hET1QzOSa8Z0-tbhnwBvt4Dx14ceeDOvxxUmty_YZ-awGS_yZYY", false)
	http.HandleFunc("/paypal", func(w http.ResponseWriter, req *http.Request) {
		var event, err = pp.GetWebhookEvent("6WJ221414R474672F", req)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(event.Id, event.ResourceType)
	})

	http.ListenAndServe(":6565", nil)
}
