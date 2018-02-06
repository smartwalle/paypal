package paypal

type WebProfiles struct {
	Id           string                   `json:"id,omitempty"`
	Name         string                   `json:"name,omitempty"`
	Presentation *WebProfilesPresentation `json:"presentation,omitempty"`
	InputFields  *WebProfilesInputFields  `json:"input_fields,omitempty"`
	FlowConfig   *WebProfilesFlowConfig   `json:"flow_config,omitempty"`
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
	LandingPageType   string `json:"landing_page_type,omitempty"`
	BankTxnPendingURL string `json:"bank_txn_pending_url,omitempty"`
}
