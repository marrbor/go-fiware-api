package iotagent

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/marrbor/gohttp"
	"github.com/marrbor/golog"
)

const (
	JsonResourceUrl = "iot/json"
)

// https://fiware-iotagent-json.letsfiware.jp/usermanual/index.html#_3
func (a *Accessor) SendJsonReport(service, path, key, id string, report interface{}) error {
	req, err := gohttp.GenRequest(gohttp.HttpMethods.POST, a.genReportUrl(JsonResourceUrl), report)
	if err != nil {
		return err
	}
	q := req.URL.Query()
	q.Add("k", key)
	q.Add("i", id)
	req.URL.RawQuery = q.Encode()
	golog.Info(fmt.Sprintf("send report(%+v) to url (%s)", report, req.URL.String()))

	res, err := a.Crud(AddIoTHeader(service, path, req))
	if err != nil {
		return err
	}

	if !gohttp.IsSuccessful(res) {
		return gohttp.ResponseToError(res, nil, nil)
	}
	return nil
}

// https://fiware-iotagent-json.letsfiware.jp/usermanual/index.html#_3
func (a *Accessor) SendJsonTextReport(service, path, key, id, report string) error {
	req, err := http.NewRequest(http.MethodPost, a.genReportUrl(JsonResourceUrl), strings.NewReader(report))
	if err != nil {
		return err
	}
	q := req.URL.Query()
	q.Add("k", key)
	q.Add("i", id)
	req.URL.RawQuery = q.Encode()
	golog.Info(fmt.Sprintf("send report(%+v) to url (%s)", report, req.URL.String()))

	res, err := a.Crud(AddIoTHeader(service, path, req))
	if err != nil {
		return err
	}

	if !gohttp.IsSuccessful(res) {
		return gohttp.ResponseToError(res, nil, nil)
	}
	return nil
}
