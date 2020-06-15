package devicemodel_test

import (
	"github.com/marrbor/go-fiware-api/datamodel"
	"github.com/marrbor/go-fiware-api/datamodel/device"
	"github.com/marrbor/go-fiware-api/datamodel/device/devicemodel"
	"github.com/stretchr/testify/assert"

	"testing"
)

func TestNewDeviceModel(t *testing.T) {
	_, err := devicemodel.NewDeviceModel("test", "", "", "", "",
		device.Categories.Misc, device.ControlledProperties.Location)
	assert.EqualError(t, err, device.EmptyCategoryNameError.Error())

	_, err = devicemodel.NewDeviceModel("", "category", "", "", "",
		device.Categories.Misc, device.ControlledProperties.Location)
	assert.EqualError(t, err, datamodel.TooShortBaseNameError.Error())

	id := ""
	for i := 0; i < datamodel.MaxLengthEntityIdentifier; i++ {
		id += "a"
	}
	_, err = devicemodel.NewDeviceModel(id, "category", "", "", "",
		device.Categories.Misc, device.ControlledProperties.Location)
	assert.EqualError(t, err, datamodel.TooLongIDLengthError.Error())

	dm, err := devicemodel.NewDeviceModel("test", "test", "brand", "model", "manufacture",
		device.Categories.Misc, device.ControlledProperties.AirPollution)
	assert.NoError(t, err)
	assert.EqualValues(t, datamodel.GenNgsiLdID("test", devicemodel.Type), dm.ID)
}
