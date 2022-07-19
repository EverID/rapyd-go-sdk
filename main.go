package rapyd

import (
	"bytes"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/EverID/rapyd-go-sdk/resources"
	"github.com/pkg/errors"
)

const (
	createWalletPath              = "/v1/user"
	createCustomerPath            = "/v1/customers"
	updateCustomerPath            = "/v1/customers/"
	retrieveCustomerPath          = "/v1/customers/"
	createPaymentPath             = "/v1/payments"
	createSenderPath              = "/v1/payouts/sender"
	createPayoutPath              = "/v1/payouts"
	createBeneficiaryPath         = "/v1/payouts/beneficiary"
	getPaymentFieldsPath          = "/v1/payment_methods/required_fields/"
	getPaymentMethodsPath         = "/v1/payment_methods/country?country="
	getPayoutMethodsPath          = "v1/payouts/supported_types?"
	getPayoutFieldsPath           = "/v1/payouts/"
	updateCustomerPaymentMethod   = "/v1/customers/%s/payment_methods/"
	retrieveCustomerPaymentMethod = "/v1/customers/%s/payment_methods/"
	customerPaymentMethodList     = "/v1/customers/%s/payment_methods"
	addCustomerPaymentMethod      = "/v1/customers/%s/payment_methods"
	deleteCustomerPaymentMethod   = "/v1/customers/%s/payment_methods/"
)

type Client interface {
	CreateCustomer(data resources.Customer) (*resources.CustomerResponse, error)
	RetrieveCustomer(customerID string) (*resources.RetrieveCustomerResponse, error)
	UpdateCustomer(customerID string, data resources.Customer) (*resources.CustomerResponse, error)

	CreateWallet(data resources.Wallet) (*resources.WalletResponse, error)
	CreatePayment(data resources.CreatePayment) (*resources.CreatePaymentResponse, error)

	GetPaymentMethodFields(method string) (*resources.PaymentMethodRequiredFieldsResponse, error)
	GetCountryPaymentMethods(country string) (*resources.CountryPaymentMethodsResponse, error)

	UpdateCustomerPaymentMethod(customerID, paymentMethodID string,
		data resources.CustomerPaymentMethod) (*resources.CustomerResponse, error)
	AddCustomerPaymentMethod(customerID string, data resources.CustomerPaymentMethod) (*resources.CustomerResponse, error)
	RetrieveCustomerPaymentMethod(customerID, paymentMethodID string) (*resources.RetrieveCustomerMethod, error)
	CustomerPaymentMethodsList(customerID string) (*resources.CustomerPaymentMethodListResponse, error)
	DeleteCustomerPaymentMethod(customerID, paymentMethodID string) (*resources.CustomerResponse, error)

	CreateSender(data resources.Sender) (*resources.SenderResponse, error)
	CreateBeneficiary(data resources.Beneficiary) (*resources.BeneficiaryResponse, error)

	CreatePayout(data resources.CreatePayout) (*resources.CreatePayoutResponse, error)
	GetPayoutMethods(category, beneficiaryCountry string) (*resources.PayoutMethodsResponse, error)
	GetPayoutRequiredFields(method, beneficiaryCountry, beneficiaryEntityType, payoutAmount, payoutCurrency,
		senderCountry, senderCurrency, senderEntityType string) (*resources.PayoutRequiredFieldsResponse, error)

	ValidateWebhook(r *http.Request) bool

	Resolve(path string) string
	GetSigned(path string) ([]byte, error)
	PostSigned(data interface{}, path string) ([]byte, error)
	DeleteSigned(path string) ([]byte, error)
}

type client struct {
	Signer
	*http.Client
	url *url.URL
}

func NewClient(signer Signer, url *url.URL, cli *http.Client) Client {
	return &client{
		Signer: signer,
		Client: cli,
		url:    url,
	}
}

func (c *client) Resolve(path string) string {
	endpoint, err := url.Parse(path)
	if err != nil {
		panic(errors.New("error parsing path"))
	}
	return c.url.ResolveReference(endpoint).String()
}

