package paypal

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"errors"
	"fmt"
)

const (
	k_WEBHOOK_API                  = "/v1/notifications/webhooks"
	k_VERITY_WEBHOOK_SIGNATURE_API = "/v1/notifications/verify-webhook-signature"
)

// CreateWebhook https://developer.paypal.com/docs/api/webhooks/#webhooks_create
func (this *PayPal) CreateWebhook(callBackURL string, eventTypeList ...string) (results *Webhook, err error) {
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

// verifyWebhookSignature https://developer.paypal.com/docs/api/webhooks/#verify-webhook-signature_post
func (this *PayPal) verifyWebhookSignature(param *verifyWebhookSignatureParam) (results *verifyWebhookSignatureResponse, err error) {
	var api = this.BuildAPI(k_VERITY_WEBHOOK_SIGNATURE_API)
	err = this.doRequestWithAuth("POST", api, param, &results)
	return results, err
}

// GetWebhookEvent 用于处理 webbook 回调
func (this *PayPal) GetWebhookEvent(webhookId string, req *http.Request) (event *Event, err error) {
	req.ParseForm()
	body, err := ioutil.ReadAll(req.Body)
	if err != nil || len(body) == 0 {
		return nil, err
	}

	if err = json.Unmarshal(body, &event); err != nil {
		return nil, err
	}

	if event == nil {
		return nil, errors.New("webhook event is nil")
	}

	var verifyParam = &verifyWebhookSignatureParam{}
	verifyParam.AuthAlgo = req.Header.Get("Paypal-Auth-Algo")
	verifyParam.CertURL = req.Header.Get("Paypal-Cert-Url")
	verifyParam.TransmissionId = req.Header.Get("Paypal-Transmission-Id")
	verifyParam.TransmissionSig = req.Header.Get("Paypal-Transmission-Sig")
	verifyParam.TransmissionTime = req.Header.Get("Paypal-Transmission-Time")
	verifyParam.WebhookId = webhookId
	verifyParam.WebhookEvent = event

	verifyResp, err := this.verifyWebhookSignature(verifyParam)
	if err != nil {
		return nil, err
	}

	if verifyResp.VerificationStatus != "SUCCESS" {
		return nil, errors.New(fmt.Sprintf("verify webhook %s", verifyResp.VerificationStatus))
	}
	return event, nil
}
