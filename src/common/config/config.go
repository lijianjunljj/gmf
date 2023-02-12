package config

import (
	"gmf/src/common/config/parser"
)

type Config struct {
	Mysql   *MysqlOptions
	Etcd    *EtcdOptions
	Web     *WebOptions
	Service *ServiceOptions
	parser  parser.Parser
	DbType  string
}

func NewConfig(parser parser.Parser) *Config {
	//fmt.Println("parser", parser)
	config := &Config{
		parser: parser,
	}
	return config
}
func (c *Config) InitEtcd() {
	registryAddr := c.parser.GetString("etcd", "registryAddr")
	//fmt.Println("registryAddr", registryAddr)
	etcdConf := NewEtcdOptions(RegistryAddr(registryAddr))
	//fmt.Println("etcdConf:", etcdConf)
	c.Etcd = &etcdConf
}
func (c *Config) InitService(serverName string) {
	serviceName := c.parser.GetString("servers", serverName, "serviceName")
	address := c.parser.GetString("servers", serverName, "address")
	//fmt.Println("serviceName", serviceName)
	serviceConf := NewServiceOptions(ServiceName(serviceName), Address(address))
	//fmt.Println("serviceConf:", serviceConf)
	c.Service = &serviceConf
}

func (c *Config) InitWeb() {
	protocol := c.parser.GetString("gateway", "protocol")
	addr := c.parser.GetString("gateway", "addr")
	//fmt.Println("protocol", protocol)
	//fmt.Println("addr", addr)
	webConf := NewWebOptions(Protocol(protocol), Addr(addr))
	//fmt.Println("webConf:", webConf)
	c.Web = &webConf
}
func (c *Config) InitDbType() {
	c.DbType = c.parser.GetString("db_type")
}
func (c *Config) InitMysql() {
	//fmt.Println("c.parser", c.parser)
	db := c.parser.GetString("mysql", "Db")
	//fmt.Println("db", db)
	dbHost := c.parser.GetString("mysql", "DbHost")
	dbPort := c.parser.GetString("mysql", "DbPort")
	dbName := c.parser.GetString("mysql", "DbName")
	dbUser := c.parser.GetString("mysql", "DbUser")
	dbPassWord := c.parser.GetString("mysql", "DbPassWord")
	mysqlConf := NewMysqlOptions(Db(db), DbHost(dbHost), DbPort(dbPort),
		DbUser(dbUser), DbPassWord(dbPassWord), DbName(dbName))
	//fmt.Println("mysqlConf:", mysqlConf)
	c.Mysql = &mysqlConf
}
