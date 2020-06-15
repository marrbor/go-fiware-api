/*
 * Data model (Location)
 *
 * https://fiware-datamodels.readthedocs.io/en/latest/howto/index.html
 * https://github.com/fiware/data-models/common-schema.json #Location-Commons
 */
package datamodel

import (
	"fmt"
	"strconv"
)

const (
	TypePoint           = "Point"
	TypeLineString      = "LineString"
	TypePolygon         = "Polygon"
	TypeMultiPoint      = "MultiPoint"
	TypeMultiLineString = "MultiLineString"
	TypeMultiPolygon    = "MultiPolygon"
)

var MismatchTypeError = fmt.Errorf("given strings are not location type")

// Location
type Location struct {
	Location   interface{} `json:"location"` //  "$ref": "http://geojson.org/schema/Geometry.json#"
	Address    Address     `json:"address"`
	AreaServed string      `json:"areaServed"`
}

// Address
type Address struct {
	StreetAddress       string `json:"streetAddress"`
	AddressLocality     string `json:"addressLocality"`
	AddressRegion       string `json:"addressRegion"`
	AddressCountry      string `json:"addressCountry"`
	PostalCode          string `json:"postalCode"`
	PostOfficeBoxNumber string `json:"postOfficeBoxNumber"`
}

// GenLocation returns specified LocationItem instance.
func GenLocation(typeName string) (*Location, error) {
	ret := Location{
		Location:   nil,
		Address:    Address{},
		AreaServed: "",
	}
	bbox := make([]float64, 4)
	switch typeName {
	case TypePoint:
		ret.Location = Point{Type: TypePoint, Coordinates: make([]float64, 2), Bbox: bbox}
	case TypeLineString:
		ret.Location = LineString{Type: TypeLineString, Coordinates: [][]float64{make([]float64, 2), make([]float64, 2)}, Bbox: bbox}
	case TypePolygon:
		ret.Location = Polygon{Type: TypePolygon, Coordinates: [][][]float64{{make([]float64, 2), make([]float64, 2), make([]float64, 2), make([]float64, 2)}}, Bbox: bbox}
	case TypeMultiPoint:
		ret.Location = MultiPoint{Type: TypeMultiPoint, Coordinates: [][]float64{make([]float64, 2)}, Bbox: bbox}
	case TypeMultiLineString:
		ret.Location = MultiLineString{Type: TypeMultiLineString, Coordinates: [][][]float64{{make([]float64, 2), make([]float64, 2)}}, Bbox: bbox}
	case TypeMultiPolygon:
		ret.Location = MultiPolygon{Type: TypeMultiPolygon, Coordinates: [][][][]float64{{{make([]float64, 2), make([]float64, 2), make([]float64, 2), make([]float64, 2)}}}, Bbox: bbox}
	default:
		return nil, MismatchTypeError
	}
	return &ret, nil
}

////// Location interfaces ref: https://geojson.org/schema/Geometry.json

// Point is an expression of GeoJSON Point.
type Point struct {
	Type        string    `json:"type"`        // must be "Point". required.
	Coordinates []float64 `json:"coordinates"` // have to hold at least 2 items. required.
	Bbox        []float64 `json:"bbox"`        // have to hold at least 4 items. optional.
}

// LineString is an expression of GeoJSON LineString
type LineString struct {
	Type        string      `json:"type"`        // must be "LineString". required.
	Coordinates [][]float64 `json:"coordinates"` // array that have to hold at least 2 array that have to hold at leaset 2 items. required.
	Bbox        []float64   `json:"bbox"`        // have to hold at least 4 items. optional.
}

// Polygon is an expresion of GeoJSON Polygon
type Polygon struct {
	Type        string        `json:"type"`        // must be "Polygon". required.
	Coordinates [][][]float64 `json:"coordinates"` // array of array that have to hold at least 4 items that holds 2 items. required.
	Bbox        []float64     `json:"bbox"`        // have to hold at least 4 items. optional.
}

// MultiPoint is an expression of GeoJSON MultiPoint.
type MultiPoint struct {
	Type        string      `json:"type"`        // must be "MultiPoint". required.
	Coordinates [][]float64 `json:"coordinates"` // array of array that have to hold at least 2 items. required.
	Bbox        []float64   `json:"bbox"`        // have to hold at least 4 items. optional.
}

// MultiLineString is an expression of GeoJSON MultiLineString.
type MultiLineString struct {
	Type        string        `json:"type"`        // must be "MultiLineString". required.
	Coordinates [][][]float64 `json:"coordinates"` // array of array that have to hold at least 2 arrays that hold at least 2 items. required.
	Bbox        []float64     `json:"bbox"`        // have to hold at least 4 items. optional.
}

// MultiPolygon is an expression of GeoJSON MultiPolygon.
type MultiPolygon struct {
	Type        string          `json:"type"`        // must be "MultiPolygon". required.
	Coordinates [][][][]float64 `json:"coordinates"` // array of array of array that have to hold at least 4 array that have to hold at least 2 items. required.
	Bbox        []float64       `json:"bbox"`        // have to hold at least 4 items. optional.
}

// XY is an expression of the coordinate holds some [X,Y]
// This is not defined at fiware data model.
type XY struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

// String returns string expresion used by sprintf("%f").
func (xy *XY) String() string {
	x := strconv.FormatFloat(xy.X, 'f', -1, 64)
	y := strconv.FormatFloat(xy.Y, 'f', -1, 64)
	return fmt.Sprintf("%s,%s", x, y)
}

// LatLng is an expression of the coordinate holds some [latitude, longitude]
// This is not defined at fiware data model.
type LatLng struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// String returns string expresion used by strconv.Ftoa64
func (ll *LatLng) String() string {
	lat := strconv.FormatFloat(ll.Latitude, 'f', -1, 64)
	lng := strconv.FormatFloat(ll.Longitude, 'f', -1, 64)
	return fmt.Sprintf("%s,%s", lat, lng)
}
