package conf

import (
	"os"
	"path/filepath"
)

var Debug = true
var confPath = "./"

func init() {
	env := "dev"
	RootPath, _ := os.Getwd()
	confPath = RootPath
	tmp := filepath.Join(RootPath, "app."+env+".yaml")
	if _, err := os.Stat(tmp); err == nil {
		confPath = tmp
	}
}

func GetConfPath() string {
	return confPath
}
