package registry

import (
	"dubbo.apache.org/dubbo-go/v3/global"
	"fmt"
)

type Options struct {
	Registry *global.RegistryConfig

	ID string
}

func defaultOptions() *Options {
	return &Options{
		Registry: global.DefaultRegistryConfig(),
	}
}

func NewOptions(opts ...Option) *Options {
	defOpts := defaultOptions()
	for _, opt := range opts {
		opt(defOpts)
	}

	if defOpts.Registry.Protocol == "" {
		panic(fmt.Sprintf("Please specify registry, eg. WithZookeeper()"))
	}
	if defOpts.ID == "" {
		defOpts.ID = defOpts.Registry.Protocol
	}

	return defOpts
}

type Option func(*Options)

type RegisterOptions struct {
}

type WatchOptions struct {
}

type DeregisterOptions struct {
}

type GetOptions struct {
}

type ListOptions struct {
}

type RegisterOption func(*RegisterOptions)

type WatchOption func(*WatchOptions)

type DeregisterOption func(*DeregisterOptions)

type GetOption func(*GetOptions)

type ListOption func(*ListOptions)
