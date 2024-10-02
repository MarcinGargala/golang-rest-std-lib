package bindings

import "rest-std-lib/mvp/vendors"

type HttpBindingConfig struct {
	Port string
}
type IBinding interface {
	Launch(config HttpBindingConfig, vendors []vendors.IVendor)
}
