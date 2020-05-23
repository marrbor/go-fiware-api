package iotagent

import (
	"fmt"
	"net/http"

	"github.com/marrbor/gohttp"
)

const (
	LogUrl = "/admin/log"
)

type (
	Level struct {
		Level string `json:"level"`
	}
)

// https://github.com/telefonicaid/iotagent-node-lib/blob/master/doc/api.md#put-adminlog-1
func (a *Accessor) ReadLogLevel() (string, error) {
	req, err := http.NewRequest(http.MethodGet, a.configUrl+LogUrl, nil)
	if err != nil {
		return "", err
	}
	res, err := a.Crud(req)
	if err != nil {
		return "", err
	}
	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf(res.Status)
	}
	var level Level
	if err := gohttp.ResponseJSONToParams(res, &level); err != nil {
		return "", err
	}
	return level.Level, nil
}

// UpdateLogLevel updates log level. If the new level is a valid level for Logops (i.e.: one of the items in the array ['INFO', 'ERROR', 'FATAL', 'DEBUG', 'WARNING']), it will be automatically changed for future logs.
// https://github.com/telefonicaid/iotagent-node-lib/blob/master/doc/api.md#put-adminlogp
func (a *Accessor) UpdateLogLevel(level string) error {
	req, err := http.NewRequest(http.MethodPut, a.genConfigUrl(LogUrl), nil)
	if err != nil {
		return err
	}

	res, err := a.Crud(gohttp.AddQueries(req, map[string]string{"level": level}))
	if err != nil {
		return err
	}

	if !gohttp.IsSuccessful(res) {
		return fmt.Errorf(res.Status)
	}
	return nil
}
