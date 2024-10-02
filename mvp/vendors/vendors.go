package vendors

import (
	"net/http"
	"rest-std-lib/mvp/managers"
)

type VendorConfig struct {
	Type       string                   `json:"type"`
	Route      string                   `json:"route"`
	Managers   []managers.ManagerConfig `json:"managers"`
	Properties map[string]interface{}   `json:"properties"`
}

type IVendor interface {
	Init(config VendorConfig, managers []managers.IManager) error
	GetEndpoints() string
	ServeHTTP(response http.ResponseWriter, request *http.Request)
}

type IVendorFactory interface {
	Create(config VendorConfig) (IVendor, error)
}
