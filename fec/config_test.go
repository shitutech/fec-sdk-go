package fec

import (
	"fmt"
	"testing"
)

func TestNewConfig(t *testing.T) {
	c := NewConfig()
	c = c.SetMerchantNo("A******-******-******").
		SetProviderNo("S******").
		SetProductNo("P********").
		SetTaskCode("STC********").
		SetPrivateKey("CustomerPrivateKey")

	fmt.Println(c)
}
