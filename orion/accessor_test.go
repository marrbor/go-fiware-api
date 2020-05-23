package orion_test

import (
	"testing"

	"github.com/marrbor/go-orion-api/orionapi"
	"github.com/stretchr/testify/assert"
)

func TestNewAccessor(t *testing.T) {
	a := orionapi.NewAccessor("localhost:9999")
	assert.NotNil(t, a)
	assert.EqualValues(t, a.BaseUrl, "localhost:9999")

	a = orionapi.NewAccessor("localhost:8888")
	assert.NotNil(t, a)
	assert.EqualValues(t, a.BaseUrl, "localhost:8888")
}

func TestIsValidService(t *testing.T) {
	// valid
	assert.True(t, orionapi.IsValidService("_"))
	assert.True(t, orionapi.IsValidService("a"))
	assert.True(t, orionapi.IsValidService("b"))
	assert.True(t, orionapi.IsValidService("y"))
	assert.True(t, orionapi.IsValidService("z"))

	assert.True(t, orionapi.IsValidService("_a"))
	assert.True(t, orionapi.IsValidService("ab"))
	assert.True(t, orionapi.IsValidService("a_b"))
	assert.True(t, orionapi.IsValidService("ab_"))

	assert.True(t, orionapi.IsValidService("yz"))
	assert.True(t, orionapi.IsValidService("y_z"))
	assert.True(t, orionapi.IsValidService("yz_"))

	assert.True(t, orionapi.IsValidService("0"))
	assert.True(t, orionapi.IsValidService("1_2"))
	assert.True(t, orionapi.IsValidService("89_"))

	assert.True(t, orionapi.IsValidService("_abcdefghijklmnopqrstuvwxyz_9012345678901234567890"))

	// invalid
	assert.False(t, orionapi.IsValidService(""))
	assert.False(t, orionapi.IsValidService("_abcdefghijklmnopqrstuvwxyz_90123456789012345678901"))
}

func TestIsValidServicePath(t *testing.T) {
	// valid
	assert.True(t, orionapi.IsValidServicePath("/"))
	assert.True(t, orionapi.IsValidServicePath("/abc"))
	assert.True(t, orionapi.IsValidServicePath("/abc/def/ghi/jkl/mno/pqr/stu/vwx/yz_/012")) // 10 depth
	assert.True(t, orionapi.IsValidServicePath("/_abcdefghijklmnopqrstuvwxyz_9012345678901234567890/_abcdefghijklmnopqrstuvwxyz_9012345678901234567890"))

	// invalid
	assert.False(t, orionapi.IsValidServicePath(""))                                                         // empty
	assert.False(t, orionapi.IsValidServicePath("//"))                                                       // empty
	assert.False(t, orionapi.IsValidServicePath("abc"))                                                      // not start with '/'
	assert.False(t, orionapi.IsValidServicePath("/ABC"))                                                     // use upper case
	assert.False(t, orionapi.IsValidServicePath("/_abcdefghijklmnopqrstuvwxyz_90123456789012345678901/abc")) // over 50 letters
	assert.False(t, orionapi.IsValidServicePath("/abc/def/ghi/jkl/mno/pqr/stu/vwx/yz_/012/345"))             // 11 depth
	assert.False(t, orionapi.IsValidServicePath("/abc/def/ghi/jkl/mno/pqr/stu/vwx/yz_/012/345/678"))         // 11 depth
}
