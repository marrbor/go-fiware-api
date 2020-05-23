// Service Group API
// https://github.com/telefonicaid/iotagent-node-lib/blob/master/doc/api.md#service-group-api
package iotagent

import (
	"net/http"
	"strings"

	"github.com/marrbor/gohttp"
)

const (
	ServiceGroupUrl       = "iot/services"
	UltraLightResourceUrl = "iot/d"
	TransportHttp         = "HTTP"
	JsonProtocol          = "json"
)

type (
	APIServiceGroup struct {
		Services []ServiceGroup `json:"services"`
	}

	ServiceGroup struct {
		Service            string            `json:"service,omitempty" bson:"service"`                         // Service of the devices of this type
		SubService         string            `json:"subservice,omitempty" bson:"subservice"`                   // Subservice of the devices of this type.
		Resource           string            `json:"resource" bson:"resource"`                                 // string representing the Southbound resource that will be used to assign a type to a device (e.g.: pathname in the southbound port).
		APIKey             string            `json:"apikey" bson:"apikey"`                                     // API Key string.
		TimeStamp          bool              `json:"timestamp,omitempty" bson:"timestamp"`                     // Optional flagw whether to include the TimeInstantwithin each entity created, as well as a TimeInstant metadata to each attribute, with the current timestamp
		EntityType         string            `json:"entity_type" bson:"entity_type"`                           // name of the Entity type to assign to the group.
		Trust              string            `json:"trust,omitempty" bson:"trust"`                             // trust token to use for secured access to the Context Broker for this type of devices (optional; only needed for secured scenarios).
		CbHost             string            `json:"cbHost" bson:"cbHost"`                                     // Context Broker connection information. This options can be used to override the global ones for specific types of devices.
		Lazy               []DeviceAttribute `json:"lazy,omitempty" bson:"lazy"`                               // list of common lazy attributes of the device. For each attribute, its name and type must be provided.
		Commands           []DeviceAttribute `json:"commands,omitempty" bson:"commands"`                       // list of common commands attributes of the device. For each attribute, its name and type must be provided, additional metadata is optional.
		Attributes         []DeviceAttribute `json:"attributes,omitempty" bson:"attributes"`                   // list of common active attributes of the device. For each attribute, its name and type must be provided, additional metadata is optional.
		StaticAttributes   []DeviceAttribute `json:"static_attributes,omitempty" bson:"static_attributes"`     // this attributes will be added to all the entities of this group 'as is', additional metadata is optional.
		InternalAttributes []interface{}     `json:"internal_attributes,omitempty" bson:"internal_attributes"` // optional section with free format, to allow specific IoT Agents to store information along with the devices in the Device Registry.
	}
)

// CreateServiceGroup registers given service group into iot agent.
func (a *Accessor) CreateServiceGroup(service, path string, body *APIServiceGroup) error {
	// set service and service path into request.
	for i := range body.Services {
		body.Services[i].Service = service
		body.Services[i].SubService = path
	}
	req, err := gohttp.GenRequest(gohttp.HttpMethods.POST, a.genConfigUrl(ServiceGroupUrl), body)
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

// ReadServiceGroup retrieves service groups from the iot agent.
func (a *Accessor) ReadServiceGroup(service, path string) (*APIServiceGroup, error) {
	req, err := gohttp.GenRequest(gohttp.HttpMethods.GET, a.genConfigUrl(ServiceGroupUrl), nil)
	if err != nil {
		return nil, err
	}

	res, err := a.Crud(AddIoTHeader(service, path, req))
	if err != nil {
		return nil, err
	}

	if !gohttp.IsSuccessful(res) {
		return nil, gohttp.ResponseToError(res, nil, nil)
	}

	var asg APIServiceGroup
	if err := gohttp.ResponseJSONToParams(res, &asg); err != nil {
		return nil, err
	}
	return &asg, nil
}

// UpdateServiceGroup modifies the information for a service group configuration, identified by the resource and apikey query parameters.
func (a *Accessor) UpdateServiceGroup(resource, apikey, json string) error {
	req, err := http.NewRequest(http.MethodPut, a.genConfigUrl(ServiceGroupUrl), strings.NewReader(json))
	if err != nil {
		return err
	}

	res, err := a.Crud(gohttp.AddQueries(req, map[string]string{"resource": resource, "apikey": apikey}))
	if err != nil {
		return err
	}

	if !gohttp.IsSuccessful(res) {
		return gohttp.ResponseToError(res, nil, nil)
	}
	return nil
}

// DeleteServiceGroup deletes specified group.
func (a *Accessor) DeleteServiceGroup(resource, apikey string) error {
	req, err := http.NewRequest(http.MethodDelete, a.genConfigUrl(ServiceGroupUrl), nil)
	if err != nil {
		return err
	}
	res, err := a.Crud(gohttp.AddQueries(req, map[string]string{"resource": resource, "apikey": apikey}))
	if err != nil {
		return err
	}

	if !gohttp.IsSuccessful(res) {
		return gohttp.ResponseToError(res, nil, nil)
	}
	return nil
}
