package paypal

type Invoice struct {
	Id                         string            `json:"id,omitempty"`
	Number                     string            `json:"number,omitempty"`
	URI                        string            `json:"uri,omitempty"`
	Status                     string            `json:"status,omitempty"`
	TemplateId                 string            `json:"template_id,omitempty"`
	MerchantInfo               *MerchantInfo     `json:"merchant_info,omitempty"`
	BillingInfo                []*BillingInfo    `json:"billing_info,omitempty"`
	ShippingInfo               *ShippingInfo     `json:"shipping_info,omitempty"`
	CCInfo                     *Participant      `json:"cc_info,omitempty"`
	Items                      []*InvoiceItem    `json:"items,omitempty"`
	InvoiceDate                string            `json:"invoice_date,omitempty"`
	PaymentTerm                *PaymentTerm      `json:"payment_term,omitempty"`
	Reference                  string            `json:"reference,omitempty"`
	Discount                   *Cost             `json:"discount,omitempty"`
	ShippingCost               *ShippingCost     `json:"shipping_cost,omitempty"`
	Custom                     *CustomAmount     `json:"custom,omitempty"`
	AllowPartialPayment        bool              `json:"allow_partial_payment,omitempty"`
	MinimumAmountDue           *Currency         `json:"minimum_amount_due,omitempty"`
	TaxCalculatedAfterDiscount bool              `json:"tax_calculated_after_discount,omitempty"`
	TaxInclusive               bool              `json:"tax_inclusive,omitempty"`
	Terms                      string            `json:"terms,omitempty"`
	Note                       string            `json:"note,omitempty"`
	MerchantMemo               string            `json:"merchant_memo,omitempty"`
	LogoURL                    string            `json:"logo_url,omitempty"`
	TotalAmount                *Currency         `json:"total_amount,omitempty"`
	Payments                   []*PaymentDetail  `json:"payments,omitempty"`
	Refunds                    []*RefundDetail   `json:"refunds,omitempty"`
	Metadata                   *Metadata         `json:"metadata,omitempty"`
	PaidAmount                 *PaymentSummary   `json:"paid_amount,omitempty"`
	RefundedAmount             *PaymentSummary   `json:"refunded_amount,omitempty"`
	Attachments                []*FileAttachment `json:"attachments,omitempty"`
	AllowTip                   bool              `json:"allow_tip,omitempty"`
	Links                      []*Link           `json:"links,omitempty"`
}

type MerchantInfo struct {
	Email          string   `json:"email,omitempty"`
	BusinessName   string   `json:"business_name,omitempty"`
	FirstName      string   `json:"first_name,omitempty"`
	LastName       string   `json:"last_name,omitempty"`
	Phone          *Phone   `json:"phone,omitempty"`
	Fax            *Phone   `json:"fax,omitempty"`
	Address        *Address `json:"address,omitempty"`
	Website        string   `json:"website,omitempty"`
	TaxId          string   `json:"tax_id,omitempty"`
	AdditionalInfo string   `json:"additional_info,omitempty"`
}

type Phone struct {
	CountryCode    string `json:"country_code,omitempty"`
	NationalNumber string `json:"national_number,omitempty"`
}

type Address struct {
	Line1       string `json:"line1,omitempty"`
	Line2       string `json:"line2,omitempty"`
	City        string `json:"city,omitempty"`
	State       string `json:"state,omitempty"`
	PostalCode  string `json:"postal_code,omitempty"`
	CountryCode string `json:"country_code,omitempty"`
	Phone       string `json:"phone,omitempty"`
}

type BillingInfo struct {
	Email        string   `json:"email,omitempty"`
	Phone        *Phone   `json:"phone,omitempty"`
	BusinessName string   `json:"business_name,omitempty"`
	FirstName    string   `json:"first_name,omitempty"`
	LastName     string   `json:"last_name,omitempty"`
	Address      *Address `json:"address,omitempty"`
	Language     string   `json:"language,omitempty"`
}

type ShippingInfo struct {
	FirstName    string   `json:"first_name,omitempty"`
	LastName     string   `json:"last_name,omitempty"`
	BusinessName string   `json:"business_name,omitempty"`
	Address      *Address `json:"address,omitempty"`
}

type Participant struct {
	Email string `json:"email,omitempty"`
}

type InvoiceItem struct {
	Name          string    `json:"name,omitempty"`
	Description   string    `json:"description,omitempty"`
	Quantity      int       `json:"quantity,omitempty"`
	UnitPrice     *Currency `json:"unit_price,omitempty"`
	Tax           *Tax      `json:"tax,omitempty"`
	Date          string    `json:"date,omitempty"`
	Discount      *Cost     `json:"discount,omitempty"`
	UnitOfMeasure string    `json:"unit_of_measure,omitempty"`
}

type PaymentTerm struct {
	TermType string `json:"term_type,omitempty"`
	DueDate  string `json:"due_date,omitempty"`
}

type Cost struct {
	Percent int       `json:"percent,omitempty"`
	Amount  *Currency `json:"amount,omitempty"`
}

type ShippingCost struct {
	Amount *Currency `json:"amount,omitempty"`
	Tax    *Tax      `json:"tax,omitempty"`
}

type CustomAmount struct {
	Label  string    `json:"label,omitempty"`
	Amount *Currency `json:"amount,omitempty"`
}

type Tax struct {
	Name    string    `json:"name,omitempty"`
	Percent int       `json:"percent,omitempty"`
	Amount  *Currency `json:"amount,omitempty"`
}

type Currency struct {
	Currency string `json:"currency,omitempty"`
	Value    string `json:"value,omitempty"`
}

type Metadata struct {
	CreatedDate     string `json:"created_date,omitempty"`
	CreatedBy       string `json:"created_by,omitempty"`
	CancelledDate   string `json:"cancelled_date,omitempty"`
	CancelledBy     string `json:"cancelled_by,omitempty"`
	LastUpdatedDate string `json:"last_updated_date,omitempty"`
	LastUpdatedBy   string `json:"last_updated_by,omitempty"`
	FirstSentDate   string `json:"first_sent_date,omitempty"`
	LastSentDate    string `json:"last_sent_date,omitempty"`
	LastSentBy      string `json:"last_sent_by,omitempty"`
	PayerViewURL    string `json:"payer_view_url,omitempty"`
}

type PaymentDetail struct {
	Type            string    `json:"type,omitempty"`
	TransactionId   string    `json:"transaction_id,omitempty"`
	TransactionType string    `json:"transaction_type,omitempty"`
	Date            string    `json:"date,omitempty"`
	Method          string    `json:"method,omitempty"`
	Note            string    `json:"note,omitempty"`
	Amount          *Currency `json:"amount,omitempty"`
}

type RefundDetail struct {
	Type          string    `json:"type,omitempty"`
	TransactionId string    `json:"transaction_id,omitempty"`
	Date          string    `json:"date,omitempty"`
	Note          string    `json:"note,omitempty"`
	Amount        *Currency `json:"amount,omitempty"`
}

type PaymentSummary struct {
	Paypal *Currency `json:"paypal,omitempty"`
	Other  *Currency `json:"other,omitempty"`
}

type FileAttachment struct {
	Name string `json:"name,omitempty"`
	URL  string `json:"url,omitempty"`
}
