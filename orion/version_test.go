package orion_test

import (
	"fmt"
	"testing"

	"github.com/marrbor/go-fiware-api/orion"
	"github.com/stretchr/testify/assert"
)

func TestAccessor_GetVersion(t *testing.T) {
	url := fmt.Sprintf("http://%s:%d", "localhost", 1026)
	a := orion.NewAccessor(url)
	v, err := a.GetVersion()
	assert.NoError(t, err)
	t.Logf("got version:%+v", v)
}
