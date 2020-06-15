/*
 * device data model
 * github.com/fiware/data-models/specs/Device/device-schema.json #Device-Common
 * https://fiware-datamodels.readthedocs.io/en/stable/Device/Device/doc/spec/index.html
 */
package device

import (
	"net/url"
	"time"

	"github.com/marrbor/go-fiware-api/datamodel"
	"github.com/marrbor/go-fiware-api/datamodel/device"
	"github.com/marrbor/go-fiware-api/datamodel/device/devicemodel"
)

const (
	Type = "Device"
)

// Device
type Device struct {
	datamodel.Core                                        // Core.Type must be "Device"
	Category              []device.CategoryEnum           `json:"category" mandatory:"true"`
	ControlledProperty    []device.ControlledPropertyEnum `json:"controlledProperty" mandatory:"true"`
	MNC                   string                          `json:"mnc" mandatory:"false"`             // mobile network code
	ControlledAsset       []url.URL                       `json:"controlledAsset" mandatory:"false"` // "https://smart-data-models.github.io/data-models/common-schema.json#/definitions/EntityIdentifierType"
	MacAddress            []string                        `json:"macAddress" mandatory:"false"`      // https://smart-data-models.github.io/data-models/specs/Device/device-schema.json#/definitions/MacAddressType
	IpAddress             []string                        `json:"ipAddress" mandatory:"false"`       // one of ipv4/ipv6
	Configuration         interface{}                     `json:"configuration" mandatory:"false"`
	Location              datamodel.Location              `json:"location" mandatory:"false"`
	Name                  string                          `json:"name" mandatory:"false"`
	DateInstalled         time.Time                       `json:"dateInstalled" mandatory:"false"`
	DateFirstUsed         time.Time                       `json:"dateFirstUsed" mandatory:"false"`
	DateManufactured      time.Time                       `json:"dateManufactured" mandatory:"false"`
	HardwareVersion       string                          `json:"hardwareVersion" mandatory:"false"`
	SoftwareVersion       string                          `json:"softwareVersion" mandatory:"false"`
	FirmwareVersion       string                          `json:"firmwareVersion" mandatory:"false"`
	OsVersion             string                          `json:"osVersion" mandatory:"false"`
	DateLastCalibration   time.Time                       `json:"dateLastCalibration" mandatory:"false"`
	SerialNumber          string                          `json:"serialNumber" mandatory:"false"`
	Provider              string                          `json:"provider" mandatory:"false"`
	RefDeviceModel        devicemodel.DeviceModel         `json:"refDeviceModel" mandatory:"false"`
	BatteryLevel          float64                         `json:"batteryLevel" mandatory:"false"` // -1 or 0.0 to 1.0. -1.0 means failed to get level.
	RSSI                  float64                         `json:"rssi" mandatory:"false"`         // 0.0(weak) to 1.0(max) -1.0 means failed to get.
	DeviceState           string                          `json:"deviceState" mandatory:"false"`
	DateLastValueReported time.Time                       `json:"dateLastValueReported" mandatory:"false"`
	Value                 string                          `json:"value" mandatory:"false"`
}

// NewDevice returns new Device instance with mandatory attributes only.
// typeName is mandatory when give 'Categories.Misc' as category parameter.
func NewDevice(id, categoryName string, category device.CategoryEnum, cp device.ControlledPropertyEnum) (*Device, error) {
	if category == device.Categories.Misc {
		if err := category.SetName(categoryName); err != nil {
			return nil, err
		}
	}

	c, err := datamodel.NewCore(id, Type)
	if err != nil {
		return nil, err
	}

	return &Device{
		Core:               *c,
		Category:           []device.CategoryEnum{category},
		ControlledProperty: []device.ControlledPropertyEnum{cp},
	}, nil
}
