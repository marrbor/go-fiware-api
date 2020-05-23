package orion

import (
	"fmt"
	"net/http"
	"time"

	"github.com/marrbor/gohttp"
)

const (
	UptimeFormatInVersion = "%d d, %d h, %d m, %d s"
)

// Version is a structure that holds version information.
type Version struct {
	Version     string        `json:"version"`
	Uptime      time.Duration `json:"uptime"`
	GitHash     string        `json:"git_hash"`
	CompileTime time.Time     `json:"compile_time"`
	CompiledBy  string        `json:"compiled_by"`
	CompiledIn  string        `json:"compiled_in"`
	ReleaseDate time.Time     `json:"release_date"`
	Doc         string        `json:"doc"`
}

// GetVersion gets version information from the server.
func (a *Accessor) GetVersion() (*Version, error) {
	req, err := gohttp.GenRequest(gohttp.HttpMethods.GET, fmt.Sprintf("%s/version", a.BaseUrl), nil)
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
	var va VersionAPI
	if err := gohttp.ResponseJSONToParams(res, &va); err != nil {
		return nil, err
	}

	v, err := va.ToVersion()
	if err != nil {
		return nil, err
	}

	return v, nil
}

// VersionAPI is a structure that received from server response
type VersionAPI struct {
	Orion VersionStrings `json:"orion"`
}

// VersionStrings is a structure that received from server response
type VersionStrings struct {
	Version     string `json:"version"`
	Uptime      string `json:"uptime"`
	GitHash     string `json:"git_hash"`
	CompileTime string `json:"compile_time"`
	CompiledBy  string `json:"compiled_by"`
	CompiledIn  string `json:"compiled_in"`
	ReleaseDate string `json:"release_date"`
	Doc         string `json:"doc"`
}

func (va *VersionAPI) ToVersion() (*Version, error) {
	o := va.Orion

	// convert uptime strings to time duration.
	var dd, hh, mm, ss int64
	format := UptimeFormatInVersion
	if _, err := fmt.Sscanf(o.Uptime, format, &dd, &hh, &mm, &ss); err != nil {
		return nil, err
	}

	ct, err := time.Parse(time.UnixDate, o.CompileTime)
	if err != nil {
		return nil, err
	}

	rd, err := time.Parse(time.UnixDate, o.ReleaseDate)
	if err != nil {
		return nil, err
	}

	v := Version{
		Version: o.Version,
		Uptime: time.Duration(dd)*24*time.Hour +
			time.Duration(hh)*time.Hour +
			time.Duration(mm)*time.Minute +
			time.Duration(ss)*time.Second,
		GitHash:     o.GitHash,
		CompileTime: ct,
		CompiledBy:  o.CompiledBy,
		CompiledIn:  o.CompiledIn,
		ReleaseDate: rd,
		Doc:         o.Doc,
	}
	return &v, nil
}
