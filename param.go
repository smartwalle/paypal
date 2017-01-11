// https://developer.paypal.com/docs/api/payments/#definitions

package paypal

import "time"

const (
	K_PAYPAL_PAYMENT_METHOD_PAYPAL = "paypal"
	K_PAYPAL_PAYMENT_METHOD_CREDIT_CARD = "credit_card"
)

type Payer struct {
	PaymentMethod string `json:"payment_method"`
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
	Quantity    int    `json:"quantity"`
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

type Transaction struct {
	Amount         *Amount         `json:"amount"`
	Description    string          `json:"description,omitempty"`
	Custom         string          `json:"custom,omitempty"`
	InvoiceNumber  string          `json:"invoice_number,omitempty"`
	PaymentOptions *PaymentOptions `json:"payment_options,omitempty"`
	SoftDescriptor string          `json:"soft_descriptor,omitempty"`
	ItemList       *ItemList       `json:"item_list,omitempty"`

	// 返回结果添加的字段
	//RelatedResources   `json:"related_resources,omitempty"`
}

type RedirectURLs struct {
	ReturnURL string `json:"return_url"`
	CancelURL string `json:"cancel_url"`
}

type Links struct {
	Href   string `json:"href"`
	Rel    string `json:"rel"`
	Method string `json:"method"`
}

const (
	K_PAYPAL_INTENT_SALE      = "sale"
	K_PAYPAL_INTENT_AUTHORIZE = "authorize"
	K_PAYPAL_INTENT_ORDER     = "order"
)

type Payment struct {
	Intent       string         `json:"intent"`
	Payer        *Payer         `json:"payer"`
	Transactions []*Transaction `json:"transactions"`
	NoteToPayer  string         `json:"note_to_payer,omitempty"`
	RedirectURLs *RedirectURLs  `json:"redirect_urls"`

	// 返回结果添加的字段
	Id         string     `json:"id,omitempty"`
	CreateTime *time.Time `json:"create_time,omitempty"`
	State      string     `json:"state,omitempty"`
	UpdateTime *time.Time `json:"update_time,omitempty"`
	Links      []*Links   `json:"links,omitempty"`
}
