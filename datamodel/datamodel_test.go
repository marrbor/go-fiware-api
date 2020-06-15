package datamodel_test

import (
	"fmt"
	"testing"

	"github.com/marrbor/go-fiware-api/datamodel"
	"github.com/stretchr/testify/assert"
)

func TestValidateIdentifier(t *testing.T) {
	assert.True(t, datamodel.ValidateIdentifier("urn:ngsi-ld:Abc:1234"))
	assert.False(t, datamodel.ValidateIdentifier(""))
}

func TestCore_Stamp(t *testing.T) {
	c, err := datamodel.NewCore("Test1", "Test")
	assert.NoError(t, err)
	assert.NotNil(t, c)
	assert.True(t, c.DateCreated.IsZero())
	assert.True(t, c.DateModified.IsZero())
	c.Stamp()
	assert.False(t, c.DateCreated.IsZero())
	assert.False(t, c.DateModified.IsZero())
	assert.EqualValues(t, c.DateCreated, c.DateModified)
	c.Stamp()
	assert.False(t, c.DateCreated.IsZero())
	assert.False(t, c.DateModified.IsZero())
	assert.NotEqual(t, c.DateCreated, c.DateModified)
}

func TestNewCore(t *testing.T) {
	c, err := datamodel.NewCore("", "")
	assert.Nil(t, c)
	assert.EqualError(t, err, datamodel.TooShortBaseNameError.Error())
}

func TestNewCore2(t *testing.T) {
	c, err := datamodel.NewCore("abc", "")
	assert.Nil(t, c)
	assert.EqualError(t, err, datamodel.TooShortTypeNameError.Error())
}

func TestNewCore3(t *testing.T) {
	c, err := datamodel.NewCore("abc", "def")
	assert.Nil(t, c)
	assert.EqualError(t, err, datamodel.InvalidTypeNameError.Error())
}

func TestNewCore4(t *testing.T) {
	// bn too long
	tn := "Def"
	bn := ""
	for (len(tn) + len(bn)) <= datamodel.AllowedIdentifierLength {
		bn += "a"
	}
	c, err := datamodel.NewCore(bn, tn)
	assert.Nil(t, c)
	assert.EqualError(t, err, datamodel.TooLongIDLengthError.Error())
}

func TestNewCore5(t *testing.T) {
	// tn too long
	bn := "abc"
	tn := "Aaa"
	for (len(tn) + len(bn)) <= datamodel.AllowedIdentifierLength {
		tn += "a"
	}
	c, err := datamodel.NewCore(bn, tn)
	assert.Nil(t, c)
	assert.EqualError(t, err, datamodel.TooLongIDLengthError.Error())
}

func TestNewCore6(t *testing.T) {
	// both too long
	bn := "a"
	tn := "A"
	for (len(bn) + len(tn)) <= datamodel.AllowedIdentifierLength {
		bn += "a"
		tn += "a"
	}
	c, err := datamodel.NewCore(bn, tn)
	assert.Nil(t, c)
	assert.EqualError(t, err, datamodel.TooLongIDLengthError.Error())
}

func TestNewCore7(t *testing.T) {
	bn := "a"
	tn := "A"
	for len(bn) < datamodel.AllowedIdentifierLength/2-1 {
		bn += "a"
	}
	for len(tn) < datamodel.AllowedIdentifierLength/2-1 {
		tn += "a"
	}
	c, err := datamodel.NewCore(bn, tn)
	assert.NotNil(t, c)
	assert.NoError(t, err)
	assert.EqualValues(t, fmt.Sprintf("urn:ngsi-ld:%s:%s", tn, bn), c.ID)
	assert.EqualValues(t, c.Type, tn)
}

func TestValidateTypeName(t *testing.T) {
	tn := "abc"
	assert.False(t, datamodel.ValidateTypeName(tn))
	tn = "123"
	assert.False(t, datamodel.ValidateTypeName(tn))
	tn = "Abc"
	assert.True(t, datamodel.ValidateTypeName(tn))
}

func TestValidateDateYearLess(t *testing.T) {
	// valid 31th
	for _, m := range []int{1, 3, 5, 7, 8, 10, 12} {
		d := fmt.Sprintf("--%02d-31", m)
		assert.True(t, datamodel.ValidateDateYearLess(d))
	}
	// valid 30th
	for m := 1; m <= 12; m++ {
		if m == 2 {
			continue
		}
		assert.True(t, datamodel.ValidateDateYearLess(fmt.Sprintf("--%02d-30", m)))
	}

	// valid 1st to 29th
	for m := 1; m <= 12; m++ {
		for d := 1; d <= 29; d++ {
			assert.True(t, datamodel.ValidateDateYearLess(fmt.Sprintf("--%02d-%02d", m, d)))
		}
	}

	// invalid 31th
	for _, m := range []int{2, 4, 6, 9, 11} {
		assert.False(t, datamodel.ValidateDateYearLess(fmt.Sprintf("--%02d-31", m)))
	}

	// invalid 30th
	assert.False(t, datamodel.ValidateDateYearLess("--02-30"))
}
