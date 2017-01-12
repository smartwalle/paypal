// https://developer.paypal.com/docs/integration/direct/webhooks/event-names/

package paypal

import "time"

type Event struct {
	Id           string      `json:"id"`
	CreateTime   *time.Time  `json:"create_time,omitempty"`
	ResourceType string      `json:"resource_type,omitempty"`
	EventVersion string      `json:"event_version,omitempty"`
	EventType    string      `json:"event_type,omitempty"`
	Summary      string      `json:"summary,omitempty"`
	Resource     interface{} `json:"resource,omitempty"`
	Status       string      `json:"status,omitempty"`
	Links        []*Link     `json:"links,omitempty"`
}

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
