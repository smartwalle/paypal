package paypal

import "net/http"

const (
	kCreditCardsAPI = "/v1/vault/credit-cards"
)

// StoreCreditCard https://developer.paypal.com/docs/api/vault/#credit-cards_create
func (this *Client) StoreCreditCard(param *CreditCard) (results *CreditCard, err error) {
	var api = this.BuildAPI(kCreditCardsAPI)
	err = this.doRequestWithAuth(http.MethodPost, api, param, &results)
	return results, err
}

// GetCreditCardList https://developer.paypal.com/docs/api/vault/#credit-cards_list
func (this *Client) GetCreditCardList(param *CreditCardListParam) (results *CreditCardList, err error) {
	var api = this.BuildAPI(kCreditCardsAPI) + param.QueryString()
	err = this.doRequestWithAuth(http.MethodGet, api, nil, &results)
	return results, err
}

// DeleteCreditCard https://developer.paypal.com/docs/api/vault/#credit-cards_delete
func (this *Client) DeleteCreditCard(creditCardId string) (err error) {
	var api = this.BuildAPI(kCreditCardsAPI, creditCardId)
	err = this.doRequestWithAuth(http.MethodDelete, api, nil, nil)
	return err
}

// GetCreditCardDetails https://developer.paypal.com/docs/api/vault/#credit-cards_delete
func (this *Client) GetCreditCardDetails(creditCardId string) (results *CreditCard, err error) {
	var api = this.BuildAPI(kCreditCardsAPI, creditCardId)
	err = this.doRequestWithAuth(http.MethodGet, api, nil, &results)
	return results, err
}
