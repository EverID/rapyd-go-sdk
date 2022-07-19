package resources

type PayoutMethod struct {
	PayoutMethodType       string                   `json:"payout_method_type"`
	Name                   string                   `json:"name"`
	IsCancelable           int                      `json:"is_cancelable"`
	IsExpirable            int                      `json:"is_expirable"`
	IsLocationSpecific     int                      `json:"is_location_specific"`
	Status                 int                      `json:"status"`
	Image                  string                   `json:"image"`
	Category               string                   `json:"category"`
	BeneficiaryCountry     string                   `json:"beneficiary_country"`
	PayoutCurrencies       []string                 `json:"payout_currencies"`
	SenderEntityTypes      []string                 `json:"sender_entity_types"`
	BeneficiaryEntityTypes []string                 `json:"beneficiary_entity_types"`
	AmountRangePerCurrency []AmountRangePerCurrency `json:"amount_range_per_currency"`
	MinExpirationSeconds   int64                    `json:"minimum_expiration_seconds"`
	MaxExpirationSeconds   int64                    `json:"maximum_expiration_seconds"`
	SenderCurrencies       []string                 `json:"sender_currencies"`
}

type AmountRangePerCurrency struct {
	MaximumAmount  int    `json:"maximum_amount"`
	MinimumAmount  int    `json:"minimum_amount"`
	PayoutCurrency string `json:"payout_currency"`
}

type PayoutMethodsResponse struct {
	Data []PayoutMethod `json:"data"`
}

type PayoutRequiredFieldsResponse struct {
	Data PayoutRequiredFields `json:"data"`
}

type PayoutRequiredFields struct {
	SenderCurrency            string          `json:"sender_currency"`
	SenderCountry             string          `json:"sender_country"`
	SenderEntityType          string          `json:"sender_entity_type"`
	BeneficiaryCountry        string          `json:"beneficiary_country"`
	PayoutCurrency            string          `json:"payout_currency"`
	BeneficiaryEntityType     string          `json:"beneficiary_entity_type"`
	IsCancelable              int             `json:"is_cancelable"`
	IsExpirable               int             `json:"is_expirable"`
	IsLocationSpecific        int             `json:"is_location_specific"`
	IsOnline                  int             `json:"is_online"`
	Status                    int             `json:"status"`
	Image                     string          `json:"image"`
	BeneficiaryRequiredFields []RequiredField `json:"beneficiary_required_fields,omitempty"`
	SenderRequiredFields      []RequiredField `json:"sender_required_fields,omitempty"`
}

type RequiredField struct {
	Name        string  `json:"name"`
	Regex       string  `json:"regex"`
	Type        string  `json:"type"`
	Required    bool    `json:"is_required"`
	Description *string `json:"description,omitempty"`
}