func (c *client) GetSigned(path string) ([]byte, error) {
	request, err := http.NewRequest("GET", c.Resolve(path), nil)

	err = c.signRequest(request, nil)
	if err != nil {
		return nil, errors.Wrap(err, "error signing request")
	}

	r, err := c.Do(request)
	if err != nil {
		return nil, errors.Wrap(err, "error sending request")
	}
	defer r.Body.Close()

	if r.StatusCode != http.StatusOK {
		errorResponse, _ := ioutil.ReadAll(r.Body)
		return nil, errors.Errorf("error: got status code %d, response %s", r.StatusCode, string(errorResponse))
	}

	return ioutil.ReadAll(r.Body)
}

func (c *client) PostSigned(data interface{}, path string) ([]byte, error) {
	body, err := json.Marshal(data)
	if err != nil {
		return nil, errors.Wrap(err, "error marshalling data")
	}

	request, err := http.NewRequest("POST", c.Resolve(path), bytes.NewBuffer(body))

	err = c.signRequest(request, body)
	if err != nil {
		return nil, errors.Wrap(err, "error signing request")
	}

	r, err := c.Do(request)
	if err != nil {
		return nil, errors.Wrap(err, "error sending request")
	}
	defer r.Body.Close()

	if r.StatusCode != http.StatusOK {
		errorResponse, _ := ioutil.ReadAll(r.Body)
		return nil, errors.Errorf("error: got status code %d, response %s", r.StatusCode, string(errorResponse))
	}

	return ioutil.ReadAll(r.Body)
}

func (c *client) DeleteSigned(path string) ([]byte, error) {
	request, err := http.NewRequest("DELETE", c.Resolve(path), nil)

	err = c.signRequest(request, nil)
	if err != nil {
		return nil, errors.Wrap(err, "error signing request")
	}

	r, err := c.Do(request)
	if err != nil {
		return nil, errors.Wrap(err, "error sending request")
	}
	defer r.Body.Close()

	if r.StatusCode != http.StatusOK {
		errorResponse, _ := ioutil.ReadAll(r.Body)
		return nil, errors.Errorf("error: got status code %d, response %s", r.StatusCode, string(errorResponse))
	}

	return ioutil.ReadAll(r.Body)
}

func (c *client) CreateWallet(data resources.Wallet) (*resources.WalletResponse, error) {
	response, err := c.PostSigned(data, createWalletPath)
	if err != nil {
		return nil, errors.Wrap(err, "error sending create wallet request")
	}

	var body resources.WalletResponse

	err = json.Unmarshal(response, &body)
	if err != nil {
		return nil, errors.Wrap(err, "error unmarshalling response")
	}
	return &body, nil
}

func (c *client) CreateCustomer(data resources.Customer) (*resources.CustomerResponse, error) {
	response, err := c.PostSigned(data, createCustomerPath)
	if err != nil {
		return nil, errors.Wrap(err, "error sending create customer request")
	}

	var body resources.CustomerResponse

	err = json.Unmarshal(response, &body)
	if err != nil {
		return nil, errors.Wrap(err, "error unmarshalling response")
	}

	return &body, nil
}

func (c *client) RetrieveCustomer(customerID string) (*resources.RetrieveCustomerResponse, error) {
	response, err := c.GetSigned(retrieveCustomerPath + customerID)
	if err != nil {
		return nil, errors.Wrap(err, "error sending retrieve customer request")
	}

	var body resources.RetrieveCustomerResponse

	err = json.Unmarshal(response, &body)
	if err != nil {
		return nil, errors.Wrap(err, "error unmarshalling response")
	}

	return &body, nil
}

func (c *client) UpdateCustomer(customerID string, data resources.Customer) (*resources.CustomerResponse, error) {
	response, err := c.PostSigned(data, updateCustomerPath+customerID)
	if err != nil {
		return nil, errors.Wrap(err, "error sending update customer request")
	}

	var body resources.CustomerResponse

	err = json.Unmarshal(response, &body)
	if err != nil {
		return nil, errors.Wrap(err, "error unmarshalling response")
	}

	return &body, nil
}

