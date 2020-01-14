package main

import (
	"flag"
	"net"
	"slowlog_elk/common"
	"slowlog_elk/tomlConfig"
	"slowlog_elk/watcher"
)
var (
	config string
)

func init()  {
	flag.StringVar(&config, "c", "conf/sl.toml", "配置文件")
}

func main()  {
	flag.Parse()
	// 解析配置文件
	configration := tomlConfig.TomlConfig(config)
	//初始化log功能
	loggs := common.WriteLog(configration.System.LogFile)
	// 监听一个端口
	_, err := net.Listen("tcp", configration.System.EndPoint)
	if err != nil {
		panic("sorry bind address is bad")
	} else {
		loggs.Println("bind address is ok, endport is", configration.System.EndPoint)
	}
	// 走起
	watcher.Start(configration, loggs)
}
