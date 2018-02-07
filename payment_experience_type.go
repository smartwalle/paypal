package paypal

type WebProfiles struct {
	// Request body
	Name         string                   `json:"name,omitempty"` // required
	Temporary    bool                     `json:"temporary,omitempty"`
	Presentation *WebProfilesPresentation `json:"presentation,omitempty"`
	InputFields  *WebProfilesInputFields  `json:"input_fields,omitempty"`
	FlowConfig   *WebProfilesFlowConfig   `json:"flow_config,omitempty"`

	Id string `json:"id,omitempty"`
}

type WebProfilesPresentation struct {
	BrandName  string `json:"brand_name,omitempty"`
	LogoImage  string `json:"logo_image,omitempty"`
	LocaleCode string `json:"locale_code,omitempty"`
}

type WebProfilesInputFields struct {
	NoShipping      int `json:"no_shipping,omitempty"`
	AddressOverride int `json:"address_override,omitempty"`
}

type WebProfilesFlowConfig struct {
	LandingPageType     string `json:"landing_page_type,omitempty"`
	BankTxnPendingURL   string `json:"bank_txn_pending_url,omitempty"`
	UserAction          string `json:"user_action,omitempty"`
	ReturnURIHttpMethod string `json:"return_uri_http_method,omitempty"`
}
