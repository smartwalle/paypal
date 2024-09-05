package paypal

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

const (
	kWebHook                = "/v1/notifications/webhooks"
	kVerityWebHookSignature = "/v1/notifications/verify-webhook-signature"
)

// CreateWebhook https://developer.paypal.com/docs/api/webhooks/#webhooks_create
func (c *Client) CreateWebhook(callBackURL string, eventTypeList ...string) (result *Webhook, err error) {
	var api = c.BuildAPI(kWebHook)
	var p = &Webhook{}

	var events = make([]*EventType, 0, len(eventTypeList))
	for _, name := range eventTypeList {
		var event = &EventType{}
		event.Name = name
		events = append(events, event)
	}
	p.EventTypes = events

	p.URL = callBackURL
	err = c.doRequestWithAuth(http.MethodPost, api, p, &result)
	return result, err
}

// GetWebhookList https://developer.paypal.com/docs/api/webhooks/#webhooks_get-all
func (c *Client) GetWebhookList() (result *WebhookList, err error) {
	var api = c.BuildAPI(kWebHook)
	err = c.doRequestWithAuth(http.MethodGet, api, nil, &result)
	return result, err
}

// GetWebhookDetails https://developer.paypal.com/docs/api/webhooks/#webhooks_get
func (c *Client) GetWebhookDetails(webhookId string) (result *Webhook, err error) {
	var api = c.BuildAPI(kWebHook, webhookId)
	err = c.doRequestWithAuth(http.MethodGet, api, nil, &result)
	return result, err
}

// DeleteWebhook https://developer.paypal.com/docs/api/webhooks/#webhooks_delete
func (c *Client) DeleteWebhook(webhookId string) (err error) {
	var api = c.BuildAPI(kWebHook, webhookId)
	err = c.doRequestWithAuth(http.MethodDelete, api, nil, nil)
	return err
}

// verifyWebhookSignature https://developer.paypal.com/docs/api/webhooks/#verify-webhook-signature_post
func (c *Client) verifyWebhookSignature(param *verifyWebhookSignatureParam) (result *verifyWebhookSignatureResponse, err error) {
	var api = c.BuildAPI(kVerityWebHookSignature)
	err = c.doRequestWithAuth(http.MethodPost, api, param, &result)
	return result, err
}

// GetWebhookEvent 用于处理 webbook 回调
func (c *Client) GetWebhookEvent(webhookId string, req *http.Request) (event *Event, err error) {
	req.ParseForm()
	body, err := io.ReadAll(req.Body)
	if err != nil || len(body) == 0 {
		return nil, err
	}

	var rawRes json.RawMessage
	event = &Event{
		Resource: &rawRes,
	}

	if err = json.Unmarshal(body, &event); err != nil {
		return nil, err
	}

	if event == nil || (event != nil && (event.Id == "" || event.EventType == "")) {
		return nil, errors.New("unknown webhook event")
	}

	switch event.ResourceType {
	case EventResourceTypeSale:
		var sale *Sale
		if err = json.Unmarshal(rawRes, &sale); err != nil {
			return nil, err
		}
		event.Resource = sale
	case EventResourceTypeRefund:
		var refund *Refund
		if err = json.Unmarshal(rawRes, &refund); err != nil {
			return nil, err
		}
		event.Resource = refund
	case EventResourceTypeInvoices:
		var invoice *Invoice
		if err = json.Unmarshal(rawRes, &invoice); err != nil {
			return nil, err
		}
		event.Resource = invoice
	case EventResourceTypeDispute:
		var dispute *Dispute
		if err = json.Unmarshal(rawRes, &dispute); err != nil {
			return nil, err
		}
		event.Resource = dispute
	default:
		var data map[string]interface{}
		if err = json.Unmarshal(rawRes, &data); err != nil {
			return nil, err
		}
		event.Resource = data
	}

	var verifyParam = &verifyWebhookSignatureParam{}
	verifyParam.AuthAlgo = req.Header.Get("Paypal-Auth-Algo")
	verifyParam.CertURL = req.Header.Get("Paypal-Cert-Url")
	verifyParam.TransmissionId = req.Header.Get("Paypal-Transmission-Id")
	verifyParam.TransmissionSig = req.Header.Get("Paypal-Transmission-Sig")
	verifyParam.TransmissionTime = req.Header.Get("Paypal-Transmission-Time")
	verifyParam.WebhookId = webhookId
	verifyParam.WebhookEvent = jsonString(body)

	verifyResp, err := c.verifyWebhookSignature(verifyParam)
	if err != nil {
		return nil, err
	}

	if verifyResp.VerificationStatus != kSuccess {
		return nil, errors.New(fmt.Sprintf("verify webhook %s", verifyResp.VerificationStatus))
	}

	return event, nil
}
