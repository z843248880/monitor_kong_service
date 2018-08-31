package server

import (
	"fmt"

	"monitor_kong_service/config"
	"monitor_kong_service/initp"
	"monitor_kong_service/kongOn"
	"strings"
)

func ServerRun() {
	for _, addr := range strings.Split(config.AppConfig.KongAddr, ",") {
		if len(addr) != 0 {
			allUpstreamList := []string{}
			url := "http://" + addr
			if config.AppConfig.CheckMethod != "tcp" {
				allUpstreamList = kongOn.GetAllUpstreamData(url)
				kongOn.CheckUpstreamHealth(url, allUpstreamList)
			} else {
				initp.InitDb(config.AppConfig.Phost, config.AppConfig.Puser, config.AppConfig.Ppassword, config.AppConfig.Pdbname, config.AppConfig.Pport)
				allUpstreamList = kongOn.GetAllUpstreamDataTcp(url)
				fmt.Println("allUpstreamList:", len(allUpstreamList), allUpstreamList)
				kongOn.CheckUpstreamHealthTcp(url, allUpstreamList)
			}

			// fmt.Println("allUpstreamList:", len(allUpstreamList), allUpstreamList)
		}
	}
}
