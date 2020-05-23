package orion_test

import (
	"fmt"
	"testing"

	"github.com/marrbor/go-orion-api/orionapi"
	"github.com/stretchr/testify/assert"
)

func TestAccessor_GetVersion(t *testing.T) {
	err := orionapi.StartTestServer(t)
	assert.NoError(t, err)

	url := fmt.Sprintf("http://%s:%d", orionapi.Host, orionapi.Port)
	a := orionapi.NewAccessor(url)
	v, err := a.GetVersion()
	assert.NoError(t, err)
	t.Logf("got version:%+v", v)
	err = orionapi.StopTestServer(t)
	assert.NoError(t, err)
}
