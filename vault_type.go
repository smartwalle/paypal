package paypal

import (
	"fmt"
	"net/url"
)

type CreditCard struct {
	// Request body
	Number             string          `json:"number,omitempty"`          // required
	Type               string          `json:"type,omitempty"`            // required
	ExpireMonth        string          `json:"expire_month,omitempty"`    // required
	ExpireYear         string          `json:"expire_year,omitempty"`     // required
	BillingAddress     *BillingAddress `json:"billing_address,omitempty"` // required
	FirstName          string          `json:"first_name,omitempty"`
	LastName           string          `json:"last_name,omitempty"`
	ExternalCustomerId string          `json:"external_customer_id,omitempty"`
	MerchantId         string          `json:"merchant_id,omitempty"`
	PayerId            string          `json:"payer_id,omitempty"`
	ExternalCardId     string          `json:"external_card_id,omitempty"`

	Id         string  `json:"id,omitempty"`
	State      string  `json:"state,omitempty"`
	CreateTime string  `json:"create_time,omitempty"`
	UpdateTime string  `json:"update_time,omitempty"`
	ValidUntil string  `json:"valid_until,omitempty"`
	Links      []*Link `json:"links,omitempty"`
}

type BillingAddress struct {
	Line1               string `json:"line1,omitempty"`
	Line2               string `json:"line2,omitempty"`
	City                string `json:"city,omitempty"`
	State               string `json:"state,omitempty"`
	CountryCode         string `json:"country_code,omitempty"`
	PostalCode          string `json:"postal_code,omitempty"`
	Phone               string `json:"phone,omitempty"`
	NormalizationStatus string `json:"normalization_status,omitempty"`
	Status              string `json:"status,omitempty"`
	Type                string `json:"type,omitempty"`
}

type CreditCardToken struct {
	CreditCardId string `json:"credit_card_id,omitempty"` // required
	PayerId      string `json:"payer_id,omitempty"`
	Last4        string `json:"last4,omitempty"`
	Type         string `json:"type,omitempty"`
	ExpireMonth  string `json:"expire_month,omitempty"`
	ExpireYear   string `json:"expire_year,omitempty"`
}

type CreditCardListParam struct {
	PageSize           int
	Page               int
	StartTime          string
	EndTime            string
	SortOrder          string
	SortBy             string
	MerchantId         string
	ExternalCardId     string
	ExternalCustomerId string
	TotalRequired      bool
}

func (p *CreditCardListParam) QueryString() string {
	var v = url.Values{}
	if p.PageSize > 0 {
		v.Set("page_size", fmt.Sprintf("%d", p.PageSize))
	}
	if p.Page > 0 {
		v.Set("page", fmt.Sprintf("%d", p.Page))
	}
	if len(p.StartTime) > 0 {
		v.Set("start_time", p.StartTime)
	}
	if len(p.EndTime) > 0 {
		v.Set("end_time", p.EndTime)
	}
	if len(p.SortOrder) > 0 {
		v.Set("sort_order", p.SortOrder)
	}
	if len(p.SortBy) > 0 {
		v.Set("sort_by", p.SortBy)
	}
	if len(p.MerchantId) > 0 {
		v.Set("merchant_id", p.MerchantId)
	}
	if len(p.ExternalCardId) > 0 {
		v.Set("external_card_id", p.ExternalCardId)
	}
	if len(p.ExternalCustomerId) > 0 {
		v.Set("external_customer_id", p.ExternalCustomerId)
	}
	v.Set("total_required", fmt.Sprintf("%t", p.TotalRequired))
	return "?" + v.Encode()
}

type CreditCardList struct {
	Items      []*CreditCard `json:"items,omitempty"`
	TotalItems int           `json:"total_items,omitempty"`
	TotalPages int           `json:"total_pages,omitempty"`
	Links      []*Link       `json:"links,omitempty"`
}
