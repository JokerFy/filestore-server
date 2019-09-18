package conf

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"sync"
)

// mysql配置
type MysqlConf struct {
	DSN string `yaml:"dsn"`
}

// app配置
type AppConf struct {
	Addr    string    `yaml:"addr"`
	AppName string    `yaml:"appname"`
	Mysql   MysqlConf `yaml:"mysql"`
}

var _conf *AppConf = nil
var _lock = sync.RWMutex{}

func LoadConf(ConfPath string) error {
	_lock.RLock()
	if _conf != nil {
		_lock.RUnlock()
		return nil
	}
	_lock.RUnlock()

	_lock.Lock()
	defer _lock.Unlock()
	_conf = &AppConf{}
	data, err := ioutil.ReadFile(ConfPath)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(data, &_conf)
	return err
}

func GetConf() *AppConf {
	return _conf
}
