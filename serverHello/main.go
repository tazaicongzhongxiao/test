package main

import (
	"Tarsgo_test/serverHello/handler"
	"fmt"
	"github.com/TarsCloud/TarsGo/tars"
	"test"
)

func main() {
	tars.ServerConfigPath = "./config/config.conf"
	cfg := tars.GetServerConfig()
	// 从配置文件解析配置

	// New servant imp
	imp := new(handler.Imp)
	// New servant
	// Register Servant
	fmt.Printf("cfg=%+v\n", cfg)
	new(test.Say).AddServantWithContext(imp, cfg.App+"."+cfg.Server+".SayHelloObj")

	// Run application
	tars.Run()
}
