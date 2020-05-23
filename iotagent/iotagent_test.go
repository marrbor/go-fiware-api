package iotagent

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVersion(t *testing.T) {
	av := APIAbout{
		LibVersion: "abc",
		Port:       "123",
		BaseRoot:   "/",
		Version:    "1.1.1",
	}
	v, err := av.toAbout()
	assert.NoError(t, err)
	assert.EqualValues(t, av.Version, v.Version)
	assert.EqualValues(t, av.BaseRoot, v.BaseRoot)
	assert.EqualValues(t, av.LibVersion, v.LibVersion)
	assert.EqualValues(t, 123, v.Port)

	av.Port = "abc"
	v, err = av.toAbout()
	assert.EqualError(t, err, "strconv.Atoi: parsing \"abc\": invalid syntax")

	av.Port = ""
	v, err = av.toAbout()
	assert.EqualError(t, err, "strconv.Atoi: parsing \"\": invalid syntax")
}

func TestAddIoTHeader(t *testing.T) {
	r, err := http.NewRequest(http.MethodPost, "", nil)
	assert.NoError(t, err)
	r2 := AddIoTHeader("a", "b", r)
	assert.EqualValues(t, "a", r2.Header.Get("Fiware-Service"))
	assert.EqualValues(t, "b", r2.Header.Get("Fiware-ServicePath"))
	assert.EqualValues(t, "no-cache", r2.Header.Get("Cache-Control"))
	assert.EqualValues(t, "application/json", r2.Header.Get("Content-Type"))
	assert.EqualValues(t, "", r2.Header.Get("Accept"))
}

func TestAddIoTHeader2(t *testing.T) {
	r, err := http.NewRequest(http.MethodPut, "", nil)
	assert.NoError(t, err)
	r2 := AddIoTHeader("a", "b", r)
	assert.EqualValues(t, "a", r2.Header.Get("Fiware-Service"))
	assert.EqualValues(t, "b", r2.Header.Get("Fiware-ServicePath"))
	assert.EqualValues(t, "no-cache", r2.Header.Get("Cache-Control"))
	assert.EqualValues(t, "application/json", r2.Header.Get("Content-Type"))
	assert.EqualValues(t, "", r2.Header.Get("Accept"))

}

func TestAddIoTHeader3(t *testing.T) {
	r, err := http.NewRequest(http.MethodGet, "", nil)
	assert.NoError(t, err)
	r2 := AddIoTHeader("a", "b", r)
	assert.EqualValues(t, "a", r2.Header.Get("Fiware-Service"))
	assert.EqualValues(t, "b", r2.Header.Get("Fiware-ServicePath"))
	assert.EqualValues(t, "no-cache", r2.Header.Get("Cache-Control"))
	assert.EqualValues(t, "", r2.Header.Get("Content-Type"))
	assert.EqualValues(t, "application/json", r2.Header.Get("Accept"))
}

func TestAddIoTHeader4(t *testing.T) {
	r, err := http.NewRequest(http.MethodDelete, "", nil)
	assert.NoError(t, err)
	r2 := AddIoTHeader("a", "b", r)
	assert.EqualValues(t, "a", r2.Header.Get("Fiware-Service"))
	assert.EqualValues(t, "b", r2.Header.Get("Fiware-ServicePath"))
	assert.EqualValues(t, "no-cache", r2.Header.Get("Cache-Control"))
	assert.EqualValues(t, "", r2.Header.Get("Content-Type"))
	assert.EqualValues(t, "", r2.Header.Get("Accept"))
}

func TestNewAccessor(t *testing.T) {
	a := NewAccessor("conf", "repo")
	assert.EqualValues(t, a.configUrl, "conf")
	assert.EqualValues(t, a.reportUrl, "repo")
}

func TestDevice(t *testing.T) {
	d := new(Device)
	j, err := json.MarshalIndent(d, "", " ")
	assert.NoError(t, err)
	t.Log(string(j))

	da := []DeviceAttribute{
		CpuUsage,
		DiskUsage,
		MemoryUsage,
		LoadAverage1,
		LoadAverage5,
		LoadAverage15,
		Processes,
		Uptime,
	}

	d.Attributes = da
	j, err = json.MarshalIndent(d, "", " ")
	assert.NoError(t, err)
	t.Log(string(j))
}
