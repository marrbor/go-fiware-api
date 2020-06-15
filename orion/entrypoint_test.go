package orion_test

import (
	"fmt"
	"testing"

	"github.com/marrbor/go-fiware-api/orion"
	"github.com/stretchr/testify/assert"
)

func TestAccessor_GetEntryPoints(t *testing.T) {
	url := fmt.Sprintf("http://%s:%d", "localhost", 1026)
	a := orion.NewAccessor(url)
	assert.EqualValues(t, url, a.BaseUrl)
	t.Logf("entrypoints: %+v", a.EntryPoints)

	ep, err := a.GetEntryPoints()
	assert.NoError(t, err)
	t.Logf("got entry points:%+v", ep)
}

// test for generate request error.
func TestAccessor_GetEntryPoints2(t *testing.T) {
	url := fmt.Sprintf("https://%s:%d", "localhost", 1026)
	a := orion.NewAccessor(url)
	assert.EqualValues(t, a.BaseUrl, url)

	// server not running.
	_, err := a.GetEntryPoints()
	assert.EqualError(t, err, "Get https://localhost:1026/v2: dial tcp [::1]:1026: connect: connection refused")
}
