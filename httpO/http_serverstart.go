package httpO

import (
	"monitor_kong_service/config"
	"net/http"
	"strconv"

	"github.com/astaxie/beego/logs"
)

func HttpServerStart() {
	http.HandleFunc("/wehealthcheck", weHealthCheck)
	server_addr := config.AppConfig.ListenAddr + ":" + strconv.Itoa(config.AppConfig.ListenPort)
	err := http.ListenAndServe(server_addr, nil)
	if err != nil {
		logs.Error("start http server failed, err: %s\n", err)
		panic("start http server failed\n")
	}
}
