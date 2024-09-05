package paypal

import (
	"fmt"
	"net/url"
)

type PayerInfo struct {
	Email           string           `json:"email,omitempty"`
	Salutation      string           `json:"salutation,omitempty"`
	FirstName       string           `json:"first_name,omitempty"`
	MiddleName      string           `json:"middle_name,omitempty"`
	LastName        string           `json:"last_name,omitempty"`
	Suffix          string           `json:"suffix,omitempty"`
	PayerId         string           `json:"payer_id,omitempty"`
	Phone           string           `json:"phone,omitempty"`
	PhoneType       string           `json:"phone_type,omitempty"`
	BirthDate       string           `json:"birth_date,omitempty"`
	TaxId           string           `json:"tax_id,omitempty"`
	TaxIdType       string           `json:"tax_id_type,omitempty"`
	CountryCode     string           `json:"country_code,omitempty"`
	BillingAddress  *BillingAddress  `json:"billing_address,omitempty"`
	ShippingAddress *ShippingAddress `json:"shipping_address,omitempty"`
}

type PaymentMethod string

const (
	PaymentMethodPayPal     PaymentMethod = "paypal"
	PaymentMethodCreditCard PaymentMethod = "credit_card"
)

type Payer struct {
	PaymentMethod     PaymentMethod      `json:"payment_method,omitempty"`
	Status            string             `json:"status,omitempty"`
	PayerInfo         *PayerInfo         `json:"payer_info,omitempty"`
	FundingInstrument *FundingInstrument `json:"funding_instrument,omitempty"`
}

type FundingInstrument struct {
	CreditCard      *CreditCard      `json:"credit_card,omitempty"`
	CreditCardToken *CreditCardToken `json:"credit_card_token,omitempty"`
}

type AmountDetails struct {
	Subtotal         string `json:"subtotal,omitempty"`
	Shipping         string `json:"shipping,omitempty"`
	Tax              string `json:"tax,omitempty"`
	HandlingFee      string `json:"handling_fee,omitempty"`
	ShippingDiscount string `json:"shipping_discount,omitempty"`
	Insurance        string `json:"insurance,omitempty"`
	GiftWrap         string `json:"gift_wrap,omitempty"`
}

type Amount struct {
	Total    string         `json:"total,omitempty"`    // required
	Currency string         `json:"currency,omitempty"` // required
	Details  *AmountDetails `json:"details,omitempty"`
}

type PaymentOptions struct {
	AllowedPaymentMethod string `json:"allowed_payment_method,omitempty"`
}

type Item struct {
	Name        string      `json:"name,omitempty"`
	Description string      `json:"description,omitempty"`
	Quantity    interface{} `json:"quantity,omitempty"` // string or int
	Price       string      `json:"price,omitempty"`
	Tax         string      `json:"tax,omitempty"`
	SKU         string      `json:"sku,omitempty"`
	Currency    string      `json:"currency,omitempty"`
}

type ShippingAddress struct {
	RecipientName string `json:"recipient_name,omitempty"`
	Line1         string `json:"line1,omitempty"`
	Line2         string `json:"line2,omitempty"`
	City          string `json:"city,omitempty"`
	CountryCode   string `json:"country_code,omitempty"`
	PostalCode    string `json:"postal_code,omitempty"`
	Phone         string `json:"phone,omitempty"`
	State         string `json:"state,omitempty"`
}

type ItemList struct {
	Items               []*Item          `json:"items,omitempty"`
	ShippingAddress     *ShippingAddress `json:"shipping_address,omitempty"`
	ShippingMethod      string           `json:"shipping_method,omitempty"`
	ShippingPhoneNumber string           `json:"shipping_phone_number,omitempty"`
}

type SaleState string

const (
	SaleStateCompleted         SaleState = "completed"
	SaleStatePartiallyRefunded SaleState = "partially_refunded"
	SaleStatePending           SaleState = "pending"
	SaleStateRefunded          SaleState = "refunded"
	SaleStateDenied            SaleState = "denied"
)

type Sale struct {
	Id                        string              `json:"id,omitempty"`
	PurchaseUnitReferenceId   string              `json:"purchase_unit_reference_id,omitempty"`
	Amount                    *Amount             `json:"amount,omitempty"`
	PaymentMode               string              `json:"payment_mode,omitempty"`
	State                     SaleState           `json:"state,omitempty"`
	ReasonCode                string              `json:"reason_code,omitempty"`
	ProtectionEligibility     string              `json:"protection_eligibility,omitempty"`
	ProtectionEligibilityType string              `json:"protection_eligibility_type,omitempty"`
	ClearingTime              string              `json:"clearing_time,omitempty"`
	PaymentHoldStatus         string              `json:"payment_hold_status,omitempty"`
	PaymentHoldReasons        []PaymentHoldReason `json:"payment_hold_reasons,omitempty"`
	TransactionFee            *Currency           `json:"transaction_fee,omitempty"`
	ReceivableAmount          *Currency           `json:"receivable_amount,omitempty"`
	ExchangeRate              string              `json:"exchange_rate,omitempty"`
	FMFDetails                *FMFDetails         `json:"fmf_details,omitempty"`
	ReceiptId                 string              `json:"receipt_id,omitempty"`
	ParentPayment             string              `json:"parent_payment,omitempty"`
	ProcessorResponse         *ProcessorResponse  `json:"processor_response,omitempty"`
	BillingAgreementId        string              `json:"billing_agreement_id,omitempty"`
	CreateTime                string              `json:"create_time,omitempty"`
	UpdateTime                string              `json:"update_time,omitempty"`
	Links                     []*Link             `json:"links,omitempty,omitempty"`
	InvoiceNumber             string              `json:"invoice_number,omitempty"`
	Custom                    string              `json:"custom,omitempty"`
	SoftDescriptor            string              `json:"soft_descriptor,omitempty"`
}

