package registry

import (
	"dubbo.apache.org/dubbo-go/v3/common"
)

type Result struct {
	Action    ActionType
	Instances []*Instance
}

type Watcher interface {
	Next() (*Result, error)
	Close()
}

type InterfaceResult struct {
	Action   ActionType
	Services []*common.URL
}

type InterfaceWatcher interface {
	Next() (*InterfaceResult, error)
	Close()
}

type ActionType int

const (
	// Create is emitted when a new service is registered.
	Create ActionType = iota
	// Delete is emitted when an existing service is deregsitered.
	Delete
	// Update is emitted when an existing servicec is updated.
	Update
)

func (t ActionType) String() string {
	switch t {
	case Create:
		return "create"
	case Delete:
		return "delete"
	case Update:
		return "update"
	default:
		return "unknown"
	}
}
