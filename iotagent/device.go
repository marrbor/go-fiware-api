// Device API
// https://github.com/telefonicaid/iotagent-node-lib/blob/master/doc/api.md#device-api

package iotagent

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/marrbor/gohttp"
)

const (
	DevicesUrl = "iot/devices"
)

var (
	DeviceNotFoundError = fmt.Errorf("deveice not found")
)

type (
	Device struct {
		DeviceID           string             `json:"device_id" bson:"id"`                                      // Device ID that will be used to identify the device.	ex) UO834IO
		Service            *string            `json:"service,omitempty" bson:"service"`                         // Name of the service the device belongs to (will be used in the fiware-service header).	ex) smartGondor
		ServicePath        *string            `json:"service_path,omitempty" bson:"subservice"`                 // Name of the subservice the device belongs to (used in the fiware-servicepath header). ex) /gardens
		EntityName         string             `json:"entity_name" bson:"name"`                                  // Name of the entity representing the device in the Context Broker ex) ParkLamplight12
		EntityType         string             `json:"entity_type" bson:"type"`                                  // Type of the entity in the Context Broker ex) Lamplights
		TimeZone           *string            `json:"timezone,omitempty" bson:"timezone"`                       // Time zone of the sensor if it has any ex) America/Santiago
		TimeStamp          *bool              `json:"timestamp,omitempty" bson:"timestamp"`                     // Optional flag about whether or not to addthe TimeInstant attribute to the device entity created, as well as a TimeInstant metadata to each attribute, with the current timestamp ex) true
		ApiKey             *string            `json:"apikey,omitempty" bson:"apikey"`                           // Optional Apikey key string to use instead of group apikey ex) 9n4hb1vpwbjozzmw9f0flf9c2
		EndPoint           *string            `json:"endpoint,omitempty" bson:"endpoint"`                       // Endpoint where the device is going to receive commands, if any. ex) http://theDeviceUrl:1234/commands
		Protocol           string             `json:"protocol" bson:"protocol"`                                 // Name of the device protocol, for its use with an IoT Manager. ex) IoTA-UL
		Transport          string             `json:"transport" bson:"transport"`                               // Name of the device transport protocol, for the IoT Agents with multiple transport protocols. ex) MQTT
		Attributes         *[]DeviceAttribute `json:"attributes,omitempty" bson:"attributes"`                   // attributes	active	List of active attributes of the device	ex) [ { "name": "attr_name", "type": "Text" } ]
		Lazy               *[]DeviceAttribute `json:"lazy,omitempty" bson:"lazy"`                               // List of lazy attributes of the device ex) [ { "name": "attr_name", "type": "Text" } ]
		Commands           *[]DeviceAttribute `json:"commands,omitempty" bson:"commands"`                       // List of commands of the device	ex) [ { "name": "attr_name", "type": "Text" } ]
		InternalAttributes *[]interface{}     `json:"internal_attributes,omitempty" bson:"internal_attributes"` // List of internal attributes with free format for specific IoT Agent configuration ex) LWM2M mappings from object URIs to attributes
		StaticAttributes   *[]DeviceAttribute `json:"static_attributes,omitempty" bson:"static_attributes"`     // List of static attributes to append to the entity. All the updateContext requests to the CB will have this set of attributes appended. ex) [ { "name": "attr_name", "type": "Text" } ]
	}

	// DeviceAttribute holds device data model list item.
	DeviceAttribute struct {
		ObjectID *string       `json:"object_id,omitempty" bson:"object_id"`
		Name     string        `json:"name" bson:"name"`
		Type     string        `json:"type" bson:"type"`
		Value    *string       `json:"value,omitempty" bson:"value"`
		UnitMeta *UnitMetadata `json:"metadata,omitempty" bson:"metadata"`
	}

	PostDevices struct {
		Devices []Device `json:"devices"`
	}

	GetDevices struct {
		Count   int      `json:"count"`
		Devices []Device `json:"devices"`
	}
)