func (c *client) UpdateCustomerPaymentMethod(customerID, paymentMethodID string, data resources.CustomerPaymentMethod) (*resources.CustomerResponse, error) {
	response, err := c.PostSigned(data, fmt.Sprintf(updateCustomerPaymentMethod, customerID)+paymentMethodID)
	if err != nil {
		return nil, errors.Wrap(err, "error sending update customer payment method request")
	}

	var body resources.CustomerResponse

	err = json.Unmarshal(response, &body)
	if err != nil {
		return nil, errors.Wrap(err, "error unmarshalling response")
	}

	return &body, nil
}

func (c *client) CustomerPaymentMethodsList(customerID string) (*resources.CustomerPaymentMethodListResponse, error) {
	response, err := c.GetSigned(fmt.Sprintf(customerPaymentMethodList, customerID))
	if err != nil {
		return nil, errors.Wrap(err, "error sending customer payment method list request")
	}

	var body resources.CustomerPaymentMethodListResponse

	err = json.Unmarshal(response, &body)
	if err != nil {
		return nil, errors.Wrap(err, "error unmarshalling response")
	}

	return &body, nil
}

func (c *client) AddCustomerPaymentMethod(customerID string, data resources.CustomerPaymentMethod) (*resources.CustomerResponse, error) {
	response, err := c.PostSigned(data, fmt.Sprintf(addCustomerPaymentMethod, customerID))
	if err != nil {
		return nil, errors.Wrap(err, "error sending add customer payment method request")
	}

	var body resources.CustomerResponse

	err = json.Unmarshal(response, &body)
	if err != nil {
		return nil, errors.Wrap(err, "error unmarshalling response")
	}

	return &body, nil
}

func (c *client) RetrieveCustomerPaymentMethod(customerID, paymentMethodID string) (*resources.RetrieveCustomerMethod, error) {
	response, err := c.GetSigned(fmt.Sprintf(retrieveCustomerPaymentMethod, customerID) + paymentMethodID)
	if err != nil {
		return nil, errors.Wrap(err, "error sending retrieve customer payment method request")
	}

	var body resources.RetrieveCustomerMethod

	err = json.Unmarshal(response, &body)
	if err != nil {
		return nil, errors.Wrap(err, "error unmarshalling response")
	}

	return &body, nil
}

func (c *client) DeleteCustomerPaymentMethod(customerID, paymentMethodID string) (*resources.CustomerResponse, error) {
	response, err := c.DeleteSigned(fmt.Sprintf(deleteCustomerPaymentMethod, customerID) + paymentMethodID)
	if err != nil {
		return nil, errors.Wrap(err, "error sending delete customer payment method request")
	}

	var body resources.CustomerResponse

	err = json.Unmarshal(response, &body)
	if err != nil {
		return nil, errors.Wrap(err, "error unmarshalling response")
	}

	return &body, nil
}

func (c *client) CreatePayment(data resources.CreatePayment) (*resources.CreatePaymentResponse, error) {
	response, err := c.PostSigned(data, createPaymentPath)
	if err != nil {
		return nil, errors.Wrap(err, "error sending create payment request")
	}

	var body resources.CreatePaymentResponse

	err = json.Unmarshal(response, &body)
	if err != nil {
		return nil, errors.Wrap(err, "error unmarshalling response")
	}

	return &body, nil
}

func (c *client) ValidateWebhook(r *http.Request) bool {
	if webhookBytes, err := ioutil.ReadAll(r.Body); err == nil {
		if err := r.Body.Close(); err != nil {
			return false
		}
		r.Body = ioutil.NopCloser(bytes.NewBuffer(webhookBytes))

		data := SignatureData{
			Path:      fmt.Sprintf("https://%s%s", r.Host, r.RequestURI),
			Salt:      r.Header.Get(SaltHeader),
			Timestamp: r.Header.Get(TimestampHeader),
			Body:      string(webhookBytes),
		}

		generatedSignature := base64.StdEncoding.EncodeToString([]byte(hex.EncodeToString(c.signData(data))))
		return generatedSignature == r.Header.Get(SignatureHeader)
	}

	return false
}

