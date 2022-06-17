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