// https://github.com/telefonicaid/iotagent-node-lib/blob/master/doc/api.md#post-iotdevices
func (a *Accessor) CreateDevice(service, path string, devices PostDevices) error {
	// set service and service path into request.
	for i := range devices.Devices {
		devices.Devices[i].Service = &service
		devices.Devices[i].ServicePath = &path
	}
	req, err := gohttp.GenRequest(gohttp.HttpMethods.POST, a.genConfigUrl(DevicesUrl), devices)
	if err != nil {
		return err
	}
	res, err := a.Crud(AddIoTHeader(service, path, req))
	if err != nil {
		return err
	}
	if !gohttp.IsSuccessful(res) {
		return gohttp.ResponseToError(res, nil, nil)
	}
	return nil
}

// https://github.com/telefonicaid/iotagent-node-lib/blob/master/doc/api.md#get-iotdevices
func (a *Accessor) ReadDevices(service, path string, limit, offset *int) (*GetDevices, error) {
	req, err := http.NewRequest(http.MethodGet, a.genConfigUrl(DevicesUrl), nil)
	if err != nil {
		return nil, err
	}

	if limit != nil {
		gohttp.AddQueries(req, map[string]string{"limit": fmt.Sprintf("%d", *limit)})
	}
	if offset != nil {
		gohttp.AddQueries(req, map[string]string{"offset": fmt.Sprintf("%d", *offset)})
	}
	res, err := a.Crud(AddIoTHeader(service, path, req))
	if err != nil {
		return nil, err
	}
	if !gohttp.IsSuccessful(res) {
		return nil, gohttp.ResponseToError(res, nil, nil)
	}
	var gd GetDevices
	if err := gohttp.ResponseJSONToParams(res, &gd); err != nil {
		return nil, err
	}
	return &gd, nil
}

// https://github.com/telefonicaid/iotagent-node-lib/blob/master/doc/api.md#get-iotdevicesdeviceid
func (a *Accessor) ReadDevice(service, path, id string) (*Device, error) {
	req, err := http.NewRequest(http.MethodGet, a.genConfigUrl(DevicesUrl+"/"+id), nil)
	if err != nil {
		return nil, err
	}

	res, err := a.Crud(AddIoTHeader(service, path, req))
	if err != nil {
		return nil, err
	}

	if gohttp.IsNotFound(res) {
		return nil, DeviceNotFoundError
	}

	if !gohttp.IsSuccessful(res) {
		return nil, gohttp.ResponseToError(res, nil, nil)
	}

	var dev Device
	if err := gohttp.ResponseJSONToParams(res, &dev); err != nil {
		return nil, err
	}
	return &dev, nil
}

// https://github.com/telefonicaid/iotagent-node-lib/blob/master/doc/api.md#put-iotdevicesdeviceid
func (a *Accessor) UpdateDevice(service, path, id, json string) error {
	req, err := http.NewRequest(http.MethodPut, a.genConfigUrl(DevicesUrl+"/"+id), strings.NewReader(json))
	if err != nil {
		return err
	}

	res, err := a.Crud(AddIoTHeader(service, path, req))
	if err != nil {
		return err
	}

	if !gohttp.IsSuccessful(res) {
		return gohttp.ResponseToError(res, nil, nil)
	}
	return nil

}

// https://github.com/telefonicaid/iotagent-node-lib/blob/master/doc/api.md#delete-iotdevicesdeviceid
func (a *Accessor) DeleteDevice(service, path, id string) error {
	req, err := http.NewRequest(http.MethodDelete, a.genConfigUrl(DevicesUrl+"/"+id), nil)
	if err != nil {
		return err
	}
	res, err := a.Crud(AddIoTHeader(service, path, req))
	if err != nil {
		return err
	}

	if !gohttp.IsSuccessful(res) {
		return gohttp.ResponseToError(res, nil, nil)
	}
	return nil
}
