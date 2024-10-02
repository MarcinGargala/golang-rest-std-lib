package managers

import "rest-std-lib/mvp/providers"

type ManagerConfig struct {
	Name       string                              `json:"name"`
	Type       string                              `json:"type"`
	Providers  map[string]providers.ProviderConfig `json:"providers"`
	Properties map[string]string                   `json:"properties"`
}
type IManager interface {
	Init(config ManagerConfig, providers map[string]providers.IProvider) error
}

type IManagerFactory interface {
	Create(config ManagerConfig) (IManager, error)
}
