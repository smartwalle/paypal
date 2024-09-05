package paypal

import "net/http"

const (
	kDisputes = "/v1/customer/disputes"
)

// GetDisputeList https://developer.paypal.com/docs/api/customer-disputes/#disputes_get-disputes
func (c *Client) GetDisputeList(param *DisputeListParam) (result *DisputeList, err error) {
	var api = c.BuildAPI(kDisputes) + param.QueryString()
	err = c.doRequestWithAuth(http.MethodGet, api, nil, &result)
	return result, err
}

// GetDisputeDetails https://developer.paypal.com/docs/api/customer-disputes/#disputes_get-dispute
func (c *Client) GetDisputeDetails(disputeId string) (result *Dispute, err error) {
	var api = c.BuildAPI(kDisputes, disputeId)
	err = c.doRequestWithAuth(http.MethodGet, api, nil, &result)
	return result, err
}
