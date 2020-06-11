package orion

import (
	"fmt"
	"net/http"
	"net/url"
	"path"

	"github.com/marrbor/go-fiware-api/common"
	"github.com/marrbor/gohttp"
)

type (
	// Accessor is base of Context Producer, Context Provider and Context Consumer.
	Accessor struct {
		HttpClient  *http.Client
		BaseUrl     string
		EntryPoints *EntryPoints
	}

	// Access Parameter holds parameter for Orion server access.
	AccessParameter struct {
		EpID           EntryPointID
		Method         gohttp.HTTPMethod
		Service        string
		ServicePath    string
		Path           string
		Query          *Query
		BodyToSend     interface{}
		ReceivedBody   interface{}
		ReceivedHeader http.Header
	}
)

var (
	IllegalEndPointIDError = fmt.Errorf("illegal end point ID")
)

// NewAccessor returns Producer instance.
func NewAccessor(baseUrl string) *Accessor {
	a := Accessor{
		HttpClient:  new(http.Client),
		BaseUrl:     baseUrl,
		EntryPoints: nil,
	}

	// try to get EntryPoints. Ignore error since take it later when failed here.
	a.EntryPoints, _ = a.GetEntryPoints()
	return &a
}

// genBaseURL returns strings url instance included entry point.
func (a *Accessor) genBaseURL(epID EntryPointID) (*url.URL, error) {
	u, err := url.Parse(a.BaseUrl)
	if err != nil {
		return nil, err
	}

	if a.EntryPoints == nil {
		if err := a.ReloadEntryPoint(); err != nil {
			return nil, err
		}
	}

	switch epID.Value() {
	case Entities:
		u.Path = path.Join(u.Path, a.EntryPoints.EntitiesURL)
	case Types:
		u.Path = path.Join(u.Path, a.EntryPoints.TypesURL)
	case Subscriptions:
		u.Path = path.Join(u.Path, a.EntryPoints.SubscriptionsURL)
	case Registrations:
		u.Path = path.Join(u.Path, a.EntryPoints.RegistrationsURL)
	default:
		return nil, IllegalEndPointIDError
	}
	return u, nil
}

// genURLString generates url string to make requests.
func (a *Accessor) genURLString(epID EntryPointID, pathTo string, q *Query) (string, error) {
	u, err := a.genBaseURL(epID)
	if err != nil {
		return "", err
	}

	u.Path = path.Join(u.Path, pathTo)
	if q != nil {
		q.SetToURL(u)
	}
	return u.String(), nil
}

func (a *Accessor) access(ap *AccessParameter) error {
	// Generate request
	uri, err := a.genURLString(ap.EpID, ap.Path, ap.Query)
	if err != nil {
		return err
	}
	req, err := gohttp.GenRequest(ap.Method, uri, ap.BodyToSend)
	if err != nil {
		return err
	}
	if err := common.AddServiceHeader(req, ap.Service, ap.ServicePath); err != nil {
		return err
	}

	// Send request
	res, err := a.HttpClient.Do(req)
	if err != nil {
		return err
	}
	if http.StatusBadRequest <= res.StatusCode {
		return fmt.Errorf(res.Status)
	}

	// set response header to return.
	ap.ReceivedHeader = res.Header

	// has response body?
	if ap.ReceivedBody == nil {
		return nil
	}

	// Parse the response
	return gohttp.ResponseJSONToParams(res, ap.ReceivedBody)
}
