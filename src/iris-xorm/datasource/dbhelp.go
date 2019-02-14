package datasource

import (
	"github.com/go-xorm/xorm"
	"sync"
)

var (
	masterEngine *xorm.Engine
	slaveEngine  *xorm.Engine
	lock         sync.Mutex
)

func InstanceMaster() *xorm.Engine {
	// TODO:
	return masterEngine
}

func InstanceSlave() *xorm.Engine {
	// TODO:
	return slaveEngine
}
