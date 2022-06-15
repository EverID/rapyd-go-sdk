# rapyd-go-sdk

[![olegfomenko](https://circleci.com/gh/olegfomenko/rapyd-go-sdk.svg?style=shield)](https://circleci.com/gh/olegfomenko/rapyd-go-sdk)

[Rapyd](https://rapyd.net/) Golang SDK for fast and easy endpoint calls.

See [official documentation](https://docs.rapyd.net/build-with-rapyd/docs/getting-started)
and [API documentation](https://docs.rapyd.net/build-with-rapyd/reference/rapyd-overview)

All endpoint calls are secured with HMAC-SHA265 signature -
described [here](https://docs.rapyd.net/build-with-rapyd/reference/message-security)

## Example

### Setup:

```go
// getting keypair from environment
var accessKey = os.Getenv("ACCESS_KEY")
var secretKey = os.Getenv("SECRET_KEY")

url, _ := url.Parse("https://sandboxapi.rapyd.net")

// creating signer instance
signer := NewRapydSigner([]byte(accessKey), []byte(secretKey))

// creating client instance
client := NewClient(signer, url, http.DefaultClient)
```

### Creating Payment
```go
import "github.com/EverID/rapyd-go-sdk/payments/usa"
```

```go
response, err = rapyd.CreatePayment(resources.CreatePayment{
		Amount:   "100.00",
		Currency: "USD",
		PaymentMethod: usa.ACHBank(true, "Oleg", "Fomenko", "111111111", "111111", "111111"),
		EWallets: []resources.EWallet{
			{
				Wallet:     "ewallet_",
				Percentage: 100,
			},
		},
	})
```