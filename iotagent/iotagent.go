// IoT Agent API common definitions.
// https://github.com/telefonicaid/iotagent-node-lib/blob/master/doc/api.md
// https://swagger.lab.fiware.org/?url=https://raw.githubusercontent.com/Fiware/specifications/master/OpenAPI/iot.IoTagent-node-lib/IoTagent-node-lib-openapi.json#/
package iotagent

import (
	"fmt"
	"net/http"
)

const (
	// https://fiware-iotagent-json.readthedocs.io/en/latest/stepbystep/index.html#provisioning-the-device
	DefaultIoTAPIKey = "1234"

	// Type of NameTypePair
	TypeCommand    = "command"
	TypeText       = "Text"
	TypeRelate     = "Relationship"
	TypeInteger    = "Integer"
	TypeFloat      = "Float"
	TypeBoolean    = "Boolean"
	TypePercentage = "percentage"

	// unitCode
	// https://www.unece.org/fileadmin/DAM/cefact/recommendations/rec20/rec20_rev3_Annex3e.pdf
	UnitDecibel     = "2N"
	UnitKilobyte    = "2P"
	UnitMeterPerMin = "2X" // m/min
	UnitMilliVolt   = "2Z" // mV
	UnitManMonth    = "3C"
	UnitMicroMetre  = "4H"  // um
	UnitMilliAmpere = "4K"  // mA
	UnitMegaByte    = "4L"  // MB 10^6
	UnitAmpereHour  = "AMH" // A-h
	UnitAmpere      = "AMP" // A
	UnitYear        = "ANN" // a
	UnitMicroAmpere = "B84" // uA
	UnitMicroSecond = "B98" // us
	UnitGigaByte    = "E34" // GB
	UnitTeraByte    = "E35" // TB
	UnitPetaByte    = "E36" // PB
	UnitPixel       = "E37"
	UnitMegaPixel   = "E38"
	UnitDpi         = "E39" // dpi
	UnitPercentage  = "P1"
	UnitFahrenheit  = "FAH" // F
	UnitHour        = "HUR" // h
	UnitKiloGram    = "KGM" // kg
	UnitKiloHertz   = "KHZ" // kHz
	UnitKiloMeter   = "KMT" // km
	UnitKiloWatt    = "KWT" // kW
	UnitMegaWatt    = "MAW" // MW
	UnitMinute      = "MIN" // min
	UnitMillimeter  = "MMT" // mm
	UnitMeter       = "MTR" // m
	UnitVolt        = "VLT" // V
	UnitWattHour    = "WHR" // W-h
	UnitWatt        = "WTT" // W
)

type (
	// NameTypePair holds name/type pair
	NameTypePair struct {
		Name string `json:"name" bson:"name"`
		Type string `json:"type" bson:"type"`
	}

	// TypeValuePair holds type/value pair
	TypeValuePair struct {
		Type  string `json:"type" bson:"type"`
		Value string `json:"value" bson:"value"`
	}

	// NameTypeValues holds name/type/value unit.
	NameTypeValues struct {
		Name  string `json:"name" bson:"name"`
		Type  string `json:"type" bson:"type"`
		Value string `json:"value" bson:"value"`
	}

	// UnitMetadata holds unit of attribute.
	UnitMetadata struct {
		UnitCode TypeValuePair `json:"unitCode" bson:"unitCode"`
	}
)

