package paypal

import "net/http"

const (
	kWebProfilesAPI = "/v1/payment-experience/web-profiles/"
)

// CreateWebExperienceProfile https://developer.paypal.com/docs/api/payment-experience/#web-profile
func (this *Client) CreateWebExperienceProfile(param *WebProfiles) (result *WebProfiles, err error) {
	var api = this.BuildAPI(kWebProfilesAPI)
	err = this.doRequestWithAuth(http.MethodPost, api, param, &result)
	return result, err
}

// GetWebExperienceProfileList https://developer.paypal.com/docs/api/payment-experience/#web-profiles_get-list
func (this *Client) GetWebExperienceProfileList() (result []*WebProfiles, err error) {
	var api = this.BuildAPI(kWebProfilesAPI)
	err = this.doRequestWithAuth(http.MethodGet, api, nil, &result)
	return result, err
}

// DeleteWebExperienceProfile https://developer.paypal.com/docs/api/payment-experience/#web-profiles_delete
func (this *Client) DeleteWebExperienceProfile(profileId string) (err error) {
	var api = this.BuildAPI(kWebProfilesAPI, profileId)
	err = this.doRequestWithAuth(http.MethodDelete, api, nil, nil)
	return err
}

// GetWebExperienceProfileDetails https://developer.paypal.com/docs/api/payment-experience/#web-profiles_get
func (this *Client) GetWebExperienceProfileDetails(profileId string) (result *WebProfiles, err error) {
	var api = this.BuildAPI(kWebProfilesAPI, profileId)
	err = this.doRequestWithAuth(http.MethodGet, api, nil, &result)
	return result, err
}

// UpdateWebExperienceProfiles https://developer.paypal.com/docs/api/payment-experience/#web-profiles_update
func (this *Client) UpdateWebExperienceProfiles(profileId string, param *WebProfiles) (err error) {
	var api = this.BuildAPI(kWebProfilesAPI, profileId)
	err = this.doRequestWithAuth(http.MethodGet, api, param, nil)
	return err
}
