package resources

import (
	"errors"
	"fmt"
)

const (
	categoryKey   string = "category"
	countryKey           = "country"
	currencyKey          = "currency"
	entityTypeKey        = "entity_type"
)

type Beneficiary map[string]string

type BeneficiaryBuilder struct {
	beneficiary Beneficiary
}

func NewBeneficiaryBuilder() *BeneficiaryBuilder {
	return &BeneficiaryBuilder{beneficiary: make(map[string]string)}
}

func (b *BeneficiaryBuilder) Category(category string) *BeneficiaryBuilder {
	b.beneficiary[categoryKey] = category
	return b
}

func (b *BeneficiaryBuilder) Country(country string) *BeneficiaryBuilder {
	b.beneficiary[countryKey] = country
	return b
}

func (b *BeneficiaryBuilder) Currency(currency string) *BeneficiaryBuilder {
	b.beneficiary[currencyKey] = currency
	return b
}

func (b *BeneficiaryBuilder) EntityType(entityType string) *BeneficiaryBuilder {
	b.beneficiary[entityTypeKey] = entityType
	return b
}

func (b *BeneficiaryBuilder) Build() (*Beneficiary, error) {
	keys := []string{categoryKey, countryKey, currencyKey, entityTypeKey}

	for _, key := range keys {
		if _, ok := b.beneficiary[key]; !ok {
			return nil, errors.New(fmt.Sprintf("Required field %s is missed", key))
		}
	}

	return &b.beneficiary, nil
}

func (b Beneficiary) AddRequiredFields(key, value string) {
	b[key] = value
}

type BeneficiaryResponse struct {
	Data Data `json:"data"`
}
