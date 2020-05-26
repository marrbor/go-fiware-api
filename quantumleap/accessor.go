package quantumleap

import (
	"net/http"
)

type (
	// Accessor is base of Context Producer, Context Provider and Context Consumer.
	Accessor struct {
		HttpClient *http.Client
		BaseUrl    string
	}
)

// NewAccessor returns Producer instance.
func NewAccessor(baseUrl string) *Accessor {
	return &Accessor{
		HttpClient: new(http.Client),
		BaseUrl:    baseUrl,
	}
}
