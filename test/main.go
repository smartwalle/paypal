package main

import (
	"github.com/smartwalle/paypal"
	"fmt"
)

func main() {
	var c = paypal.NewClient("AT2V6Y2Kh7mFN5tE_c-BdeAyqS4HBcGF4Kl8seWPvA-jhY2CW6MMSr-t-mvf9F6GNFkobgtp6L2GDPxI", "EDjPl5bFyIrydUS__Nsd34l8t4O7aPEsd8Z_xY7cI_5_jTl1jt-Wn6QnzaB6-J3coWtrNoIc31DrblC3", paypal.PAY_PAL_SANDBOX_API_URL)
	fmt.Println(c.GetAccessToken())

}
