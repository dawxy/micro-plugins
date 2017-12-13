package consultags

import (
	"github.com/micro/go-micro/registry"
	"golang.org/x/net/context"
)

func Tags(tags ...string) registry.Option {
	return func(o *registry.Options) {
		if o.Context == nil {
			o.Context = context.Background()
		}
		o.Context = context.WithValue(o.Context, "tags", tags)
	}
}
