package paypal

import "net/http"

const (
	kWebProfiles = "/v1/payment-experience/web-profiles/"
)

// CreateWebExperienceProfile https://developer.paypal.com/docs/api/payment-experience/#web-profile
func (c *Client) CreateWebExperienceProfile(param *WebProfiles) (result *WebProfiles, err error) {
	var api = c.BuildAPI(kWebProfiles)
	err = c.doRequestWithAuth(http.MethodPost, api, param, &result)
	return result, err
}

// GetWebExperienceProfileList https://developer.paypal.com/docs/api/payment-experience/#web-profiles_get-list
func (c *Client) GetWebExperienceProfileList() (result []*WebProfiles, err error) {
	var api = c.BuildAPI(kWebProfiles)
	err = c.doRequestWithAuth(http.MethodGet, api, nil, &result)
	return result, err
}

// DeleteWebExperienceProfile https://developer.paypal.com/docs/api/payment-experience/#web-profiles_delete
func (c *Client) DeleteWebExperienceProfile(profileId string) (err error) {
	var api = c.BuildAPI(kWebProfiles, profileId)
	err = c.doRequestWithAuth(http.MethodDelete, api, nil, nil)
	return err
}

// GetWebExperienceProfileDetails https://developer.paypal.com/docs/api/payment-experience/#web-profiles_get
func (c *Client) GetWebExperienceProfileDetails(profileId string) (result *WebProfiles, err error) {
	var api = c.BuildAPI(kWebProfiles, profileId)
	err = c.doRequestWithAuth(http.MethodGet, api, nil, &result)
	return result, err
}

// UpdateWebExperienceProfiles https://developer.paypal.com/docs/api/payment-experience/#web-profiles_update
func (c *Client) UpdateWebExperienceProfiles(profileId string, param *WebProfiles) (err error) {
	var api = c.BuildAPI(kWebProfiles, profileId)
	err = c.doRequestWithAuth(http.MethodGet, api, param, nil)
	return err
}
