package paypal

import (
	"fmt"
	"net/url"
)

type PayerInfo struct {
	Email           string           `json:"email"`
	FirstName       string           `json:"first_name"`
	LastName        string           `json:"last_name"`
	PayerId         string           `json:"payer_id"`
	ShippingAddress *ShippingAddress `json:"shipping_address,omitempty"`
}

const (
	K_PAYPAL_PAYMENT_PAYER_METHOD_PAYPAL      = "paypal"
	K_PAYPAL_PAYMENT_PAYER_METHOD_CREDIT_CARD = "credit_card"
)

type Payer struct {
	PaymentMethod string     `json:"payment_method"`
	Status        string     `json:"status,omitempty"`
	PayerInfo     *PayerInfo `json:"payer_info,omitempty"`
}

type AmountDetails struct {
	Subtotal         string `json:"subtotal,omitempty"`
	Tax              string `json:"tax,omitempty"`
	Shipping         string `json:"shipping,omitempty"`
	HandlingFee      string `json:"handling_fee,omitempty"`
	ShippingDiscount string `json:"shipping_discount,omitempty"`
	Insurance        string `json:"insurance,omitempty"`
}

type Amount struct {
	Total    string         `json:"total,omitempty"`
	Currency string         `json:"currency,omitempty"`
	Details  *AmountDetails `json:"details,omitempty"`
}

type PaymentOptions struct {
	AllowedPaymentMethod string `json:"allowed_payment_method"`
}

type Item struct {
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Quantity    interface{} `json:"quantity"` // string or int
	Price       string      `json:"price"`
	Tax         string      `json:"tax"`
	SKU         string      `json:"sku"`
	Currency    string      `json:"currency"`
}

type ShippingAddress struct {
	RecipientName string `json:"recipient_name"`
	Line1         string `json:"line1"`
	Line2         string `json:"line2"`
	City          string `json:"city"`
	CountryCode   string `json:"country_code"`
	PostalCode    string `json:"postal_code"`
	Phone         string `json:"phone"`
	State         string `json:"state"`
}

type ItemList struct {
	Items           []*Item          `json:"items,omitempty"`
	ShippingAddress *ShippingAddress `json:"shipping_address,omitempty"`
}

type TransactionFee struct {
	Value    string `json:"value,omitempty"`
	Currency string `json:"currency,omitempty"`
}

const (
	K_PAYPAL_SALE_STATUS_COMPLETED          = "completed"
	K_PAYPAL_SALE_STATUS_PARTIALLY_REFUNDED = "partially_refunded"
	K_PAYPAL_SALE_STATUS_PENDING            = "pending"
	K_PAYPAL_SALE_STATUS_REFUNDED           = "refunded"
	K_PAYPAL_SALE_STATUS_DENIED             = "denied"
)

type Sale struct {
	Id                        string          `json:"id,omitempty"`
	CreateTime                string          `json:"create_time,omitempty"`
	UpdateTime                string          `json:"update_time,omitempty"`
	Amount                    *Amount         `json:"amount,omitempty"`
	PaymentMode               string          `json:"payment_mode,omitempty"`
	State                     string          `json:"state,omitempty"`
	ProtectionEligibility     string          `json:"protection_eligibility,omitempty"`
	ProtectionEligibilityType string          `json:"protection_eligibility_type,omitempty"`
	ParentPayment             string          `json:"parent_payment,omitempty"`
	TransactionFee            *TransactionFee `json:"transaction_fee,omitempty"`
	Links                     []*Link         `json:"links,omitempty,omitempty"`
	InvoiceNumber             string          `json:"invoice_number,omitempty"`
	Custom                    string          `json:"custom,omitempty"`
	ReceiptId                 string          `json:"receipt_id,omitempty"`
	SoftDescriptor            string          `json:"soft_descriptor,omitempty"`
}

