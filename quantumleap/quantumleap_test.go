// QuantumLeap test
package quantumleap_test

import (
	"os"
	"testing"

	"github.com/marrbor/go-fiware-api/quantumleap"
	"github.com/stretchr/testify/assert"
)

const (
	EnvQuantumLeapUrl     = "QL_URL"
	DefaultQuantumLeapUrl = "http://localhost/quantumleap"
)

func TestAccessor_GetVersion(t *testing.T) {
	url := DefaultQuantumLeapUrl
	if 0 < len(os.Getenv(EnvQuantumLeapUrl)) {
		url = os.Getenv(EnvQuantumLeapUrl)
	}
	a := quantumleap.NewAccessor(url)
	v, err := a.GetVersion()
	assert.NoError(t, err)
	assert.EqualValues(t, "0.7.5", v.Version)
}
