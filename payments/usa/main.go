package usa

import "github.com/EverID/rapyd-go-sdk/resources"

const ACHBankPaymentType = "us_ach_bank"

func ACHBank(proofOfAuthorization bool, firstName, lastName, routingNumber, paymentPurpose, accountNumber string) *resources.PaymentMethod {
	return &resources.PaymentMethod{
		Fields: map[string]interface{}{
			"proof_of_authorization": proofOfAuthorization,
			"first_name":             firstName,
			"last_name":              lastName,
			"routing_number":         routingNumber,
			"payment_purpose":        paymentPurpose,
			"account_number":         accountNumber,
		},
		Type: ACHBankPaymentType,
	}
}
