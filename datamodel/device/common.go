// common schema for device and deviceModel.
//  https://github.com/FIWARE/data-models/blob/master/specs/Device/device-schema.json
package device

import (
	"fmt"
)

// enums
type (
	CategoryEnum              struct{ value string }
	ControlledPropertyEnum    struct{ value string }
	SupportedProtocolEnum     struct{ value string }
	ClassEnum                 struct{ value string }
	FunctionEnum              struct{ value string }
	EnergyLimitationClassEnum struct{ value string }
)

// SetName sets given name for this instance when this instance have no name.
func (c *CategoryEnum) SetName(name string) error {
	if 0 < len(c.value) {
		return OverWriteCategoryNameError
	}
	if len(name) <= 0 {
		return EmptyCategoryNameError
	}
	c.value = name
	return nil
}

var (
	// EmptyCategoryNameError is an error that should be returned when called NewDeviceModel with empty typeName and Categories.Misc category.
	EmptyCategoryNameError = fmt.Errorf("empty category name")

	// OverWriteCategoryNameError is an error that should be returned when an attempt to overwrite the category name.
	OverWriteCategoryNameError = fmt.Errorf("over write category name ")

	// Categories holds allowed value for "category" attribute.
	Categories = struct {
		// sensor : A device that detects and responds to events or changes in the physical environment such as light, motion, or temperature changes. https://w3id.org/saref#Sensor.
		Sensor CategoryEnum
		// actuator : A device responsible for moving or controlling a mechanism or system. https://w3id.org/saref#Actuator.
		Actuator CategoryEnum
		// meter : A device built to accurately detect and display a quantity in a form readable by a human being. Partially defined by SAREF.
		Meter CategoryEnum
		// HVAC : Heating, Ventilation and Air Conditioning (HVAC) device that provides indoor environmental comfort. https://w3id.org/saref#HVAC.
		HVAC CategoryEnum
		// network : A device used to connect other devices in a network, such as hub, switch or router in a LAN or Sensor network. (https://w3id.org/saref#Network.
		Network CategoryEnum
		// multimedia : A device designed to display, store, record or play multimedia content such as audio, images, animation, video. https://w3id.org/saref#Multimedia
		Multimedia CategoryEnum
		// implement: A device used or needed in a given activity; tool, instrument, utensil, etc. https://github.com/ADAPT/ADAPT/blob/develop/source/ADAPT/Equipment/ImplementConfiguration.cs
		Implement CategoryEnum
		// irrSystem: A mobile or fixed irrigation system such as a center pivot, linear, traveling gun, solid set, etc. https://github.com/ADAPT/ADAPT/blob/develop/source/ADAPT/Equipment/IrrSystemConfiguration.cs
		IrrSystem CategoryEnum
		// irrSection: A section of an IrrSystem. Different enough from a regular section. https://github.com/ADAPT/ADAPT/blob/develop/source/ADAPT/Equipment/IrrSectionConfiguration.cs
		IrrSection CategoryEnum
		// endgun: A device attached to an irrigation system that projects water beyond it https://github.com/ADAPT/ADAPT/blob/develop/source/ADAPT/Equipment/EndgunConfiguration.cs
		Endgun CategoryEnum
		// misc: any other meaningful to the application
		Misc CategoryEnum
	}{
		Sensor:     CategoryEnum{"sensor"},
		Actuator:   CategoryEnum{"actuator"},
		Meter:      CategoryEnum{"meter"},
		HVAC:       CategoryEnum{"HVAC"},
		Network:    CategoryEnum{"network"},
		Multimedia: CategoryEnum{"multimedia"},
		Implement:  CategoryEnum{"implement"},
		IrrSystem:  CategoryEnum{"irrSystem"},
		IrrSection: CategoryEnum{"irrSection"},
		Endgun:     CategoryEnum{"endgun"},
		Misc:       CategoryEnum{""},
	}

	// ControlledProperties holds allowed value for "controlledProperty" attribute.
	// some of this values are defined as instances of the class Property in SAREF(Smart Appliances REFerence ontology)
	ControlledProperties = struct {
		Temperature            ControlledPropertyEnum
		Humidity               ControlledPropertyEnum
		Light                  ControlledPropertyEnum
		Motion                 ControlledPropertyEnum
		FillingLevel           ControlledPropertyEnum
		Occupancy              ControlledPropertyEnum
		Power                  ControlledPropertyEnum
		Pressure               ControlledPropertyEnum
		Smoke                  ControlledPropertyEnum
		Energy                 ControlledPropertyEnum
		AirPollution           ControlledPropertyEnum
		NoiseLevel             ControlledPropertyEnum
		WeatherConditions      ControlledPropertyEnum
		Precipitation          ControlledPropertyEnum
		WindSpeed              ControlledPropertyEnum
		WindDirection          ControlledPropertyEnum
		AtmosphericPressure    ControlledPropertyEnum
		SolarRadiation         ControlledPropertyEnum
		Depth                  ControlledPropertyEnum
		PH                     ControlledPropertyEnum
		Conductivity           ControlledPropertyEnum
		Conductance            ControlledPropertyEnum
		Tss                    ControlledPropertyEnum
		Tds                    ControlledPropertyEnum // Total Dissolved Solids
		Turbidity              ControlledPropertyEnum
		Salinity               ControlledPropertyEnum
		Orp                    ControlledPropertyEnum // oxygen reduction potential
		Cdom                   ControlledPropertyEnum // Colored Dissolved Organic Matter
		WaterPollution         ControlledPropertyEnum
		Location               ControlledPropertyEnum
		Speed                  ControlledPropertyEnum
		Heading                ControlledPropertyEnum
		Weight                 ControlledPropertyEnum
		WaterConsumption       ControlledPropertyEnum
		GasComsumption         ControlledPropertyEnum
		ElectricityConsumption ControlledPropertyEnum
		EatingActivity         ControlledPropertyEnum
		Milking                ControlledPropertyEnum
		MovementActivity       ControlledPropertyEnum
		SoilMoisture           ControlledPropertyEnum
	}{
		Temperature:            ControlledPropertyEnum{"temperature"},
		Humidity:               ControlledPropertyEnum{"humidity"},
		Light:                  ControlledPropertyEnum{"light"},
		Motion:                 ControlledPropertyEnum{"motion"},
		FillingLevel:           ControlledPropertyEnum{"fillingLevel"},
		Occupancy:              ControlledPropertyEnum{"occupancy"},
		Power:                  ControlledPropertyEnum{"power"},
		Pressure:               ControlledPropertyEnum{"pressure"},
		Smoke:                  ControlledPropertyEnum{"smoke"},
		Energy:                 ControlledPropertyEnum{"energy"},
		AirPollution:           ControlledPropertyEnum{"airPollution"},
		NoiseLevel:             ControlledPropertyEnum{"noiseLevel"},
		WeatherConditions:      ControlledPropertyEnum{"weatherConditions"},
		Precipitation:          ControlledPropertyEnum{"precipitation"},
		WindSpeed:              ControlledPropertyEnum{"windSpeed"},
		WindDirection:          ControlledPropertyEnum{"windDirection"},
		AtmosphericPressure:    ControlledPropertyEnum{"atmosphericPressure"},
		SolarRadiation:         ControlledPropertyEnum{"solarRadiation"},
		Depth:                  ControlledPropertyEnum{"depth"},
		PH:                     ControlledPropertyEnum{"pH"},
		Conductivity:           ControlledPropertyEnum{"conductivity"},
		Conductance:            ControlledPropertyEnum{"conductance"},
		Tss:                    ControlledPropertyEnum{"tss"},
		Tds:                    ControlledPropertyEnum{"tds"},
		Turbidity:              ControlledPropertyEnum{"turbidity"},
		Salinity:               ControlledPropertyEnum{"salinity"},
		Orp:                    ControlledPropertyEnum{"orp"},
		Cdom:                   ControlledPropertyEnum{"cdom"},
		WaterPollution:         ControlledPropertyEnum{"waterPollution"},
		Location:               ControlledPropertyEnum{"location"},
		Speed:                  ControlledPropertyEnum{"speed"},
		Heading:                ControlledPropertyEnum{"heading"},
		Weight:                 ControlledPropertyEnum{"weight"},
		WaterConsumption:       ControlledPropertyEnum{"waterConsumption"},
		GasComsumption:         ControlledPropertyEnum{"gasComsumption"},
		ElectricityConsumption: ControlledPropertyEnum{"electricityConsumption"},
		EatingActivity:         ControlledPropertyEnum{"eatingActivity"},
		Milking:                ControlledPropertyEnum{"milking"},
		MovementActivity:       ControlledPropertyEnum{"movementActivity"},
		SoilMoisture:           ControlledPropertyEnum{"soilMoisture"},
	}

	// SupportedProtocols holds allowed value for "supportedUnits" attribute.
	SupportedProtocols = struct {
		Ul20      SupportedProtocolEnum
		Mqtt      SupportedProtocolEnum
		Lwm2m     SupportedProtocolEnum
		Http      SupportedProtocolEnum
		Websocket SupportedProtocolEnum
		Onem2m    SupportedProtocolEnum
		Sigfox    SupportedProtocolEnum
		Lora      SupportedProtocolEnum
		NbIot     SupportedProtocolEnum
		EcGsmIot  SupportedProtocolEnum
		LteM      SupportedProtocolEnum
		CatM      SupportedProtocolEnum
		ThreeG    SupportedProtocolEnum
		Gprs      SupportedProtocolEnum
		Coap      SupportedProtocolEnum
	}{
		Ul20:      SupportedProtocolEnum{"ul20"},
		Mqtt:      SupportedProtocolEnum{"mqtt"},
		Lwm2m:     SupportedProtocolEnum{"lwm2m"},
		Http:      SupportedProtocolEnum{"http"},
		Websocket: SupportedProtocolEnum{"websocket"},
		Onem2m:    SupportedProtocolEnum{"onem2m"},
		Sigfox:    SupportedProtocolEnum{"sigfox"},
		Lora:      SupportedProtocolEnum{"lora"},
		NbIot:     SupportedProtocolEnum{"nb-iot"},
		EcGsmIot:  SupportedProtocolEnum{"ec-gsm-iot"},
		LteM:      SupportedProtocolEnum{"lte-m"},
		CatM:      SupportedProtocolEnum{"cat-m"},
		ThreeG:    SupportedProtocolEnum{"3g"},
		Gprs:      SupportedProtocolEnum{"gprs"},
		Coap:      SupportedProtocolEnum{"coap"},
	}

	// DeviceClasses holds allowed value for "deviceClass" attribute.
	DeviceClasses = struct {
		C0 ClassEnum
		C1 ClassEnum
		C2 ClassEnum
	}{
		C0: ClassEnum{"C0"},
		C1: ClassEnum{"C1"},
		C2: ClassEnum{"C2"},
	}

	// Functions holds allowed value for "functions" attribute.
	Functions = struct {
		LevelControl      FunctionEnum
		Sensing           FunctionEnum
		OnOff             FunctionEnum
		OpenClose         FunctionEnum
		Metering          FunctionEnum
		EventNotification FunctionEnum
	}{
		LevelControl:      FunctionEnum{"levelControl"},
		Sensing:           FunctionEnum{"sensing"},
		OnOff:             FunctionEnum{"onOff"},
		OpenClose:         FunctionEnum{"openClose"},
		Metering:          FunctionEnum{"metering"},
		EventNotification: FunctionEnum{"eventNotification"},
	}

	// EnergyLimitationClasses holds allowed value for "energyLimitationClasses" attribute.
	EnergyLimitationClasses = struct {
		E0 EnergyLimitationClassEnum
		E1 EnergyLimitationClassEnum
		E2 EnergyLimitationClassEnum
		E9 EnergyLimitationClassEnum
	}{
		E0: EnergyLimitationClassEnum{"E0"},
		E1: EnergyLimitationClassEnum{"E1"},
		E2: EnergyLimitationClassEnum{"E2"},
		E9: EnergyLimitationClassEnum{"E9"},
	}
)

// DeviceCommon based on #Device-Common
type DeviceCommon struct {
	Category           []CategoryEnum           `json:"category"`
	ControlledProperty []ControlledPropertyEnum `json:"controlledProperty"`
	SupportedProtocol  []SupportedProtocolEnum  `json:"supportedProtocol"`
}
