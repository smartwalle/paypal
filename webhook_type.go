package paypal

type ResourceType string

const (
	EventResourceTypeInvoices ResourceType = "invoices"
	EventResourceTypeSale     ResourceType = "sale"
	EventResourceTypeRefund   ResourceType = "refund"
	EventResourceTypeDispute  ResourceType = "dispute"
)

// https://developer.paypal.com/docs/integration/direct/webhooks/event-names/
const (
	EventTypePaymentSaleCompleted = "PAYMENT.SALE.COMPLETED"
	EventTypePaymentSaleDenied    = "PAYMENT.SALE.DENIED"
	EventTypePaymentSalePending   = "PAYMENT.SALE.PENDING"

	// 退款成功
	EventTypePaymentSaleRefunded = "PAYMENT.SALE.REFUNDED"
	EventTypePaymentSaleReversed = "PAYMENT.SALE.REVERSED"

	// 用户从 paypal 网站申请退款
	EventTypeCustomerDisputeCreated  = "CUSTOMER.DISPUTE.CREATED"
	EventTypeCustomerDisputeResolved = "CUSTOMER.DISPUTE.RESOLVED"
	EventTypeCustomerDisputeUpdated  = "CUSTOMER.DISPUTE.UPDATED"
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

func (e *Event) Sale() *Sale {
	if s, ok := e.Resource.(*Sale); ok {
		return s
	}
	return nil
}

func (e *Event) Invoice() *Invoice {
	if s, ok := e.Resource.(*Invoice); ok {
		return s
	}
	return nil
}

func (e *Event) Dispute() *Dispute {
	if s, ok := e.Resource.(*Dispute); ok {
		return s
	}
	return nil
}

func (e *Event) Refund() *Refund {
	if s, ok := e.Resource.(*Refund); ok {
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
