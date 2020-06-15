/*
 * device model data model
 * https://github.com/FIWARE/data-models/blob/master/specs/Device/DeviceModel/schema.json
 * https://fiware-apis.readthedocs.io/en/stable/Device/DeviceModel/doc/spec/index.html
 */
package devicemodel

import (
	"github.com/marrbor/go-fiware-api/datamodel"
	"github.com/marrbor/go-fiware-api/datamodel/device"
)

const (
	Type = "DeviceModel"
)

// DeviceModel
type DeviceModel struct {
	datamodel.Core                                         // Core.Type must be "DeviceModel"
	Category              []device.CategoryEnum            `json:"category" mandatory:"true"`
	DeviceClass           device.ClassEnum                 `json:"deviceClass" mandatory:"false"`
	ControlledProperty    []device.ControlledPropertyEnum  `json:"controlledProperty" mandatory:"true"`
	Function              []device.FunctionEnum            `json:"function" mandatory:"false"`
	SupportedUnits        []string                         `json:"supportedUnits" mandatory:"false"`
	EnergyLimitationClass device.EnergyLimitationClassEnum `json:"energyLimitationClass" mandatory:"false"`
	Documentation         string                           `json:"documentation" mandatory:"false"` // format: uri
	BrandName             string                           `json:"brandName" mandatory:"true"`
	ModelName             string                           `json:"modelName" mandatory:"true"`
	ManufacturerName      string                           `json:"manufacturerName" mandatory:"true"`
}

// NewDeviceModel returns new DeviceModel instance with mandatory attributes only.
// typeName is mandatory when give 'Categories.Misc' as category parameter.
func NewDeviceModel(id, categoryName, brandName, modelName, manufactureName string, category device.CategoryEnum, cp device.ControlledPropertyEnum) (*DeviceModel, error) {
	if category == device.Categories.Misc {
		if err := category.SetName(categoryName); err != nil {
			return nil, err
		}
	}

	c, err := datamodel.NewCore(id, Type)
	if err != nil {
		return nil, err
	}

	return &DeviceModel{
		Core:               *c,
		Category:           []device.CategoryEnum{category},
		ControlledProperty: []device.ControlledPropertyEnum{cp},
		BrandName:          brandName,
		ModelName:          modelName,
		ManufacturerName:   manufactureName,
	}, nil
}
