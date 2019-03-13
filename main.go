package main

import (
	"financer/config"
	financerconfig "financer/financer_config"
	"financer/stockspider"
	"flag"

	"github.com/golang/glog"
)

func main() {
	defer glog.Flush()
	var confPath string
	flag.StringVar(&confPath, "c", "./default_config.yml", "configuration file path")
	flag.Parse()
	config := initConfig(confPath)

	switch config.StartMode {
	case "list":
		doList()
	default:
		glog.Fatalf("UNKNOWN start mode : %s", config.StartMode)
	}
}

func doList() {
	allStocks := stockspider.ListAllStocks()
	for _, s := range allStocks {
		glog.Infof("name:%s, symbol:%s", s.Name, s.Symbol)
	}
}

func initConfig(confPath string) *financerconfig.FinancerConfig {
	config.Initialize(confPath, &financerconfig.FinancerConfig{})
	return config.Get().(*financerconfig.FinancerConfig)
}
