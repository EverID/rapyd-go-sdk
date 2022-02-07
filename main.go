package rapyd

import (
	"bytes"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/olegfomenko/rapyd-go-sdk/resources"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	createWalletPath      = "/v1/user"
	createCustomerPath    = "/v1/customers"
	createPaymentPath     = "/v1/payments"
	getPaymentFieldsPath  = "/v1/payment_methods/required_fields/"
	getPaymentMethodsPath = "/v1/payment_methods/country?country="
)

type Client interface {
	CreateCustomer(data resources.Customer) (*resources.CustomerResponse, error)
	CreateWallet(data resources.Wallet) (*resources.WalletResponse, error)
	CreatePayment(data resources.CreatePayment) (*resources.CreatePaymentResponse, error)
	GetPaymentMethodFields(method string) (*resources.PaymentMethodRequiredFieldsResponse, error)
	GetCountryPaymentMethods(country string) (*resources.CountryPaymentMethodsResponse, error)
	ValidateWebhook(r *http.Request) bool

	Resolve(path string) string
	GetSigned(path string) ([]byte, error)
	PostSigned(data interface{}, path string) ([]byte, error)
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
	path = c.Resolve(path)
	fmt.Println("Sending GET request on ", path)

	request, err := http.NewRequest("GET", path, nil)

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

	resp, err := ioutil.ReadAll(r.Body)
	fmt.Println("Got response ", string(resp))

	return resp, err
}

func (c *client) PostSigned(data interface{}, path string) ([]byte, error) {
	body, err := json.Marshal(data)
	if err != nil {
		return nil, errors.Wrap(err, "error marshalling data")
	}

	path = c.Resolve(path)
	fmt.Println("Sending POST request on ", path, " ", string(body))

	request, err := http.NewRequest("POST", path, bytes.NewBuffer(body))

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

	resp, err := ioutil.ReadAll(r.Body)
	fmt.Println("Got response ", string(resp))

	return resp, err
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
