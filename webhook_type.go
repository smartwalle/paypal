package paypal

type ResourceType string

const (
	K_EVENT_RESOURCE_TYPE_INVOICES ResourceType = "invoices"
	K_EVENT_RESOURCE_TYPE_SALE     ResourceType = "sale"
	K_EVENT_RESOURCE_TYPE_REFUND   ResourceType = "refund"
)

// https://developer.paypal.com/docs/integration/direct/webhooks/event-names/
const (
	K_EVENT_TYPE_PAYMENT_SALE_COMPLETED = "PAYMENT.SALE.COMPLETED"
	K_EVENT_TYPE_PAYMENT_SALE_REFUNDED  = "PAYMENT.SALE.REFUNDED"
)

type EventType struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	Status      string `json:"status,omitempty"`
}

type Webhook struct {
	Id         string       `json:"id,omitempty"`
	URL        string       `json:"url"`
	EventTypes []*EventType `json:"event_types,omitempty"`
	Links      []*Link      `json:"links,omitempty"`
}

type WebhookList struct {
	Webhooks []*Webhook `json:"webhooks,omitempty"`
}

type Event struct {
	Id           string       `json:"id"`
	CreateTime   string       `json:"create_time,omitempty"`
	ResourceType ResourceType `json:"resource_type,omitempty"`
	EventVersion string       `json:"event_version,omitempty"`
	EventType    string       `json:"event_type,omitempty"`
	Summary      string       `json:"summary,omitempty"`
	Resource     interface{}  `json:"resource,omitempty"`
	Status       string       `json:"status,omitempty"`
	Links        []*Link      `json:"links,omitempty"`
}

func (this *Event) Sale() *Sale {
	if s, ok := this.Resource.(*Sale); ok {
		return s
	}
	return nil
}

func (this *Event) Invoice() *Invoice {
	if s, ok := this.Resource.(*Invoice); ok {
		return s
	}
	return nil
}

type verifyWebhookSignatureParam struct {
	AuthAlgo         string      `json:"auth_algo"`
	CertURL          string      `json:"cert_url"`
	TransmissionId   string      `json:"transmission_id"`
	TransmissionSig  string      `json:"transmission_sig"`
	TransmissionTime string      `json:"transmission_time"`
	WebhookId        string      `json:"webhook_id"`
	WebhookEvent     interface{} `json:"webhook_event"`
}

type verifyWebhookSignatureResponse struct {
	VerificationStatus string `json:"verification_status"`
}
