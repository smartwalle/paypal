package paypal

const (
	k_WEBHOOK_API = "/v1/notifications/webhooks"
)

// CreateWebhook https://developer.paypal.com/docs/api/webhooks/#webhooks_create
func (this *PayPal) CreateWebhook(callBackURL string, eventTypeList []string) (results *Webhook, err error) {
	var api = this.BuildAPI(k_WEBHOOK_API)
	var p = &Webhook{}

	var events = make([]*EventType, 0, len(eventTypeList))
	for _, name := range eventTypeList {
		var event = &EventType{}
		event.Name = name
		events = append(events, event)
	}
	p.EventTypes = events

	p.URL = callBackURL
	err = this.doRequestWithAuth("POST", api, p, &results)
	return results, err
}

// GetWebhookList https://developer.paypal.com/docs/api/webhooks/#webhooks_get-all
func (this *PayPal) GetWebhookList() (results *WebhookList, err error) {
	var api = this.BuildAPI(k_WEBHOOK_API)
	err = this.doRequestWithAuth("GET", api, nil, &results)
	return results, err
}

// GetWebhookDetails https://developer.paypal.com/docs/api/webhooks/#webhooks_get
func (this *PayPal) GetWebhookDetails(webhookId string) (results *Webhook, err error) {
	var api = this.BuildAPI(k_WEBHOOK_API, webhookId)
	err = this.doRequestWithAuth("GET", api, nil, &results)
	return results, err
}

// DeleteWebhook https://developer.paypal.com/docs/api/webhooks/#webhooks_delete
func (this *PayPal) DeleteWebhook(webhookId string) (results *Webhook, err error) {
	var api = this.BuildAPI(k_WEBHOOK_API, webhookId)
	err = this.doRequestWithAuth("DELETE", api, nil, &results)
	return results, err
}
