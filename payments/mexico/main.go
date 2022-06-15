package mexico

import (
	"github.com/EverID/rapyd-go-sdk/resources"
)

const (
	SPEIPaymentType            = "mx_spei_bank"
	BBVAPaymentType            = "mx_bbva_bank"
	BanortePaymentType         = "mx_banorte_bank"
	BanamexPaymentType         = "mx_banamex_bank"
	BanorteEmpresasPaymentType = "mx_banorteempresas_bank"
	SantandermexPaymentType    = "mx_santandermex_bank"
	ScotiaPaymentType          = "mx_scotiabankmexico_bank"
	SPEIMxPaymentType          = "mx_speimx_bank"
)

func SPEI(amount string, description string, expiration int64, metadata interface{}) *resources.PaymentMethod {
	return &resources.PaymentMethod{
		Fields: map[string]interface{}{
			"amount":      amount,
			"description": description,
			"expiration":  expiration,
			"metadata":    metadata,
		},
		Type: SPEIPaymentType,
	}
}

func BBVABank() *resources.PaymentMethod {
	return &resources.PaymentMethod{
		Fields: map[string]interface{}{},
		Type:   BBVAPaymentType,
	}
}

func BanorteBank() *resources.PaymentMethod {
	return &resources.PaymentMethod{
		Fields: map[string]interface{}{},
		Type:   BanortePaymentType,
	}
}

func BanamexBank() *resources.PaymentMethod {
	return &resources.PaymentMethod{
		Fields: map[string]interface{}{},
		Type:   BanamexPaymentType,
	}
}

func BanorteEmpresasBank() *resources.PaymentMethod {
	return &resources.PaymentMethod{
		Fields: map[string]interface{}{},
		Type:   BanorteEmpresasPaymentType,
	}
}

func SantandermexBank() *resources.PaymentMethod {
	return &resources.PaymentMethod{
		Fields: map[string]interface{}{},
		Type:   SantandermexPaymentType,
	}
}

func ScotiaBank() *resources.PaymentMethod {
	return &resources.PaymentMethod{
		Fields: map[string]interface{}{},
		Type:   ScotiaPaymentType,
	}
}

func SPEIMxBank() *resources.PaymentMethod {
	return &resources.PaymentMethod{
		Fields: map[string]interface{}{},
		Type:   SPEIMxPaymentType,
	}
}
