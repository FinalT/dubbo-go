package interface_registry

import (
	"dubbo.apache.org/dubbo-go/v3/cluster/router"
	"dubbo.apache.org/dubbo-go/v3/common"
	"dubbo.apache.org/dubbo-go/v3/directory_v1"
	"dubbo.apache.org/dubbo-go/v3/protocol"
	"dubbo.apache.org/dubbo-go/v3/registry_v1"
)

// router is the default router.
type instanceDirectory struct {
	opts directory_v1.Options

	routerChain router.Chain

	invokers []protocol.Invoker

	registry registry.Registry

	consumerURL *common.URL
}

func newDirectory(opts ...directory_v1.Option) directory_v1.Directory {
	dir := &instanceDirectory{}

	go dir.watch()
	go dir.refresh()
	return dir
}

func (dir *instanceDirectory) List(invocation protocol.Invocation) []protocol.Invoker {
	return dir.routerChain.Route(dir.consumerURL, invocation)
}

func (dir *instanceDirectory) watch() {
	//w, err := dir.registry.Watch()
	//go func() {
	//	for {
	//		res, err := w.Next()
	//		dir.process(res)
	//	}
	//}()
}

func (dir *instanceDirectory) process(result *registry.InterfaceResult) {
	// 更新invoker列表
}

func (dir *instanceDirectory) refresh() {
	// 定时去刷新
}

func (dir *instanceDirectory) GetURL() *common.URL {
	//TODO implement me
	panic("implement me")
}

func (dir *instanceDirectory) IsAvailable() bool {
	//TODO implement me
	panic("implement me")
}

func (dir *instanceDirectory) Destroy() {
	//TODO implement me
	panic("implement me")
}
