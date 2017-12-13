package consultags

import (
	"github.com/micro/go-micro/registry"
)

func NewRegistry(opts ...registry.Option) registry.Registry {
	return newConsulRegistry(opts...)
}
