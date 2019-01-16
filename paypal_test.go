package paypal

import (
	"os"
	"testing"
)

var paypal *PayPal

func TestMain(t *testing.M) {
	paypal = New("AS8XSa9JrOJ3rf0kxVqCgRLIlMpgaKhLTShpYxISysR1VpnN6AMLfrvj-upOMuNkXdb9bTIzsFH4umB5", "ECA3_usif2DUgGxgcBTddOKgg2rbjUT7J3B3-Ud9z9y54AK9mYTDDFyadmMLSo1QOiO2rci99FSq1PbZ", false)
	exitCode := t.Run()
	os.Exit(exitCode)
}
