package device_test

import (
	"testing"

	"github.com/marrbor/go-fiware-api/datamodel"
	common "github.com/marrbor/go-fiware-api/datamodel/device"
	"github.com/marrbor/go-fiware-api/datamodel/device/device"
	"github.com/stretchr/testify/assert"
)

func TestNewDevice(t *testing.T) {
	_, err := device.NewDevice("test", "", common.Categories.Misc, common.ControlledProperties.Location)
	assert.EqualError(t, err, common.EmptyCategoryNameError.Error())

	_, err = device.NewDevice("", "category", common.Categories.Misc, common.ControlledProperties.Location)
	assert.EqualError(t, err, datamodel.TooShortBaseNameError.Error())

	id := ""
	for i := 0; i < datamodel.MaxLengthEntityIdentifier; i++ {
		id += "a"
	}
	_, err = device.NewDevice(id, "category", common.Categories.Misc, common.ControlledProperties.Location)
	assert.EqualError(t, err, datamodel.TooLongIDLengthError.Error())

	dm, err := device.NewDevice("test", "test", common.Categories.Misc, common.ControlledProperties.AirPollution)
	assert.NoError(t, err)
	assert.EqualValues(t, datamodel.GenNgsiLdID("test", device.Type), dm.ID)

}
