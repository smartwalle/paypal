package paypal

const (
	kWebProfilesAPI = "/v1/payment-experience/web-profiles/"
)

// CreateWebExperienceProfile https://developer.paypal.com/docs/api/payment-experience/#web-profile
func (this *PayPal) CreateWebExperienceProfile(param *WebProfiles) (results *WebProfiles, err error) {
	var api = this.BuildAPI(kWebProfilesAPI)
	err = this.doRequestWithAuth("POST", api, param, &results)
	return results, err
}

// GetWebExperienceProfileList https://developer.paypal.com/docs/api/payment-experience/#web-profiles_get-list
func (this *PayPal) GetWebExperienceProfileList() (results []*WebProfiles, err error) {
	var api = this.BuildAPI(kWebProfilesAPI)
	err = this.doRequestWithAuth("GET", api, nil, &results)
	return results, err
}

// DeleteWebExperienceProfile https://developer.paypal.com/docs/api/payment-experience/#web-profiles_delete
func (this *PayPal) DeleteWebExperienceProfile(profileId string) (err error) {
	var api = this.BuildAPI(kWebProfilesAPI, profileId)
	err = this.doRequestWithAuth("DELETE", api, nil, nil)
	return err
}

// GetWebExperienceProfileDetails https://developer.paypal.com/docs/api/payment-experience/#web-profiles_get
func (this *PayPal) GetWebExperienceProfileDetails(profileId string) (results *WebProfiles, err error) {
	var api = this.BuildAPI(kWebProfilesAPI, profileId)
	err = this.doRequestWithAuth("GET", api, nil, &results)
	return results, err
}

// UpdateWebExperienceProfiles https://developer.paypal.com/docs/api/payment-experience/#web-profiles_update
func (this *PayPal) UpdateWebExperienceProfiles(profileId string, param *WebProfiles) (err error) {
	var api = this.BuildAPI(kWebProfilesAPI, profileId)
	err = this.doRequestWithAuth("GET", api, param, nil)
	return err
}
