package resources

const (
	FixedSideSell = "sell"
	FixedSideBuy  = "buy"
)

type PaymentMethod struct {
	Fields map[string]interface{} `json:"fields"`
	Type   string                 `json:"type"`
}

type EWallet struct {
	Wallet     string `json:"ewallet"`
	Percentage int32  `json:"percentage"`
}

type CreatePayment struct {
	Amount            string    `json:"amount"`
	Currency          string    `json:"currency"`
	Customer          *string   `json:"customer,omitempty"`
	Expiration        *int64    `json:"expiration,omitempty"`
	FixedSide         *string   `json:"fixed_side,omitempty"`
	RequestedCurrency *string   `json:"requested_currency,omitempty"`
	EWallets          []EWallet `json:"ewallets"`
	PaymentMethod     string    `json:"payment_method,omitempty"`
}

type CreatePaymentResponse struct {
	Data Data `json:"data"`
}
