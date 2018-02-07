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

func (this *CreditCardListParam) QueryString() string {
	var p = url.Values{}
	if this.PageSize > 0 {
		p.Set("page_size", fmt.Sprintf("%d", this.PageSize))
	}
	if this.Page > 0 {
		p.Set("page", fmt.Sprintf("%d", this.Page))
	}
	if len(this.StartTime) > 0 {
		p.Set("start_time", this.StartTime)
	}
	if len(this.EndTime) > 0 {
		p.Set("end_time", this.EndTime)
	}
	if len(this.SortOrder) > 0 {
		p.Set("sort_order", this.SortOrder)
	}
	if len(this.SortBy) > 0 {
		p.Set("sort_by", this.SortBy)
	}
	if len(this.MerchantId) > 0 {
		p.Set("merchant_id", this.MerchantId)
	}
	if len(this.ExternalCardId) > 0 {
		p.Set("external_card_id", this.ExternalCardId)
	}
	if len(this.ExternalCustomerId) > 0 {
		p.Set("external_customer_id", this.ExternalCustomerId)
	}
	p.Set("total_required", fmt.Sprintf("%t", this.TotalRequired))
	return "?" + p.Encode()
}

type CreditCardList struct {
	Items      []*CreditCard `json:"items,omitempty"`
	TotalItems int           `json:"total_items,omitempty"`
	TotalPages int           `json:"total_pages,omitempty"`
	Links      []*Link       `json:"links,omitempty"`
}
