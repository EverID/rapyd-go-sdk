package rapyd

import (
	"fmt"
	"github.com/EverID/rapyd-go-sdk/resources"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"testing"
	"time"
)

const (
	endpoint = "https://sandboxapi.rapyd.net"
)

var accessKey = os.Getenv("ACCESS_KEY")
var secretKey = os.Getenv("SECRET_KEY")

func TestClient_CreateWallet(t *testing.T) {
	addr, err := url.Parse(endpoint)
	assert.NoError(t, err)

	rapyd := NewClient(NewRapydSigner([]byte(accessKey), []byte(secretKey)), addr, http.DefaultClient)

	rand.Seed(time.Now().Unix())
	var randNumber = strconv.Itoa(rand.Int())

	fmt.Println(randNumber)

	_, err = rapyd.CreateWallet(resources.Wallet{
		FirstName: "Oleg",
		LastName:  "Fomenko",
		Email:     randNumber + "oleg@mail.com",
		Reference: randNumber + "-Oleg",
		Type:      resources.PersonWalletType,
		Contact: resources.Contact{
			Email:       randNumber + "oleg@mail.com",
			FirstName:   "Oleg",
			LastName:    "Fomenko",
			ContactType: resources.PersonalContactType,
			Address: resources.Address{
				Name:    "Oleg Fomenko",
				Line1:   "111 Main Street",
				City:    "Anytown",
				State:   "NY",
				Country: "US",
				Zip:     "11111",
			},
			Birth:       "11/22/2000",
			Country:     "US",
			Nationality: "FR",
		},
	})

	assert.NoError(t, err)
}

func TestClient_CreateCustomer(t *testing.T) {
	addr, err := url.Parse(endpoint)
	assert.NoError(t, err)

	rand.Seed(time.Now().Unix())
	var randNumber = strconv.Itoa(rand.Int())

	rapyd := NewClient(NewRapydSigner([]byte(accessKey), []byte(secretKey)), addr, http.DefaultClient)
	_, err = rapyd.CreateCustomer(resources.Customer{
		Name:  "Oleg Fomenko",
		Email: randNumber + "oleg@mail.com",
		PaymentMethod: resources.PaymentMethod{
			Fields: map[string]interface{}{
				"proof_of_authorization": false,
				"routing_number":         "111111111",
				"payment_purpose":        "111111",
				"account_number":         "111111",
			},
			Type: "us_ach_bank",
		},
	})

	assert.NoError(t, err)
}

func TestClient_CreatePayment(t *testing.T) {
	addr, err := url.Parse(endpoint)
	assert.NoError(t, err)

	rapyd := NewClient(NewRapydSigner([]byte(accessKey), []byte(secretKey)), addr, http.DefaultClient)
	_, err = rapyd.CreatePayment(resources.CreatePayment{
		Amount:   "100.23",
		Currency: "USD",
		PaymentMethod: &resources.PaymentMethod{
			Fields: map[string]interface{}{
				"proof_of_authorization": false,
				"routing_number":         "111111111",
				"payment_purpose":        "111111",
				"account_number":         "111111",
			},
			Type: "us_ach_bank",
		},
		EWallets: []resources.EWallet{
			{
				Wallet:     "ewallet_8d3fdd0929856f5a30ec2933f4bd6cf1",
				Percentage: 100,
			},
		},
	})

	assert.NoError(t, err)
}

func TestClient_GetPaymentMethodFields(t *testing.T) {
	addr, err := url.Parse(endpoint)
	assert.NoError(t, err)

	rapyd := NewClient(NewRapydSigner([]byte(accessKey), []byte(secretKey)), addr, http.DefaultClient)

	resp, err := rapyd.GetPaymentMethodFields("br_itau_bank")
	fmt.Println(resp)

	assert.NoError(t, err)
}

func TestClient_GetPayoutMethods(t *testing.T) {
	addr, err := url.Parse(endpoint)
	assert.NoError(t, err)

	rapyd := NewClient(NewRapydSigner([]byte(accessKey), []byte(secretKey)), addr, http.DefaultClient)

	resp, err := rapyd.GetPayoutMethods("bank", "US")
	fmt.Println(resp)

	assert.NoError(t, err)
}

func TestClient_GetPayoutRequiredFields(t *testing.T) {
	addr, err := url.Parse(endpoint)
	assert.NoError(t, err)

	rapyd := NewClient(NewRapydSigner([]byte(accessKey), []byte(secretKey)), addr, http.DefaultClient)

	resp, err := rapyd.GetPayoutRequiredFields("us_general_bank", "us", "individual",
		"251", "USD", "us", "USD", "individual")

	fmt.Println(resp)

	assert.NoError(t, err)
}

func TestClient_GetCountryPaymentMethods(t *testing.T) {
	addr, err := url.Parse(endpoint)
	assert.NoError(t, err)

	rapyd := NewClient(NewRapydSigner([]byte(accessKey), []byte(secretKey)), addr, http.DefaultClient)

	resp, err := rapyd.GetCountryPaymentMethods("US")
	fmt.Println(resp)

	assert.NoError(t, err)
}

