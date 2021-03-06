package contracts_test

import (
	"testing"

	"github.com/onflow/flow-go-sdk"
	"github.com/stretchr/testify/assert"

	"github.com/onflow/flow-ft/lib/go/contracts"
)

var addrA = flow.HexToAddress("0A")

func TestFungibleTokenContract(t *testing.T) {
	contract := contracts.FungibleToken()
	assert.NotNil(t, contract)
}

func TestExampleTokenContract(t *testing.T) {
	contract := contracts.ExampleToken(addrA.Hex())
	assert.NotNil(t, contract)
	assert.Contains(t, string(contract), addrA.Hex())
}

func TestCustomExampleTokenContract(t *testing.T) {
	contract := contracts.CustomToken(addrA.Hex(), "UtilityCoin", "utilityCoin", "100.0")
	assert.NotNil(t, contract)
	assert.Contains(t, string(contract), addrA.Hex())
}

func TestTokenForwardingContract(t *testing.T) {
	contract := contracts.TokenForwarding(addrA.Hex())
	assert.NotNil(t, contract)
	assert.Contains(t, string(contract), addrA.Hex())
}

func TestCustomTokenForwardingContract(t *testing.T) {
	contract := contracts.CustomTokenForwarding(addrA.Hex(), "UtilityCoin", "utilityCoin")
	assert.NotNil(t, contract)
	assert.Contains(t, string(contract), addrA.Hex())
}
