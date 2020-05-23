// https://open-apis.letsfiware.jp/fiware-orion/api/v2/stable/	#API Entry Point
package orion

import (
	"fmt"
	"net/http"

	"github.com/marrbor/gohttp"
)

// Entry Point ID
const (
	Entities = iota
	Types
	Subscriptions
	Registrations
)

type EntryPointID struct{ value int }

func (e EntryPointID) Value() int {
	return e.value
}

var EntryPointIDs = struct {
	Entities      EntryPointID
	Types         EntryPointID
	Subscriptions EntryPointID
	Registrations EntryPointID
}{
	Entities:      EntryPointID{Entities},
	Types:         EntryPointID{Types},
	Subscriptions: EntryPointID{Subscriptions},
	Registrations: EntryPointID{Registrations},
}

type (
	EntryPoints struct {
		EntitiesURL      string `json:"entities_url"`
		TypesURL         string `json:"types_url"`
		SubscriptionsURL string `json:"subscriptions_url"`
		RegistrationsURL string `json:"registrations_url"`
	}
)

// ReloadEntryPoint load entry point list from server.
func (a *Accessor) ReloadEntryPoint() error {
	req, err := gohttp.GenRequest(gohttp.HttpMethods.GET, fmt.Sprintf("%s/v2", a.BaseUrl), nil)
	if err != nil {
		return err
	}
	res, err := a.HttpClient.Do(req)
	if err != nil {
		return err
	}
	if http.StatusBadRequest <= res.StatusCode {
		return fmt.Errorf(res.Status)
	}
	var ep EntryPoints
	if err := gohttp.ResponseJSONToParams(res, &ep); err != nil {
		return err
	}
	a.EntryPoints = &ep
	return nil
}

// GetEntryPoints gets api resources from Orion.
func (a *Accessor) GetEntryPoints() (*EntryPoints, error) {
	if err := a.ReloadEntryPoint(); err != nil {
		return nil, err
	}
	return a.EntryPoints, nil
}
