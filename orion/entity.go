// based on https://fiware.github.io/specifications/ngsiv2/stable/
package orion

import (
	"fmt"

	"github.com/marrbor/gohttp"
)

/// Create

// CreateEntry create entity
func (a *Accessor) CreateEntity(service, servicePath string, q *Query, entity interface{}) error {
	return a.access(&AccessParameter{
		EpID:         EntryPointIDs.Entities,
		Method:       gohttp.HttpMethods.POST,
		Service:      service,
		ServicePath:  servicePath,
		Path:         "",
		Query:        q,
		BodyToSend:   entity,
		ReceivedBody: nil,
	})
}

/// Read

// GetEntityList gets entity list.
func (a *Accessor) GetEntityList(service, servicePath string, q *Query, entities interface{}) error {
	return a.access(&AccessParameter{
		EpID:         EntryPointIDs.Entities,
		Method:       gohttp.HttpMethods.GET,
		Service:      service,
		ServicePath:  servicePath,
		Path:         "",
		Query:        q,
		BodyToSend:   nil,
		ReceivedBody: entities,
	})
}

// GetEntity gets entity that has specified ID.
func (a *Accessor) GetEntity(service, servicePath, id string, q *Query, entity interface{}) error {
	return a.access(&AccessParameter{
		EpID:         EntryPointIDs.Entities,
		Method:       gohttp.HttpMethods.GET,
		Service:      service,
		ServicePath:  servicePath,
		Path:         fmt.Sprintf("/%s", id),
		Query:        q,
		BodyToSend:   nil,
		ReceivedBody: entity,
	})
}

// GetEntityAttribute gets specified attribute of specified entity.
func (a *Accessor) GetEntityAttribute(service, servicePath, id, attrName string, q *Query, attr interface{}) error {
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

// UpdateEntity updates entity
func (a *Accessor) UpdateEntity(service, servicePath, id, typeName string, param interface{}) error {
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

// UpdateEntityAttribute updates or append entity attribute.
func (a *Accessor) UpdateEntityAttribute(service, servicePath, id, attrName string, q *Query, attr interface{}) error {
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
func (a *Accessor) DeleteEntity(service, servicePath, id, typeName string) error {
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