type PaymentHoldReason struct {
	PaymentHoldReason string `json:"payment_hold_reason,omitempty"`
}

type FMFDetails struct {
	FilterType  string `json:"filter_type,omitempty"`
	FilterId    string `json:"filter_id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

type ProcessorResponse struct {
	ResponseCode string `json:"response_code,omitempty"`
	AVSCode      string `json:"avs_code,omitempty"`
	CVVCode      string `json:"cvv_code,omitempty"`
	AdviceCode   string `json:"advice_code,omitempty"`
	ECISubmitted string `json:"eci_submitted,omitempty"`
	Vpas         string `json:"vpas,omitempty"`
}

type RefundState string

const (
	RefundStatePending   RefundState = "pending"
	RefundStateCompleted RefundState = "completed"
	RefundStateCancelled RefundState = "cancelled"
	RefundStateFailed    RefundState = "failed"
)

type Refund struct {
	Id               string      `json:"id,omitempty"`
	Amount           *Amount     `json:"amount,omitempty"`
	State            RefundState `json:"state,omitempty"`
	Reason           string      `json:"reason,omitempty"`
	RefundReasonCode string      `json:"refund_reason_code,omitempty"`
	InvoiceNumber    string      `json:"invoice_number,omitempty"`
	SaleId           string      `json:"sale_id,omitempty"`
	CaptureId        string      `json:"capture_id,omitempty"`
	ParentPayment    string      `json:"parent_payment,omitempty"`
	Description      string      `json:"description,omitempty"`
	CreateTime       string      `json:"create_time,omitempty"`
	UpdateTime       string      `json:"update_time,omitempty"`
	Custom           string      `json:"custom,omitempty"`
	RefundToPayer    *Currency   `json:"refund_to_payer,omitempty"`
	Links            []*Link     `json:"links,omitempty,omitempty"`
}

type Transaction struct {
	ReferenceId    string          `json:"reference_id,omitempty"`
	Amount         *Amount         `json:"amount,omitempty"` // required
	Payee          *Payee          `json:"payee,omitempty"`
	Description    string          `json:"description,omitempty"`
	NoteToPayee    string          `json:"note_to_payee,omitempty"`
	Custom         string          `json:"custom,omitempty"`
	InvoiceNumber  string          `json:"invoice_number,omitempty"`
	PurchaseOrder  string          `json:"purchase_order,omitempty"`
	SoftDescriptor string          `json:"soft_descriptor,omitempty"`
	PaymentOptions *PaymentOptions `json:"payment_options,omitempty"`
	ItemList       *ItemList       `json:"item_list,omitempty"`
	NotifyURL      string          `json:"notify_url,omitempty"`
	OrderURL       string          `json:"order_url,omitempty"`
	// 返回结果添加的字段
	RelatedResources []*RelatedResources `json:"related_resources,omitempty"`
}

type Payee struct {
	MerchantID string `json:"merchant_id,omitempty"`
	Email      string `json:"email,omitempty"`
}

type RedirectURLs struct {
	ReturnURL string `json:"return_url,omitempty"`
	CancelURL string `json:"cancel_url,omitempty"`
}

type PaymentIntent string

const (
	PaymentIntentSale      PaymentIntent = "sale"
	PaymentIntentAuthorize PaymentIntent = "authorize"
	PaymentIntentOrder     PaymentIntent = "order"
)

type PaymentState string

const (
	PaymentStateCreated  PaymentState = "created"
	PaymentStateApproved PaymentState = "approved"
	PaymentStateFailed   PaymentState = "failed"
)

type Payment struct {
	// Request body
	Intent              PaymentIntent  `json:"intent,omitempty"`       // required
	Payer               *Payer         `json:"payer,omitempty"`        // required
	Transactions        []*Transaction `json:"transactions,omitempty"` // required
	ExperienceProfileId string         `json:"experience_profile_id,omitempty"`
	NoteToPayer         string         `json:"note_to_payer,omitempty"`
	RedirectURLs        *RedirectURLs  `json:"redirect_urls"`

	// 返回结果添加的字段
	Id            string       `json:"id,omitempty"`
	CreateTime    string       `json:"create_time,omitempty"`
	State         PaymentState `json:"state,omitempty"`
	FailureReason string       `json:"failure_reason,omitempty"`
	UpdateTime    string       `json:"update_time,omitempty"`
	Links         []*Link      `json:"links,omitempty"`
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

func (p *PaymentListParam) QueryString() string {
	var v = url.Values{}
	if len(p.StartId) > 0 {
		v.Set("start_id", p.StartId)
	}
	if len(p.StartTime) > 0 {
		v.Set("start_time", p.StartTime)
	}
	if len(p.EndTime) > 0 {
		v.Set("end_time", p.EndTime)
	}
	if p.StartIndex > 0 {
		v.Set("start_index", fmt.Sprintf("%d", p.StartIndex))
	}
	if p.Count > 0 {
		v.Set("count", fmt.Sprintf("%d", p.Count))
	}
	if len(p.SortBy) > 0 {
		v.Set("sort_by", p.SortBy)
	}
	if len(p.SortOrder) > 0 {
		v.Set("sort_order", p.SortOrder)
	}
	return "?" + v.Encode()
}

type PaymentList struct {
	Payments []*Payment `json:"payments,omitempty"`
	Count    int        `json:"count,omitempty"`
	NextId   string     `json:"next_id,omitempty"`
}

type RefundSaleParam struct {
	Amount        *Amount `json:"amount,omitempty"`
	Description   string  `json:"description,omitempty"`
	Reason        string  `json:"reason,omitempty"`
	InvoiceNumber string  `json:"invoice_number,omitempty"`
}
