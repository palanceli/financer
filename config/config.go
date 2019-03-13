package config

import (
	"io/ioutil"
	"sync/atomic"

	"github.com/golang/glog"
	"gopkg.in/yaml.v2"
)

var globalConfig atomic.Value

// Initialize 从指定文件读取配置
func Initialize(confPath string, config interface{}) interface{} {
	data, err := ioutil.ReadFile(confPath)
	if err != nil {
		glog.Fatalf("FAILED to read config file : %s, err=%v", confPath, err)
	}
	err = yaml.Unmarshal(data, config)
	if err != nil {
		glog.Fatalf("FAILED to parse config file : %s, err=%v", confPath, err)
	}
	globalConfig.Store(config)
	return globalConfig.Load()
}

// Get 返回线程安全的config实例
func Get() interface{} {
	return globalConfig.Load()
}
