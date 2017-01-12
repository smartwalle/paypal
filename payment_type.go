// https://developer.paypal.com/docs/api/payments/#definitions

package paypal

import "time"

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
	Subtotal         string `json:"subtotal"`
	Tax              string `json:"tax"`
	Shipping         string `json:"shipping"`
	HandlingFee      string `json:"handling_fee"`
	ShippingDiscount string `json:"shipping_discount"`
	Insurance        string `json:"insurance"`
}

type Amount struct {
	Total    string         `json:"total"`
	Currency string         `json:"currency"`
	Details  *AmountDetails `json:"details,omitempty"`
}

type PaymentOptions struct {
	AllowedPaymentMethod string `json:"allowed_payment_method"`
}

type Item struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Quantity    string `json:"quantity"`
	Price       string `json:"price"`
	Tax         string `json:"tax"`
	SKU         string `json:"sku"`
	Currency    string `json:"currency"`
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
	Value    string `json:"value"`
	Currency string `json:"currency"`
}

const (
	K_PAYPAL_SALE_STATUS_COMPLETED          = "completed"
	K_PAYPAL_SALE_STATUS_PARTIALLY_REFUNDED = "partially_refunded"
	K_PAYPAL_SALE_STATUS_PENDING            = "pending"
	K_PAYPAL_SALE_STATUS_REFUNDED           = "refunded"
	K_PAYPAL_SALE_STATUS_DENIED             = "denied"
)

type Sale struct {
	Id                        string          `json:"id"`
	CreateTime                *time.Time      `json:"create_time"`
	UpdateTime                *time.Time      `json:"update_time"`
	Amount                    *Amount         `json:"amount"`
	PaymentMode               string          `json:"payment_mode"`
	State                     string          `json:"state"`
	ProtectionEligibility     string          `json:"protection_eligibility"`
	ProtectionEligibilityType string          `json:"protection_eligibility_type"`
	ParentPayment             string          `json:"parent_payment"`
	TransactionFee            *TransactionFee `json:"transaction_fee,omitempty"`
	Links                     []*Link         `json:"links,omitempty"`
	InvoiceNumber             string          `json:"invoice_number"`
	Custom                    string          `json:"custom"`
	ReceiptId                 string          `json:"receipt_id"`
}

type Refund struct {
	Id            string     `json:"id"`
	CreateTime    *time.Time `json:"create_time"`
	UpdateTime    *time.Time `json:"update_time"`
	State         string     `json:"state"`
	Amount        *Amount    `json:"amount"`
	SaleId        string     `json:"sale_id"`
	ParentPayment string     `json:"parent_payment"`
	InvoiceNumber string     `json:"invoice_number"`
	Links         []*Link    `json:"links,omitempty"`
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
	Intent       string         `json:"intent"`
	Payer        *Payer         `json:"payer"`
	Transactions []*Transaction `json:"transactions"`
	NoteToPayer  string         `json:"note_to_payer,omitempty"`
	RedirectURLs *RedirectURLs  `json:"redirect_urls"`

	// 返回结果添加的字段
	Id            string     `json:"id,omitempty"`
	CreateTime    *time.Time `json:"create_time,omitempty"`
	State         string     `json:"state,omitempty"`
	FailureReason string     `json:"failure_reason,omitempty"`
	UpdateTime    *time.Time `json:"update_time,omitempty"`
	Links         []*Link    `json:"links,omitempty"`
}

type PaymentList struct {
	Payments []*Payment `json:"payments"`
	Count    int        `json:"count"`
	NextId   string     `json:"next_id"`
}