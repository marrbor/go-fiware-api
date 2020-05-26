/* fiware orion registration api
 * https://fiware.github.io/specifications/ngsiv2/stable/
 * TODO not tested yet
 */
package orion

import (
	"fmt"
	"time"

	"github.com/marrbor/gohttp"
)

type RegistrationProvider struct {
	Http                    Http   `json:"http"`
	SupportedForwardingMode string `json:"supportedForwardingMode,omitempty"`
}

type RegistrationDataProvided struct {
	Entities   []map[string]string `json:"entities"`
	Attrs      []string            `json:"attrs,omitempty"`
	Expression string              `json:"expression,omitempty"`
}

type RegistrationForwardingInformation struct {
	TimesSent        int64     `json:"timesSent,omitempty"`        // not editable, only present in GET operations
	LastNotification time.Time `json:"lastNotification,omitempty"` // not editable, only present in GET operations
	LastFailure      time.Time `json:"lastFailure,omitempty"`      // not editable, only present in GET operations
	LastSuccess      time.Time `json:"lastSuccess,omitempty"`      // not editable, only present in GET operations
}

type Registration struct {
	Id                    string                            `json:"id,omitempty"`
	Description           string                            `json:"description,omitempty"`
	Provider              RegistrationProvider              `json:"provider"`
	DataProvided          RegistrationDataProvided          `json:"dataProvided"`
	Status                string                            `json:"status,omitempty"`
	Expires               time.Time                         `json:"expires,omitempty"`
	ForwardingInformation RegistrationForwardingInformation `json:"forwardingInformation,omitempty"`
}

// CreateRegistration
func (a Accessor) CreateRegistration(service, servicePath string, q *Query, registration interface{}) error {
	return a.access(&AccessParameter{
		EpID:         EntryPointIDs.Registrations,
		Method:       gohttp.HttpMethods.POST,
		Service:      service,
		ServicePath:  servicePath,
		Path:         "",
		Query:        q,
		BodyToSend:   registration,
		ReceivedBody: nil,
	})
}

// GetRegistrationList
func (a Accessor) GetRegistrationList(service, servicePath string, q *Query, registrations interface{}) error {
	return a.access(&AccessParameter{
		EpID:         EntryPointIDs.Registrations,
		Method:       gohttp.HttpMethods.GET,
		Service:      service,
		ServicePath:  servicePath,
		Path:         "",
		Query:        q,
		BodyToSend:   nil,
		ReceivedBody: registrations,
	})
}

// GetRegistration gets registration that has specified ID.
func (a *Accessor) GetRegistration(service, servicePath, id string, q *Query, registration interface{}) error {
	return a.access(&AccessParameter{
		EpID:         EntryPointIDs.Entities,
		Method:       gohttp.HttpMethods.GET,
		Service:      service,
		ServicePath:  servicePath,
		Path:         fmt.Sprintf("/%s", id),
		Query:        q,
		BodyToSend:   nil,
		ReceivedBody: registration,
	})
}

// GetRegistrationAttribute gets specified attribute of specified registration.
func (a *Accessor) GetRegistrationAttribute(service, servicePath, id, attrName string, q *Query, attr interface{}) error {
	return a.access(&AccessParameter{
		EpID:         EntryPointIDs.Entities,
		Method:       gohttp.HttpMethods.GET,
		Service:      service,
		ServicePath:  servicePath,
		Path:         fmt.Sprintf("/%s/attrs/%s", id, attrName),
		Query:        q,
		BodyToSend:   nil,
		ReceivedBody: attr,
	})
}

/// Update

// UpdateRegistration updates registration
func (a *Accessor) UpdateRegistration(service, servicePath, id, typeName string, param interface{}) error {
	return a.access(&AccessParameter{
		EpID:         EntryPointIDs.Entities,
		Method:       gohttp.HttpMethods.PATCH,
		Service:      service,
		ServicePath:  servicePath,
		Path:         fmt.Sprintf("/%s/attrs", id),
		Query:        NewQuery().SetQuery("type", typeName),
		BodyToSend:   param,
		ReceivedBody: nil,
	})
}

// UpdateRegistrationAttribute updates or append registration attribute.
func (a *Accessor) UpdateRegistrationAttribute(service, servicePath, id, attrName string, q *Query, attr interface{}) error {
	return a.access(&AccessParameter{
		EpID:         EntryPointIDs.Entities,
		Method:       gohttp.HttpMethods.POST,
		Service:      service,
		ServicePath:  servicePath,
		Path:         fmt.Sprintf("/%s/attrs", id),
		Query:        q,
		BodyToSend:   attr,
		ReceivedBody: nil,
	})
}

/// Delete
func (a *Accessor) DeleteRegistration(service, servicePath, id, typeName string) error {
	return a.access(&AccessParameter{
		EpID:         EntryPointIDs.Entities,
		Method:       gohttp.HttpMethods.DELETE,
		Service:      service,
		ServicePath:  servicePath,
		Path:         fmt.Sprintf("/%s", id),
		Query:        NewQuery().SetQuery("type", typeName),
		BodyToSend:   nil,
		ReceivedBody: nil,
	})
}
