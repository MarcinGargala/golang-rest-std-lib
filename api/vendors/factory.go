package vendors

import (
	"rest-std-lib/mvp/vendors"
)

type VendorFactory struct {
}

func (s *VendorFactory) Create(config vendors.VendorConfig) (vendor vendors.IVendor, err error) {
	switch config.Type {
	case "vendors.components":
		componentVendor := &ComponentsVendor{}
		return componentVendor, nil
	}

	return nil, nil
}
