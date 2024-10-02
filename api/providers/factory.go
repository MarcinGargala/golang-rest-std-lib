package providers

import (
	"rest-std-lib/api/providers/states"
	"rest-std-lib/mvp/providers"
)

type ProviderFactory struct {
}

func (s *ProviderFactory) Create(config providers.ProviderConfig) (provider providers.IProvider, err error) {

	switch config.Type {
	case "providers.state.memory":
		return &states.InMemoryStateProvider{}, err
	case "providers.state.postgres":
		return &states.PostgresComponentStateProvider{}, err
	}

	return nil, nil
}
