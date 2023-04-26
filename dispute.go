package paypal

import "net/http"

const (
	kDisputes = "/v1/customer/disputes"
)

// GetDisputeList https://developer.paypal.com/docs/api/customer-disputes/#disputes_get-disputes
func (this *Client) GetDisputeList(param *DisputeListParam) (result *DisputeList, err error) {
	var api = this.BuildAPI(kDisputes) + param.QueryString()
	err = this.doRequestWithAuth(http.MethodGet, api, nil, &result)
	return result, err
}

// GetDisputeDetails https://developer.paypal.com/docs/api/customer-disputes/#disputes_get-dispute
func (this *Client) GetDisputeDetails(disputeId string) (result *Dispute, err error) {
	var api = this.BuildAPI(kDisputes, disputeId)
	err = this.doRequestWithAuth(http.MethodGet, api, nil, &result)
	return result, err
}
