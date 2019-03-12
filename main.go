package main

import (
	"flag"

	"github.com/golang/glog"
)

func main() {
	defer glog.Flush()
	var mode string
	modeTip := "start mode :\n  l: list all stock names\n  r: real time data \n  h: history data"
	flag.StringVar(&mode, "mode", "", modeTip)
	flag.Parse()
	// flag.Usage()
}
