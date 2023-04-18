package paypal_test

import (
	"testing"
)

func TestPayPal_CreateWebhook(t *testing.T) {
	var webhook, err = client.CreateWebhook("https://smartwalle.tk/paypal", "PAYMENT.SALE.COMPLETED")
	if err != nil {
		t.Fatal(err)
	}
	if webhook != nil {
		t.Log("CreateWebhook", webhook.Id, webhook.URL)
	}
}

func TestPayPal_GetWebhookList(t *testing.T) {
	var webhookList, err = client.GetWebhookList()
	if err != nil {
		t.Fatal(err)
	}
	for _, webhook := range webhookList.Webhooks {
		t.Log(webhook.Id, webhook.URL)
	}
}
