package config

import (
	"fmt"
	"github.com/jinzhu/configor"
	"os"
)

var Conf *Config

type Config struct {
	Etcd  Etcd
	MySql MySql
	Log   Log
}

type Etcd struct {
	Hosts    string
	Username string
	Password string
}

type MySql struct {
	Url      string
	User     string
	Password string
	Port     uint
}

type Log struct {
	Level string
}

func init() {
	fmt.Fprintln(os.Stdout, "init config, use config.yml")
	configor.Load(&Conf, "config.yml")
}