// UnitCodes
var (
	UcPercentage      = TypeValuePair{Type: "Text", Value: UnitPercentage}
	UcUnitDecibel     = TypeValuePair{Type: "Text", Value: UnitDecibel}
	UcUnitKilobyte    = TypeValuePair{Type: "Text", Value: UnitKilobyte}
	UcUnitMeterPerMin = TypeValuePair{Type: "Text", Value: UnitMeterPerMin}
	UcUnitMilliVolt   = TypeValuePair{Type: "Text", Value: UnitMilliVolt}
	UcUnitManMonth    = TypeValuePair{Type: "Text", Value: UnitManMonth}
	UcUnitMicroMetre  = TypeValuePair{Type: "Text", Value: UnitMicroMetre}
	UcUnitMilliAmpere = TypeValuePair{Type: "Text", Value: UnitMilliAmpere}
	UcUnitMegaByte    = TypeValuePair{Type: "Text", Value: UnitMegaByte}
	UcUnitAmpereHour  = TypeValuePair{Type: "Text", Value: UnitAmpereHour}
	UcUnitAmpere      = TypeValuePair{Type: "Text", Value: UnitAmpere}
	UcUnitYear        = TypeValuePair{Type: "Text", Value: UnitYear}
	UcUnitMicroAmpere = TypeValuePair{Type: "Text", Value: UnitMicroAmpere}
	UcUnitMicroSecond = TypeValuePair{Type: "Text", Value: UnitMicroSecond}
	UcUnitGigaByte    = TypeValuePair{Type: "Text", Value: UnitGigaByte}
	UcUnitTeraByte    = TypeValuePair{Type: "Text", Value: UnitTeraByte}
	UcUnitPetaByte    = TypeValuePair{Type: "Text", Value: UnitPetaByte}
	UcUnitPixel       = TypeValuePair{Type: "Text", Value: UnitPixel}
	UcUnitMegaPixel   = TypeValuePair{Type: "Text", Value: UnitMegaPixel}
	UcUnitDpi         = TypeValuePair{Type: "Text", Value: UnitDpi}
	UcUnitPercentage  = TypeValuePair{Type: "Text", Value: UnitPercentage}
	UcUnitFahrenheit  = TypeValuePair{Type: "Text", Value: UnitFahrenheit}
	UcUnitHour        = TypeValuePair{Type: "Text", Value: UnitHour}
	UcUnitKiloGram    = TypeValuePair{Type: "Text", Value: UnitKiloGram}
	UcUnitKiloHertz   = TypeValuePair{Type: "Text", Value: UnitKiloHertz}
	UcUnitKiloMeter   = TypeValuePair{Type: "Text", Value: UnitKiloMeter}
	UcUnitKiloWatt    = TypeValuePair{Type: "Text", Value: UnitKiloWatt}
	UcUnitMegaWatt    = TypeValuePair{Type: "Text", Value: UnitMegaWatt}
	UcUnitMinute      = TypeValuePair{Type: "Text", Value: UnitMinute}
	UcUnitMillimeter  = TypeValuePair{Type: "Text", Value: UnitMillimeter}
	UcUnitMeter       = TypeValuePair{Type: "Text", Value: UnitMeter}
	UcUnitVolt        = TypeValuePair{Type: "Text", Value: UnitVolt}
	UcUnitWattHour    = TypeValuePair{Type: "Text", Value: UnitWattHour}
	UcUnitWatt        = TypeValuePair{Type: "Text", Value: UnitWatt}
)

// NewReference returns new reference NameTypeValues instance.
func NewReference(category, id string) NameTypeValues {
	return NameTypeValues{Name: fmt.Sprintf("ref%s", category), Type: TypeRelate, Value: id}
}

// NewText returns new text type NameTypeValues instance.
func NewText(name, value string) NameTypeValues {
	return NameTypeValues{Name: name, Type: TypeText, Value: value}
}


// AddIoTHeader adds common header to given request and return it.
func AddIoTHeader(service, path string, req *http.Request) *http.Request {
	h := req.Header
	h.Add("Fiware-Service", service)
	h.Add("Fiware-ServicePath", path)
	h.Add("Cache-Control", "no-cache")
	if req.Method == http.MethodPost || req.Method == http.MethodPut {
		h.Add("Content-Type", "application/json")
	}
	if req.Method == http.MethodGet {
		h.Add("Accept", "application/json")
	}
	return req
}
