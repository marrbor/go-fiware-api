// Quantum
package quantumleap

import (
	"fmt"
	"net/http"

	"github.com/marrbor/gohttp"
)

type (
	Version struct {
		Version string `json:"version"`
	}
)

// GetVersion gets version information from the server.
func (a Accessor) GetVersion() (*Version, error){
	req, err := gohttp.GenRequest(gohttp.HttpMethods.GET, fmt.Sprintf("%s/v2/version", a.BaseUrl), nil)
	if err != nil {
		return nil, err
	}
	res, err := a.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}
	if http.StatusBadRequest <= res.StatusCode {
		return nil, fmt.Errorf(res.Status)
	}
	var v Version
	if err := gohttp.ResponseJSONToParams(res, &v); err != nil {
		return nil, err
	}
	return &v, nil
}
