package paypal

const (
	kCreditCardsAPI = "/v1/vault/credit-cards"
)

// StoreCreditCard https://developer.paypal.com/docs/api/vault/#credit-cards_create
func (this *PayPal) StoreCreditCard(param *CreditCard) (results *CreditCard, err error) {
	var api = this.BuildAPI(kCreditCardsAPI)
	err = this.doRequestWithAuth("POST", api, param, &results)
	return results, err
}

// GetCreditCardList https://developer.paypal.com/docs/api/vault/#credit-cards_list
func (this *PayPal) GetCreditCardList(param *CreditCardListParam) (results *CreditCardList, err error) {
	var api = this.BuildAPI(kCreditCardsAPI) + param.QueryString()
	err = this.doRequestWithAuth("GET", api, nil, &results)
	return results, err
}

// DeleteCreditCard https://developer.paypal.com/docs/api/vault/#credit-cards_delete
func (this *PayPal) DeleteCreditCard(creditCardId string) (err error) {
	var api = this.BuildAPI(kCreditCardsAPI, creditCardId)
	err = this.doRequestWithAuth("DELETE", api, nil, nil)
	return err
}

// GetCreditCardDetails https://developer.paypal.com/docs/api/vault/#credit-cards_delete
func (this *PayPal) GetCreditCardDetails(creditCardId string) (results *CreditCard, err error) {
	var api = this.BuildAPI(kCreditCardsAPI, creditCardId)
	err = this.doRequestWithAuth("GET", api, nil, &results)
	return results, err
}
