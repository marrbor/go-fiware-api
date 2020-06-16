/* fiware orion subscription api
 * https://fiware.github.io/specifications/ngsiv2/stable/
 */
package orion

import (
	"fmt"
	"path"
	"time"

	"github.com/marrbor/gohttp"
)

type SubscriptionSubjectCondition struct {
	Attrs      *[]string `json:"attrs,omitempty"`
	Expression *string   `json:"expression,omitempty"`
}

type SubscriptionSubject struct {
	Entities  []map[string]string           `json:"entities"`
	Condition *SubscriptionSubjectCondition `json:"condition,omitempty"`
}

type SubscriptionHttpCustom struct {
	Url     string               `json:"url"`
	Headers *[]map[string]string `json:"headers,omitempty"`
	Qs      *[]map[string]string `json:"qs,omitempty"`
	Method  *string              `json:"method,omitempty"`
	Payload *string              `json:"payload,omitempty"`
}

// TODO. add optional attributes `attrsFormat`
type SubscriptionNotification struct {
	Attrs            *[]string               `json:"attrs,omitempty"`
	MetaData         *[]string               `json:"metadata,omitempty"`
	ExceptAttrs      *[]string               `json:"exceptAttrs,omitempty"`
	Http             *Http                   `json:"http,omitempty"`
	HttpCustom       *SubscriptionHttpCustom `json:"httpCustom,omitempty"`
	TimesSent        *int64                  `json:"timesSent,omitempty"`        // not editable, only present in GET operations
	LastNotification *time.Time              `json:"lastNotification,omitempty"` // not editable, only present in GET operations
	LastFailure      *time.Time              `json:"lastFailure,omitempty"`      // not editable, only present in GET operations
	LastSuccess      *time.Time              `json:"lastSuccess,omitempty"`      // not editable, only present in GET operations
}

type Subscription struct {
	Id           *string                  `json:"id,omitempty"`
	Description  *string                  `json:"description,omitempty"`
	Subject      SubscriptionSubject      `json:"subject"`
	Notification SubscriptionNotification `json:"notification"`
	Expires      *time.Time               `json:"expires,omitempty"`
	Status       *string                  `json:"status,omitempty"`
	Throttling   *int64                   `json:"throttling,omitempty"`
}

// CreateSubscription post request to create subscription and return subscription ID or error.
func (a Accessor) CreateSubscription(service, servicePath string, subscription *Subscription) (string, error) {
	ap := AccessParameter{
		EpID:        EntryPointIDs.Subscriptions,
		Method:      gohttp.HttpMethods.POST,
		Service:     service,
		ServicePath: servicePath,
		Path:        "",
		BodyToSend:  subscription,
	}

	err := a.access(&ap)
	if err != nil {
		return "", err
	}
	return path.Base(ap.ReceivedHeader.Get("Location")), err
}

// GetSubscriptionList gets current subscription list
func (a Accessor) GetSubscriptionList(service, servicePath string, subscriptions *[]Subscription) error {
	return a.access(&AccessParameter{
		EpID:         EntryPointIDs.Subscriptions,
		Method:       gohttp.HttpMethods.GET,
		Service:      service,
		ServicePath:  servicePath,
		Path:         "",
		ReceivedBody: subscriptions,
	})
}

// GetSubscription gets subscription that has specified ID.
func (a *Accessor) GetSubscription(service, servicePath, id string, subscription *Subscription) error {
	return a.access(&AccessParameter{
		EpID:         EntryPointIDs.Subscriptions,
		Method:       gohttp.HttpMethods.GET,
		Service:      service,
		ServicePath:  servicePath,
		Path:         fmt.Sprintf("/%s", id),
		ReceivedBody: subscription,
	})
}

// GetSubscriptionAttribute gets specified attribute of specified subscription.
func (a *Accessor) GetSubscriptionAttribute(service, servicePath, id, attrName string, q *Query, attr interface{}) error {
	return a.access(&AccessParameter{
		EpID:         EntryPointIDs.Subscriptions,
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

// UpdateSubscription updates subscription
func (a *Accessor) UpdateSubscription(service, servicePath, id, typeName string, param interface{}) error {
	return a.access(&AccessParameter{
		EpID:         EntryPointIDs.Subscriptions,
		Method:       gohttp.HttpMethods.PATCH,
		Service:      service,
		ServicePath:  servicePath,
		Path:         fmt.Sprintf("/%s/attrs", id),
		Query:        NewQuery().SetQuery("type", typeName),
		BodyToSend:   param,
		ReceivedBody: nil,
	})
}

// UpdateSubscriptionAttribute updates or append subscription attribute.
func (a *Accessor) UpdateSubscriptionAttribute(service, servicePath, id, attrName string, q *Query, attr interface{}) error {
	return a.access(&AccessParameter{
		EpID:         EntryPointIDs.Subscriptions,
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
func (a *Accessor) DeleteSubscription(service, servicePath, id, typeName string) error {
	return a.access(&AccessParameter{
		EpID:         EntryPointIDs.Subscriptions,
		Method:       gohttp.HttpMethods.DELETE,
		Service:      service,
		ServicePath:  servicePath,
		Path:         fmt.Sprintf("/%s", id),
		Query:        NewQuery().SetQuery("type", typeName),
		BodyToSend:   nil,
		ReceivedBody: nil,
	})
}
