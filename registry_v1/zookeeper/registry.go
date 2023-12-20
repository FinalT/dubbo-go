package zookeeper

import (
	"dubbo.apache.org/dubbo-go/v3/common"
	"dubbo.apache.org/dubbo-go/v3/registry_v1"
)

type zkRegistry struct {
	url *common.URL
}

func (r *zkRegistry) GetURL() *common.URL {
	return r.url
}

func (r *zkRegistry) IsAvailable() bool {
	//TODO implement me
	panic("implement me")
}

func (r *zkRegistry) Destroy() {
	//TODO implement me
	panic("implement me")
}

func (r *zkRegistry) Register(instance *registry.Instance, option ...registry.RegisterOption) error {
	//TODO implement me
	panic("implement me")
}

func (r *zkRegistry) Unregister(instance *registry.Instance, option ...registry.DeregisterOption) error {
	//TODO implement me
	panic("implement me")
}

func (r *zkRegistry) GetService(s string, option ...registry.GetOption) ([]*registry.Instance, error) {
	//TODO implement me
	panic("implement me")
}

func (r *zkRegistry) ListServices(option ...registry.ListOption) ([]*registry.Instance, error) {
	//TODO implement me
	panic("implement me")
}

func (r *zkRegistry) Watch(option ...registry.WatchOption) (registry.Watcher, error) {
	//TODO implement me
	panic("implement me")
}
