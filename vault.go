package paypal

import "net/http"

const (
	kCreditCardsAPI = "/v1/vault/credit-cards"
)

// StoreCreditCard https://developer.paypal.com/docs/api/vault/#credit-cards_create
func (this *Client) StoreCreditCard(param *CreditCard) (result *CreditCard, err error) {
	var api = this.BuildAPI(kCreditCardsAPI)
	err = this.doRequestWithAuth(http.MethodPost, api, param, &result)
	return result, err
}

// GetCreditCardList https://developer.paypal.com/docs/api/vault/#credit-cards_list
func (this *Client) GetCreditCardList(param *CreditCardListParam) (result *CreditCardList, err error) {
	var api = this.BuildAPI(kCreditCardsAPI) + param.QueryString()
	err = this.doRequestWithAuth(http.MethodGet, api, nil, &result)
	return result, err
}

// DeleteCreditCard https://developer.paypal.com/docs/api/vault/#credit-cards_delete
func (this *Client) DeleteCreditCard(creditCardId string) (err error) {
	var api = this.BuildAPI(kCreditCardsAPI, creditCardId)
	err = this.doRequestWithAuth(http.MethodDelete, api, nil, nil)
	return err
}

// GetCreditCardDetails https://developer.paypal.com/docs/api/vault/#credit-cards_delete
func (this *Client) GetCreditCardDetails(creditCardId string) (result *CreditCard, err error) {
	var api = this.BuildAPI(kCreditCardsAPI, creditCardId)
	err = this.doRequestWithAuth(http.MethodGet, api, nil, &result)
	return result, err
}
