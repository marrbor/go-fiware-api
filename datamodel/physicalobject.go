/*
 * https://fiware-datamodels.readthedocs.io/en/latest/howto/index.html
 * https://github.com/fiware/data-models/common-schema.json #PhysicalObject-Commons
 *
 */
package datamodel

import "net/url"

type (
	PhysicalObject struct {
		Color       string   `json:"color"`
		Image       string   `json:"image"` // "format": "uri"
		Annotations []string `json:"annotations"`
	}
)

// ValidateImage validates whether image string is valid uri or not
func ValidatePhysicalObjectImageString(image string) error {
	_, err := url.ParseRequestURI(image)
	return err
}
