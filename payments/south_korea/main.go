package south_korea

import "github.com/EverID/rapyd-go-sdk/resources"

const LocalPaymentType = "kr_localredirect_bank"

func LocalBank(isMobile bool) *resources.PaymentMethod {
	return &resources.PaymentMethod{
		Fields: map[string]interface{}{
			"is_mobile": isMobile,
		},
		Type: LocalPaymentType,
	}
}
