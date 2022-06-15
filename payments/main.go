package payments

import "github.com/EverID/rapyd-go-sdk/resources"

func CustomPayment(fields map[string]interface{}, paymentType string) *resources.PaymentMethod {
	return &resources.PaymentMethod{
		Fields: fields,
		Type:   paymentType,
	}
}
