// common/inter subsystem data
package common

import (
	"fmt"
	"net/http"
	"regexp"
)

const (
	ServiceHeader     = "fiware-service"
	ServicePathHeader = "fiware-servicepath"
)

var (
	InvalidServiceName     = fmt.Errorf("invalid service name")
	InvalidServicePath     = fmt.Errorf("invalid service path")

	ReServiceHeader = regexp.MustCompile(`^[_a-z0-9]{1,50}$`)
	ReServicePath   = regexp.MustCompile(`^/$|^/[_a-z0-9]{1,50}(/([_a-z0-9]{1,50})){0,9}$`) // max 10 path depth. TODO total 50 letter.
)

// IsValidService returns whether given strings suit to fiware-service or not.
func IsValidService(s string) bool {
	return ReServiceHeader.MatchString(s)
}

// IsValidServicePath returns whether given strings suit to fiware-service or not.
func IsValidServicePath(s string) bool {
	return ReServicePath.MatchString(s)
}

// AddServiceHeader(req, service, servicePath)
func AddServiceHeader(req *http.Request, service, servicePath string) error {
	if 0 < len(service) {
		if !IsValidService(service) {
			return InvalidServiceName
		}
		req.Header.Add(ServiceHeader, service)
	}
	if 0 < len(servicePath) {
		if !IsValidServicePath(servicePath) {
			return InvalidServicePath
		}
		req.Header.Add(ServicePathHeader, servicePath)
	}
	return nil
}