func (c *client) GetPaymentMethodFields(method string) (*resources.PaymentMethodRequiredFieldsResponse, error) {
	response, err := c.GetSigned(getPaymentFieldsPath + method)
	if err != nil {
		return nil, errors.Wrap(err, "error getting payment method fields")
	}

	var body resources.PaymentMethodRequiredFieldsResponse

	err = json.Unmarshal(response, &body)
	if err != nil {
		return nil, errors.Wrap(err, "error unmarshalling response")
	}

	return &body, nil
}

func (c *client) GetCountryPaymentMethods(country string) (*resources.CountryPaymentMethodsResponse, error) {
	response, err := c.GetSigned(getPaymentMethodsPath + country)
	if err != nil {
		return nil, errors.Wrap(err, "error getting country payment methods")
	}

	var body resources.CountryPaymentMethodsResponse

	err = json.Unmarshal(response, &body)
	if err != nil {
		return nil, errors.Wrap(err, "error unmarshalling response")
	}

	return &body, nil
}

func (c *client) GetPayoutMethods(category, beneficiaryCountry string) (*resources.PayoutMethodsResponse, error) {
	reqPath := fmt.Sprintf("%scategory=%s&beneficiary_country=%s", getPayoutMethodsPath, category, beneficiaryCountry)
	response, err := c.GetSigned(reqPath)
	if err != nil {
		return nil, errors.Wrap(err, "error getting payout methods list")
	}

	var body resources.PayoutMethodsResponse

	err = json.Unmarshal(response, &body)
	if err != nil {
		return nil, errors.Wrap(err, "error unmarshalling response")
	}

	return &body, nil
}

func (c *client) CreateSender(data resources.Sender) (*resources.SenderResponse, error) {
	response, err := c.PostSigned(data, createSenderPath)
	if err != nil {
		return nil, errors.Wrap(err, "error sending create sender request")
	}

	var body resources.SenderResponse

	err = json.Unmarshal(response, &body)
	if err != nil {
		return nil, errors.Wrap(err, "error unmarshalling response")
	}

	return &body, nil
}

func (c *client) CreateBeneficiary(data resources.Beneficiary) (*resources.BeneficiaryResponse, error) {
	response, err := c.PostSigned(data, createBeneficiaryPath)
	if err != nil {
		return nil, errors.Wrap(err, "error sending create beneficiary request")
	}

	var body resources.BeneficiaryResponse

	err = json.Unmarshal(response, &body)
	if err != nil {
		return nil, errors.Wrap(err, "error unmarshalling response")
	}

	return &body, nil
}

func (c *client) GetPayoutRequiredFields(method, beneficiaryCountry, beneficiaryEntityType, payoutAmount,
	payoutCurrency, senderCountry, senderCurrency, senderEntityType string) (*resources.PayoutRequiredFieldsResponse, error) {

	reqPath := fmt.Sprintf("%s%s/details?beneficiary_country=%s&beneficiary_entity_type=%s&payout_amount=%s&payout_currency=%s&sender_country=%s&sender_currency=%s&sender_entity_type=%s", getPayoutFieldsPath, method, beneficiaryCountry, beneficiaryEntityType, payoutAmount, payoutCurrency, senderCountry, senderCurrency, senderEntityType)

	response, err := c.GetSigned(reqPath)
	if err != nil {
		return nil, errors.Wrap(err, "error getting payment method fields")
	}

	var body resources.PayoutRequiredFieldsResponse

	responseCopy := response[:]
	err = json.Unmarshal(response, &body)
	if err != nil {
		var emptyResponseStructure struct {
			data interface{}
		}
		if err = json.Unmarshal(responseCopy, &emptyResponseStructure); err != nil {
			return nil, errors.Wrap(err, "error unmarshalling response")
		}

		return &body, nil
	}

	return &body, nil
}

func (c *client) CreatePayout(data resources.CreatePayout) (*resources.CreatePayoutResponse, error) {
	response, err := c.PostSigned(data, createPayoutPath)
	if err != nil {
		return nil, errors.Wrap(err, "error sending create beneficiary request")
	}

	var body resources.CreatePayoutResponse

	err = json.Unmarshal(response, &body)
	if err != nil {
		return nil, errors.Wrap(err, "error unmarshalling response")
	}

	return &body, nil
}
