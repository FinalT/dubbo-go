package interface_registry

import (
	"dubbo.apache.org/dubbo-go/v3/cluster/router"
	"dubbo.apache.org/dubbo-go/v3/common"
	"dubbo.apache.org/dubbo-go/v3/directory_v1"
	"dubbo.apache.org/dubbo-go/v3/protocol"
	"dubbo.apache.org/dubbo-go/v3/registry_v1"
)

// router is the default router.
type interfaceDirectory struct {
	opts directory_v1.Options

	routerChain router.Chain

	invokers []protocol.Invoker

	registry registry.InterfaceRegistry

	consumerURL *common.URL
}

func newDirectory(opts ...directory_v1.Option) directory_v1.Directory {
	dir := &interfaceDirectory{}
	// range opts

	go dir.watch()
	go dir.refresh()
	return dir
}

func (dir *interfaceDirectory) List(invocation protocol.Invocation) []protocol.Invoker {
	return dir.routerChain.Route(dir.consumerURL, invocation)
}

func (dir *interfaceDirectory) watch() {
	//w, err := dir.registry.Watch()
	//go func() {
	//	for {
	//		res, err := w.Next()
	//		dir.process(res)
	//	}
	//}()
}

func (dir *interfaceDirectory) process(result *registry.InterfaceResult) {
	// 更新invoker列表
}

func (dir *interfaceDirectory) refresh() {
	// 定时刷新列表
}

func (dir *interfaceDirectory) GetURL() *common.URL {
	//TODO implement me
	panic("implement me")
}

func (dir *interfaceDirectory) IsAvailable() bool {
	//TODO implement me
	panic("implement me")
}

func (dir *interfaceDirectory) Destroy() {
	//TODO implement me
	panic("implement me")
}
