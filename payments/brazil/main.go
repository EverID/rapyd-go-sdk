package brazil

import "github.com/EverID/rapyd-go-sdk/resources"

const (
	BradescoPaymentType    = "br_bradesco_bank"
	BancodoPaymentType     = "br_bancodobrazil_bank"
	BancoitauPaymentType   = "br_bancoitau_bank"
	BanrisulPaymentType    = "br_banrisulbank_bank"
	BancodoCBPaymentType   = "br_bancodobrasil_cb_bank"
	BancoitauCBPaymentType = "br_bancoitau_cb_bank"
	BanrisulCBPaymentType  = "br_banrisulbank_cb_bank"
	BradescoCBPaymentType  = "br_bradesco_cb_bank"
	SantanderCBPaymentType = "br_santander_cb_bank"
	BrancodoPaymentType    = "br_brancodobrazil_bank"
	BoletoPaymentType      = "br_boleto_bank"
	ITAUPaymentType        = "br_itau_bank"
)

func BradescoBank() *resources.PaymentMethod {
	return &resources.PaymentMethod{
		Fields: map[string]interface{}{},
		Type:   BradescoPaymentType,
	}
}

func BancodoBank() *resources.PaymentMethod {
	return &resources.PaymentMethod{
		Fields: map[string]interface{}{},
		Type:   BancodoPaymentType,
	}
}

func BancoitauBank() *resources.PaymentMethod {
	return &resources.PaymentMethod{
		Fields: map[string]interface{}{},
		Type:   BancoitauPaymentType,
	}
}

func BanrisulBank() *resources.PaymentMethod {
	return &resources.PaymentMethod{
		Fields: map[string]interface{}{},
		Type:   BanrisulPaymentType,
	}
}

func BancodoCBBank() *resources.PaymentMethod {
	return &resources.PaymentMethod{
		Fields: map[string]interface{}{},
		Type:   BancodoCBPaymentType,
	}
}

func BancoitauCBBank() *resources.PaymentMethod {
	return &resources.PaymentMethod{
		Fields: map[string]interface{}{},
		Type:   BancoitauCBPaymentType,
	}
}

func BanrisulCBBank() *resources.PaymentMethod {
	return &resources.PaymentMethod{
		Fields: map[string]interface{}{},
		Type:   BanrisulCBPaymentType,
	}
}

func BradescoCBBank() *resources.PaymentMethod {
	return &resources.PaymentMethod{
		Fields: map[string]interface{}{},
		Type:   BradescoCBPaymentType,
	}
}

func SantanderCBBank() *resources.PaymentMethod {
	return &resources.PaymentMethod{
		Fields: map[string]interface{}{},
		Type:   SantanderCBPaymentType,
	}
}

func BrancodoBank(addressName, firstName, lastName, line1, city, country, zip, state, phone, email, productName, taxId string) *resources.PaymentMethod {
	return &resources.PaymentMethod{
		Fields: map[string]interface{}{
			"address_name":    addressName,
			"first_name":      firstName,
			"last_name":       lastName,
			"address_line_1":  line1,
			"address_city":    city,
			"address_country": country,
			"address_zip":     zip,
			"address_state":   state,
			"phone_number":    phone,
			"email":           email,
			"product_name":    productName,
			"cpf|cnpj":        taxId,
		},
		Type: BrancodoPaymentType,
	}
}

func BoletoBank(taxId, payeeTaxId, name string) *resources.PaymentMethod {
	return &resources.PaymentMethod{
		Fields: map[string]interface{}{
			"identification_number":       taxId,
			"payee_identification_number": payeeTaxId,
			"merchant_name":               name,
		},
		Type: BoletoPaymentType,
	}
}

func ITAUBank(addressName, firstName, lastName, line1, city, country, zip, state, phone, email, productName, taxId string) *resources.PaymentMethod {
	return &resources.PaymentMethod{
		Fields: map[string]interface{}{
			"address_name":    addressName,
			"first_name":      firstName,
			"last_name":       lastName,
			"address_line_1":  line1,
			"address_city":    city,
			"address_country": country,
			"address_zip":     zip,
			"address_state":   state,
			"phone_number":    phone,
			"email":           email,
			"product_name":    productName,
			"cpf|cnpj":        taxId,
		},
		Type: ITAUPaymentType,
	}
}