type Refund struct {
	Id            string  `json:"id"`
	CreateTime    string  `json:"create_time"`
	UpdateTime    string  `json:"update_time"`
	State         string  `json:"state"`
	Amount        *Amount `json:"amount"`
	SaleId        string  `json:"sale_id"`
	ParentPayment string  `json:"parent_payment"`
	InvoiceNumber string  `json:"invoice_number"`
	Links         []*Link `json:"links,omitempty"`
}

type RelatedResources struct {
	Sale *Sale `json:"sale,omitempty"`
}

type Transaction struct {
	Amount         *Amount         `json:"amount"`
	Description    string          `json:"description,omitempty"`
	Custom         string          `json:"custom,omitempty"`
	InvoiceNumber  string          `json:"invoice_number,omitempty"`
	PaymentOptions *PaymentOptions `json:"payment_options,omitempty"`
	SoftDescriptor string          `json:"soft_descriptor,omitempty"`
	ItemList       *ItemList       `json:"item_list,omitempty"`

	// 返回结果添加的字段
	RelatedResources []*RelatedResources `json:"related_resources,omitempty"`
}

type Payee struct {
	MerchantID string `json:"merchant_id"`
	Email      string `json:"email"`
}

type RedirectURLs struct {
	ReturnURL string `json:"return_url"`
	CancelURL string `json:"cancel_url"`
}

type Link struct {
	Href   string `json:"href"`
	Rel    string `json:"rel"`
	Method string `json:"method"`
}

const (
	K_PAYPAL_PAYMENT_INTENT_SALE      = "sale"
	K_PAYPAL_PAYMENT_INTENT_AUTHORIZE = "authorize"
	K_PAYPAL_PAYMENT_INTENT_ORDER     = "order"
)

const (
	K_PAYPAL_PAYMENT_STATUS_CREATED  = "created"
	K_PAYPAL_PAYMENT_STATUS_APPROVED = "approved"
	K_PAYPAL_PAYMENT_STATUS_FAILED   = "failed"
)

type Payment struct {
	Intent              string         `json:"intent"`
	ExperienceProfileId string         `json:"experience_profile_id,omitempty"`
	Payer               *Payer         `json:"payer"`
	Transactions        []*Transaction `json:"transactions"`
	NoteToPayer         string         `json:"note_to_payer,omitempty"`
	RedirectURLs        *RedirectURLs  `json:"redirect_urls"`

	// 返回结果添加的字段
	Id            string  `json:"id,omitempty"`
	CreateTime    string  `json:"create_time,omitempty"`
	State         string  `json:"state,omitempty"`
	FailureReason string  `json:"failure_reason,omitempty"`
	UpdateTime    string  `json:"update_time,omitempty"`
	Links         []*Link `json:"links,omitempty"`
}

type PaymentListParam struct {
	Count      int
	StartId    string
	StartIndex int
	StartTime  string
	EndTime    string
	SortBy     string
	SortOrder  string
}

func (this *PaymentListParam) QueryString() string {
	var p = url.Values{}
	if len(this.StartId) > 0 {
		p.Set("start_id", this.StartId)
	}
	if len(this.StartTime) > 0 {
		p.Set("start_time", this.StartTime)
	}
	if len(this.EndTime) > 0 {
		p.Set("end_time", this.EndTime)
	}
	if this.StartIndex > 0 {
		p.Set("start_index", fmt.Sprintf("%d", this.StartIndex))
	}
	if this.Count > 0 {
		p.Set("count", fmt.Sprintf("%f", this.Count))
	}
	if len(this.SortBy) > 0 {
		p.Set("sort_by", this.SortBy)
	}
	if len(this.SortOrder) > 0 {
		p.Set("sort_order", this.SortOrder)
	}
	return "?" + p.Encode()
}

type PaymentList struct {
	Payments []*Payment `json:"payments"`
	Count    int        `json:"count"`
	NextId   string     `json:"next_id"`
}

type refundSaleParam struct {
	Amount struct {
		Total    string `json:"total"`
		Currency string `json:"currency"`
	} `json:"amount"`
	InvoiceNumber string `json:"invoice_number"`
}
