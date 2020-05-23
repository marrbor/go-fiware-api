// API accessor
package iotagent

import (
	"fmt"
	"net/http"
)

// Accessor holds http client instance, base url of iot agent.
type Accessor struct {
	client    *http.Client
	configUrl string
	reportUrl string
}

// genConfigUrl returns configUrl + givenUrl.
func (a *Accessor) genConfigUrl(url string) string {
	return fmt.Sprintf("%s/%s", a.configUrl, url)
}

// genReportUrl returns reportUrl + givenUrl.
func (a *Accessor) genReportUrl(url string) string {
	return fmt.Sprintf("%s/%s", a.reportUrl, url)
}

// Crud do api access.
func (a *Accessor) Crud(req *http.Request) (*http.Response, error) {
	return a.client.Do(req)
}

func NewAccessor(configUrl, reportUrl string) *Accessor {
	return &Accessor{
		client:    new(http.Client),
		configUrl: configUrl,
		reportUrl: reportUrl,
	}
}
