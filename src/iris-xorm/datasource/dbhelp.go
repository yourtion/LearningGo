package datasource

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"iris-xorm/conf"
	"log"
	"sync"
)

var (
	masterEngine *xorm.Engine
	slaveEngine  *xorm.Engine
	lock         sync.Mutex
)

func InstanceMaster() *xorm.Engine {
	if masterEngine != nil {
		return masterEngine
	}

	lock.Lock()
	defer lock.Unlock()

	// 防止一开始并发情况下没拿到锁的进程获取不到链接导致重复创建
	if masterEngine != nil {
		return masterEngine
	}

	c := conf.MasterDbConfig
	driveSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", c.User, c.Pwd, c.Host, c.Port, c.DbName)

	engine, err := xorm.NewEngine(conf.DriverName, driveSource)
	if err != nil {
		log.Fatal("dbhelper.DbInstanceMaster,", err)
		return nil
	}

	//engine.ShowSQL(false)
	masterEngine = engine
	return engine

}

func InstanceSlave() *xorm.Engine {
	if slaveEngine != nil {
		return slaveEngine
	}

	lock.Lock()
	defer lock.Unlock()

	if slaveEngine != nil {
		return slaveEngine
	}

	c := conf.SlaveDbConfig
	driveSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", c.User, c.Pwd, c.Host, c.Port, c.DbName)

	engine, err := xorm.NewEngine(conf.DriverName, driveSource)
	if err != nil {
		log.Fatal("dbhelper.DbInstanceSlave,", err)
		return nil
	}

	//engine.ShowSQL(false)
	slaveEngine = engine
	return engine
}
