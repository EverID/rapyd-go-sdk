package rapyd

import (
	"bytes"
	"encoding/json"
	"github.com/olegfomenko/rapyd-go-sdk/resources"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	createWalletPath   = "/v1/user"
	createCustomerPath = "/v1/customers"
	createPaymentPath  = "/v1/payments"
)

type Client interface {
	CreateCustomer(data resources.Customer) (*resources.CustomerResponse, error)
	CreateWallet(data resources.Wallet) (*resources.WalletResponse, error)
	CreatePayment(data resources.CreatePayment) (*resources.CreatePaymentResponse, error)
	ValidateWebhook(wh resources.Webhook, path string) bool

	Resolve(path string) string
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
		return nil, errors.Wrap(err, "error sending create wallet request")
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
		return nil, errors.Wrap(err, "error sending create wallet request")
	}

	var body resources.CreatePaymentResponse

	err = json.Unmarshal(response, &body)
	if err != nil {
		return nil, errors.Wrap(err, "error unmarshalling response")
	}

	return &body, nil
}

func (c *client) ValidateWebhook(wh resources.Webhook, path string) bool {
	h := wh.Headers

	data := SignatureData{
		Path:      path,
		Salt:      h.Get(SaltHeader),
		Timestamp: h.Get(TimestampHeader),
		Body:      string(wh.Body),
	}

	return string(c.signData(data)) == h.Get(SignatureHeader)
}
