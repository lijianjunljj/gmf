package config

import (
	"fmt"
	"gmf/src/common/config/parser"
)

type Config struct {
	Mysql  *MysqlOptions
	Etcd   *EtcdOptions
	Web    *WebOptions
	parser parser.Parser
}

func NewConfig(parser parser.Parser) *Config {
	fmt.Println("parser", parser)
	config := &Config{
		parser: parser,
	}
	return config
}
func (c *Config) InitEtcd() {
	registryAddr := c.parser.GetString("etcd", "registryAddr")
	fmt.Println("registryAddr", registryAddr)
	etcdConf := NewEtcdOptions(RegistryAddr(registryAddr))
	fmt.Println("etcdConf:", etcdConf)
	c.Etcd = &etcdConf
}

func (c *Config) InitWeb() {
	protocol := c.parser.GetString("web", "protocol")
	addr := c.parser.GetString("web", "addr")
	fmt.Println("protocol", protocol)
	fmt.Println("addr", addr)
	webConf := NewWebOptions(Protocol(protocol), Addr(addr))
	fmt.Println("webConf:", webConf)
	c.Web = &webConf
}

func (c *Config) InitMysql() {
	fmt.Println("c.parser", c.parser)
	db := c.parser.GetString("mysql", "Db")
	fmt.Println("db", db)
	dbHost := c.parser.GetString("mysql", "DbHost")
	dbPort := c.parser.GetString("mysql", "DbPort")
	dbName := c.parser.GetString("mysql", "DbName")
	dbUser := c.parser.GetString("mysql", "DbUser")
	dbPassWord := c.parser.GetString("mysql", "DbPassWord")
	mysqlConf := NewMysqlOptions(Db(db), DbHost(dbHost), DbPort(dbPort),
		DbUser(dbUser), DbPassWord(dbPassWord), DbName(dbName))
	fmt.Println("mysqlConf:", mysqlConf)
	c.Mysql = &mysqlConf
}
