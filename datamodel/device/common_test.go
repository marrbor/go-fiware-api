package device_test

import (
	"testing"

	"github.com/marrbor/go-fiware-api/datamodel/device"
	"github.com/stretchr/testify/assert"
)

func TestCategoryEnum_SetName(t *testing.T) {
	ca := device.Categories.Actuator
	err := ca.SetName("xxx")
	assert.EqualError(t, err, device.OverWriteCategoryNameError.Error())

	cm := device.Categories.Misc
	err = cm.SetName("")
	assert.EqualError(t, err, device.EmptyCategoryNameError.Error())

	err = cm.SetName("xxx")
	assert.NoError(t, err)
}
