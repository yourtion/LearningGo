package conf

const DriverName = "mysql"

type DbConf struct {
	Host   string
	Port   int
	User   string
	Pwd    string
	DbName string
}

var MasterDbConfig = DbConf{
	Host:   "127.0.0.1",
	Port:   3306,
	User:   "root",
	Pwd:    "123456",
	DbName: "demo",
}

var SlaveDbConfig = DbConf{
	Host:   "127.0.0.1",
	Port:   3306,
	User:   "root",
	Pwd:    "123456",
	DbName: "demo",
}
