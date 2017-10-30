package paypal

const (
	k_WEB_PROFILES_API = "/v1/payment-experience/web-profiles/"
)

// CreateWebExperienceProfile https://developer.paypal.com/docs/api/payment-experience/#web-profile
func (this *PayPal) CreateWebExperienceProfile(param *WebProfiles) (results *WebProfiles, err error) {
	var api = this.BuildAPI(k_WEB_PROFILES_API)
	err = this.doRequestWithAuth("POST", api, param, &results)
	return results, err
}
