package paypal

import "net/http"

const (
	kCreditCards = "/v1/vault/credit-cards"
)

// StoreCreditCard https://developer.paypal.com/docs/api/vault/#credit-cards_create
func (c *Client) StoreCreditCard(param *CreditCard) (result *CreditCard, err error) {
	var api = c.BuildAPI(kCreditCards)
	err = c.doRequestWithAuth(http.MethodPost, api, param, &result)
	return result, err
}

// GetCreditCardList https://developer.paypal.com/docs/api/vault/#credit-cards_list
func (c *Client) GetCreditCardList(param *CreditCardListParam) (result *CreditCardList, err error) {
	var api = c.BuildAPI(kCreditCards) + param.QueryString()
	err = c.doRequestWithAuth(http.MethodGet, api, nil, &result)
	return result, err
}

// DeleteCreditCard https://developer.paypal.com/docs/api/vault/#credit-cards_delete
func (c *Client) DeleteCreditCard(creditCardId string) (err error) {
	var api = c.BuildAPI(kCreditCards, creditCardId)
	err = c.doRequestWithAuth(http.MethodDelete, api, nil, nil)
	return err
}

// GetCreditCardDetails https://developer.paypal.com/docs/api/vault/#credit-cards_delete
func (c *Client) GetCreditCardDetails(creditCardId string) (result *CreditCard, err error) {
	var api = c.BuildAPI(kCreditCards, creditCardId)
	err = c.doRequestWithAuth(http.MethodGet, api, nil, &result)
	return result, err
}
