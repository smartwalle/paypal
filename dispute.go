package paypal

const (
	k_DISPUTES_API = "/v1/customer/disputes"
)

// GetDisputeList https://developer.paypal.com/docs/api/customer-disputes/#disputes_get-disputes
func (this *PayPal) GetDisputeList(param *DisputeListParam) (results *DisputeList, err error) {
	var api = this.BuildAPI(k_DISPUTES_API) + param.QueryString()
	err = this.doRequestWithAuth("GET", api, nil, &results)
	return results, err
}

// GetDisputeDetails https://developer.paypal.com/docs/api/customer-disputes/#disputes_get-dispute
func (this *PayPal) GetDisputeDetails(disputeId string) (results *Dispute, err error) {
	var api = this.BuildAPI(k_DISPUTES_API, disputeId)
	err = this.doRequestWithAuth("GET", api, nil, &results)
	return results, err
}
