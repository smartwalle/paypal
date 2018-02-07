package paypal

import (
	"fmt"
	"net/url"
)

type DisputeStatus string

const (
	K_DISPUTE_STATUS_OPEN                        = "OPEN"
	K_DISPUTE_STATUS_WAITING_FOR_BUYER_RESPONSE  = "WAITING_FOR_BUYER_RESPONSE"
	K_DISPUTE_STATUS_WAITING_FOR_SELLER_RESPONSE = "WAITING_FOR_SELLER_RESPONSE"
	K_DISPUTE_STATUS_UNDER_REVIEW                = "UNDER_REVIEW"
	K_DISPUTE_STATUS_RESOLVED                    = "RESOLVED"
	K_DISPUTE_STATUS_OTHER                       = "OTHER"
)

type Dispute struct {
	DisputeId             string                 `json:"dispute_id,omitempty"`
	CreateTime            string                 `json:"create_time,omitempty"`
	UpdateTime            string                 `json:"update_time,omitempty"`
	DisputedTransactions  []*DisputedTransaction `json:"disputed_transactions,omitempty"`
	Reason                string                 `json:"reason,omitempty"`
	Status                DisputeStatus          `json:"status,omitempty"`
	DisputeAmount         *Currency              `json:"dispute_amount,omitempty"`
	DisputeOutcome        *DisputeOutcome        `json:"dispute_outcome,omitempty"`
	Messages              []*Message             `json:"messages,omitempty"`
	SellerResponseDueDate string                 `json:"seller_response_due_date,omitempty"`
	Links                 []*Link                `json:"links,omitempty"`
}

type DisputedTransaction struct {
	BuyerTransactionId       string                     `json:"buyer_transaction_id,omitempty"`
	SellerTransactionId      string                     `json:"seller_transaction_id,omitempty"`
	CreateTime               string                     `json:"create_time,omitempty"`
	TransactionStatus        string                     `json:"transaction_status,omitempty"`
	GrossAmount              *Currency                  `json:"gross_amount,omitempty"`
	InvoiceNumber            string                     `json:"invoice_number,omitempty"`
	Custom                   string                     `json:"custom,omitempty"`
	Buyer                    *Buyer                     `json:"buyer,omitempty"`
	Seller                   *Seller                    `json:"seller,omitempty"`
	Items                    []*DisputedTransactionItem `json:"items,omitempty"`
	SellerProtectionEligible bool                       `json:"seller_protection_eligible,omitempty"`
}

type DisputeOutcome struct {
	OutcomeCode    string    `json:"outcome_code,omitempty"`
	AmountRefunded *Currency `json:"amount_refunded,omitempty"`
}

type Message struct {
	PostedBy   string `json:"posted_by,omitempty"`
	TimePosted string `json:"time_posted,omitempty"`
	Content    string `json:"content,omitempty"`
}

type Buyer struct {
	Email string `json:"email,omitempty"`
	Name  string `json:"name,omitempty"`
}

type Seller struct {
	Email      string `json:"email,omitempty"`
	Name       string `json:"name,omitempty"`
	MerchantId string `json:"merchant_id,omitempty"`
}

type DisputedTransactionItem struct {
	ItemId               string    `json:"item_id,omitempty"`
	PartnerTransactionId string    `json:"partner_transaction_id,omitempty"`
	Reason               string    `json:"reason,omitempty"`
	DisputeAmount        *Currency `json:"dispute_amount,omitempty"`
	Notes                string    `json:"notes,omitempty"`
}

type DisputeListParam struct {
	StartTime             string
	DisputedTransactionId string
	PageSize              int
	NextPageToken         string
	//TotalRequired         bool
	DisputeState string
}

func (this *DisputeListParam) QueryString() string {
	var p = url.Values{}
	if len(this.StartTime) > 0 {
		p.Set("start_time", this.StartTime)
	}
	if len(this.DisputedTransactionId) > 0 {
		p.Set("disputed_transaction_id", this.DisputedTransactionId)
	}
	if this.PageSize > 0 {
		p.Set("page_size", fmt.Sprintf("%d", this.PageSize))
	}
	if len(this.NextPageToken) > 0 {
		p.Set("next_page_token", this.NextPageToken)
	}
	//p.Set("total_required", fmt.Sprintf("%t", this.TotalRequired))
	if len(this.DisputeState) > 0 {
		p.Set("dispute_state", this.DisputeState)
	}
	return "?" + p.Encode()
}

type DisputeList struct {
	Items      []*Dispute `json:"items"`
	TotalItems int        `json:"total_items"`
	TotalPages int        `json:"total_pages"`
	Links      []*Link    `json:"links"`
}
