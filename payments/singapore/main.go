package singapore

import "github.com/EverID/rapyd-go-sdk/resources"

const (
	FASTPaymentType   = "sg_fast_bank"
	PayNowPaymentType = "sg_paynow_bank"
)

func FAST(paymentPurpose string) *resources.PaymentMethod {
	return &resources.PaymentMethod{
		Fields: map[string]interface{}{
			"purpose_of_payment": paymentPurpose,
		},
		Type: FASTPaymentType,
	}
}

func PayNow(paymentPurpose string) *resources.PaymentMethod {
	return &resources.PaymentMethod{
		Fields: map[string]interface{}{
			"purpose_of_payment": paymentPurpose,
		},
		Type: PayNowPaymentType,
	}
}
