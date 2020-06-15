package orion_test

import (
	"testing"

	"github.com/marrbor/go-fiware-api/orion"
	"github.com/stretchr/testify/assert"
)

func TestNewAccessor(t *testing.T) {
	a := orion.NewAccessor("localhost:9999")
	assert.NotNil(t, a)
	assert.EqualValues(t, a.BaseUrl, "localhost:9999")

	a = orion.NewAccessor("localhost:8888")
	assert.NotNil(t, a)
	assert.EqualValues(t, a.BaseUrl, "localhost:8888")
}
