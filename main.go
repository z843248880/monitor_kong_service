package main

import (
	"fmt"
	"monitor_kong_service/config"
	"monitor_kong_service/httpO"
	"monitor_kong_service/server"
	"time"
)

func main() {
	go httpO.HttpServerStart()
	for {
		server.ServerRun()
		fmt.Println(time.Now().Unix())
		time.Sleep(time.Duration(config.AppConfig.CheckInternal) * time.Second)
	}
}