func TestClient_CreateSender(t *testing.T) {
	addr, err := url.Parse(endpoint)
	assert.NoError(t, err)

	rapyd := NewClient(NewRapydSigner([]byte(accessKey), []byte(secretKey)), addr, http.DefaultClient)

	_, err = rapyd.CreateSender(resources.Sender{
		FirstName:               "Nikita",
		LastName:                "Shaburov",
		DateOfBirth:             "04/16/2001",
		Country:                 "US",
		Currency:                "USD",
		Address:                 "1 Second Street",
		City:                    "Montreal",
		State:                   "Quebec",
		PostCode:                "12345",
		PhoneNumber:             "0632606012",
		IdentificationType:      "identification_id",
		IdentificationValue:     "163",
		Occupation:              "occ",
		SourceOfIncome:          "salary",
		BeneficiaryRelationship: "spouse",
		PurposeCode:             "salary",
		EntityType:              "individual",
	})

	assert.NoError(t, err)
}

func TestClient_CreateBeneficiary(t *testing.T) {
	addr, err := url.Parse(endpoint)
	assert.NoError(t, err)

	rapyd := NewClient(NewRapydSigner([]byte(accessKey), []byte(secretKey)), addr, http.DefaultClient)

	_, err = rapyd.CreateBeneficiary(resources.Beneficiary{
		Category:            "rapyd_ewallet",
		Country:             "US",
		Currency:            "USD",
		EntityType:          "individual",
		FirstName:           "Nikita",
		IdentificationType:  "identification_id",
		IdentificationValue: "16345",
		LastName:            "Shaburov",
		AccountNumber:       "1234567",
		Address:             "1 Second Street",
		City:                "Montreal",
		State:               "Quebec",
		PostCode:            "12345",
	})

	assert.NoError(t, err)
}

func TestClient_CreatePayout(t *testing.T) {
	addr, err := url.Parse(endpoint)
	assert.NoError(t, err)

	rapyd := NewClient(NewRapydSigner([]byte(accessKey), []byte(secretKey)), addr, http.DefaultClient)

	_, err = rapyd.CreatePayout(resources.CreatePayout{
		Beneficiary: resources.Beneficiary{
			Category:            "bank",
			Country:             "CA",
			Currency:            "USD",
			EntityType:          "individual",
			FirstName:           "John",
			LastName:            "Doe",
			IdentificationType:  "identification_id",
			IdentificationValue: "123456789",
			PaymentType:         "priority",
			Address:             "1 Main Street",
			City:                "Montreal",
			State:               "Quebec",
			PostCode:            "12345",
			AccountNumber:       "1234567",
			BicSwift:            "12345678XXX",
		},
		BeneficiaryCountry:    "CA",
		BeneficiaryEntityType: "individual",
		PayoutAmount:          "110",
		PayoutCurrency:        "USD",
		PayoutMethodType:      "ca_general_bank",
		Sender: resources.Sender{
			Country:             "CA",
			Currency:            "USD",
			EntityType:          "individual",
			FirstName:           "Jane",
			LastName:            "Smith",
			IdentificationType:  "identification_id",
			IdentificationValue: "987654321",
			DateOfBirth:         "12/12/2000",
			Address:             "1 Second Street",
			City:                "Montreal",
			State:               "Quebec",
			PostCode:            "12345",
		},
		SenderCountry:    "CA",
		SenderCurrency:   "USD",
		SenderEntityType: "individual",
	})

	assert.NoError(t, err)
}

func TestClient_CreatePayout2(t *testing.T) {
	addr, err := url.Parse(endpoint)
	assert.NoError(t, err)

	rapyd := NewClient(NewRapydSigner([]byte(accessKey), []byte(secretKey)), addr, http.DefaultClient)
	beneficiary, err := rapyd.CreateBeneficiary(
		resources.Beneficiary{
			Category:            "bank",
			Country:             "CA",
			Currency:            "USD",
			EntityType:          "individual",
			FirstName:           "John",
			LastName:            "Doe",
			IdentificationType:  "identification_id",
			IdentificationValue: "123456789",
			PaymentType:         "priority",
			Address:             "1 Main Street",
			City:                "Montreal",
			State:               "Quebec",
			PostCode:            "12345",
			AccountNumber:       "1234567",
			BicSwift:            "12345678XXX",
		})
	assert.NoError(t, err)

	sender, err := rapyd.CreateSender(
		resources.Sender{
			Country:             "CA",
			Currency:            "USD",
			EntityType:          "individual",
			FirstName:           "Jane",
			LastName:            "Smith",
			IdentificationType:  "identification_id",
			IdentificationValue: "987654321",
			DateOfBirth:         "12/12/2000",
			Address:             "1 Second Street",
			City:                "Montreal",
			State:               "Quebec",
			PostCode:            "12345",
		})
	assert.NoError(t, err)

	_, err = rapyd.CreatePayout(resources.CreatePayout{
		Beneficiary:           beneficiary.Data.GetId(),
		BeneficiaryCountry:    "CA",
		BeneficiaryEntityType: "individual",
		PayoutAmount:          "110",
		PayoutCurrency:        "USD",
		PayoutMethodType:      "ca_general_bank",
		Sender:                sender.Data.GetId(),
		SenderCountry:         "CA",
		SenderCurrency:        "USD",
		SenderEntityType:      "individual",
	})

	assert.NoError(t, err)
}
