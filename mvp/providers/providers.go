package providers

type ProviderConfig struct {
	Name string `json:"name"`
	Type string `json:"type"`
}
type IProvider interface {
	Init(config ProviderConfig) error
}

type IProviderFactory interface {
	Create(config ProviderConfig) (IProvider, error)
}
