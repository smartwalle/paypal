package paypal

const (
	kDisputesAPI = "/v1/customer/disputes"
)

// GetDisputeList https://developer.paypal.com/docs/api/customer-disputes/#disputes_get-disputes
func (this *PayPal) GetDisputeList(param *DisputeListParam) (results *DisputeList, err error) {
	var api = this.BuildAPI(kDisputesAPI) + param.QueryString()
	err = this.doRequestWithAuth("GET", api, nil, &results)
	return results, err
}

// GetDisputeDetails https://developer.paypal.com/docs/api/customer-disputes/#disputes_get-dispute
func (this *PayPal) GetDisputeDetails(disputeId string) (results *Dispute, err error) {
	var api = this.BuildAPI(kDisputesAPI, disputeId)
	err = this.doRequestWithAuth("GET", api, nil, &results)
	return results, err
}
