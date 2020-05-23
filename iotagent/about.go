package iotagent

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/marrbor/gohttp"
)

const (
	AboutUrl = "iot/about"
)

type (
	About struct {
		LibVersion string `json:"libVersion"`
		Port       int    `json:"port"`
		BaseRoot   string `json:"baseRoot"`
		Version    string `json:"version"`
	}

	APIAbout struct {
		LibVersion string `json:"libVersion"`
		Port       string `json:"port"`
		BaseRoot   string `json:"baseRoot"`
		Version    string `json:"version"`
	}
)

func (av *APIAbout) toAbout() (*About, error) {
	port, err := strconv.Atoi(av.Port)
	if err != nil {
		return nil, err
	}
	return &About{
		LibVersion: av.LibVersion,
		Port:       port,
		BaseRoot:   av.BaseRoot,
		Version:    av.Version,
	}, nil
}

// ReadAbout gets IoT Agent information. It can be used as a heartbeat operation to check the health of the IoT Agent if required.
func (a *Accessor) ReadAbout() (*About, error) {
	req, err := http.NewRequest(http.MethodGet, a.genConfigUrl(AboutUrl), nil)
	if err != nil {
		return nil, err
	}
	res, err := a.Crud(req)
	if err != nil {
		return nil, err
	}

	if !gohttp.IsSuccessful(res) {
		return nil, fmt.Errorf(res.Status)
	}

	var av APIAbout
	if err := gohttp.ResponseJSONToParams(res, &av); err != nil {
		return nil, err
	}
	return av.toAbout()
}
