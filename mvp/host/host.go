package host

import (
	"log/slog"
	"rest-std-lib/mvp/bindings"
	"rest-std-lib/mvp/managers"
	"rest-std-lib/mvp/providers"
	"rest-std-lib/mvp/vendors"
)

type Config struct {
	API      APIConfig
	Bindings BindingConfig
}

type APIConfig struct {
	Vendors []vendors.VendorConfig
}

type BindingConfig struct {
	Type   string                 `json:"type"`
	Config map[string]interface{} `json:"config"`
}
type APIHost struct {
	config          Config
	vendorFactory   vendors.IVendorFactory
	managerFactory  managers.IManagerFactory
	providerFactory providers.IProviderFactory
}

func (s *APIHost) Init(config Config,
	vendorFactory vendors.IVendorFactory,
	managerFactory managers.IManagerFactory,
	providerFactory providers.IProviderFactory) (err error) {

	s.config = config
	s.vendorFactory = vendorFactory
	s.managerFactory = managerFactory
	s.providerFactory = providerFactory
	return
}

func (s *APIHost) Launch() error {

	var createdVendors []vendors.IVendor
	for _, vendorConfig := range s.config.API.Vendors {
		var createdManagers []managers.IManager
		for _, managerConfig := range vendorConfig.Managers {
			var createdProviders = make(map[string]providers.IProvider)
			for k, providerConfig := range managerConfig.Providers {
				provider, err := s.providerFactory.Create(providerConfig)
				err2 := provider.Init(providerConfig)
				if err != nil || err2 != nil {
					slog.Warn("Can not create provider from ", "config", providerConfig)
					panic("Launching error on Provider creation")
				}
				createdProviders[k] = provider
			}
			manager, err := s.managerFactory.Create(managerConfig)
			err2 := manager.Init(managerConfig, createdProviders)
			if err != nil || err2 != nil {
				slog.Warn("Can not create manager from ", "config", managerConfig)
				panic("Launching error on Manager creation")
			}
			createdManagers = append(createdManagers, manager)
		}
		vendor, err := s.vendorFactory.Create(vendorConfig)
		err2 := vendor.Init(vendorConfig, createdManagers)
		if err != nil || err2 != nil {
			slog.Warn("Can not create vendor from ", "config", vendorConfig)
			panic("Launching error on Vendor creation")

		}
		createdVendors = append(createdVendors, vendor)
	}

	switch s.config.Bindings.Type {
	case "bindings.http":
		var http = bindings.HttpBinding{}
		var httpConfig = bindings.HttpBindingConfig{
			Port: s.config.Bindings.Config["port"].(string),
		}
		http.Init(httpConfig)
		http.Launch(createdVendors)
	}
	return nil
}
