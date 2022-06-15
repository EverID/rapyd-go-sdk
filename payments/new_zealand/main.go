package new_zealand

import "github.com/EverID/rapyd-go-sdk/resources"

const PoliPaymentType = "nz_poli_bank"

func PoliBank() *resources.PaymentMethod {
	return &resources.PaymentMethod{
		Fields: map[string]interface{}{},
		Type:   PoliPaymentType,
	}
}
