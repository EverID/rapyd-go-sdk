package resources

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBeneficiaryBuilder_Build(t *testing.T) {
	bb := NewBeneficiaryBuilder()
	bb.Country("US").
		Category("bank").
		Currency("USD")

	beneficiary, err := bb.Build()

	fmt.Println(beneficiary)

	assert.Error(t, err)
}