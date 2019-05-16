package paypal

import "net/http"

const (
	kDisputesAPI = "/v1/customer/disputes"
)

// GetDisputeList https://developer.paypal.com/docs/api/customer-disputes/#disputes_get-disputes
func (this *Client) GetDisputeList(param *DisputeListParam) (results *DisputeList, err error) {
	var api = this.BuildAPI(kDisputesAPI) + param.QueryString()
	err = this.doRequestWithAuth(http.MethodGet, api, nil, &results)
	return results, err
}

// GetDisputeDetails https://developer.paypal.com/docs/api/customer-disputes/#disputes_get-dispute
func (this *Client) GetDisputeDetails(disputeId string) (results *Dispute, err error) {
	var api = this.BuildAPI(kDisputesAPI, disputeId)
	err = this.doRequestWithAuth(http.MethodGet, api, nil, &results)
	return results, err
}
