package datamodel_test

import (
	"testing"

	"github.com/marrbor/go-fiware-api/datamodel"
	"github.com/stretchr/testify/assert"
)

func TestValidatePhysicalObjectImageString(t *testing.T) {
	// invalid
	err := datamodel.ValidatePhysicalObjectImageString("abc")
	assert.Error(t, err)
	err = datamodel.ValidatePhysicalObjectImageString("")
	assert.Error(t, err)
	err = datamodel.ValidatePhysicalObjectImageString("google.com")
	assert.Error(t, err)

	// valid
	err = datamodel.ValidatePhysicalObjectImageString(string(datamodel.GenNgsiLdID("abc", "Abc")))
	assert.NoError(t, err)
	err = datamodel.ValidatePhysicalObjectImageString("/path/to")
	assert.NoError(t, err)
	err = datamodel.ValidatePhysicalObjectImageString("http://google.com")
	assert.NoError(t, err)
	err = datamodel.ValidatePhysicalObjectImageString("https://google.com")
	assert.NoError(t, err)
	err = datamodel.ValidatePhysicalObjectImageString("http://google.com/path/to")
	assert.NoError(t, err)
	err = datamodel.ValidatePhysicalObjectImageString("https://google.com/path/to/")
	assert.NoError(t, err)
}
