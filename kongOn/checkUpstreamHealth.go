package kongOn

import (
	"encoding/json"
	"fmt"
	"monitor_kong_service/alarm"
	"monitor_kong_service/config"
	"monitor_kong_service/httpopt"
	"strings"

	// "time"
	"github.com/astaxie/beego/logs"
)

type UpstreamStatus struct {
	Data []struct {
		CreatedAt  int    `json:"created_at"`
		Health     string `json:"health"`
		ID         string `json:"id"`
		Target     string `json:"target"`
		UpstreamID string `json:"upstream_id"`
		Weight     int    `json:"weight"`
	} `json:"data"`
	NodeID string `json:"node_id"`
	Total  int    `json:"total"`
}

func CheckUpstreamHealth(url string, allUpstreamList []string) {
	for _, uname := range allUpstreamList {
		checkUpstreamHealth(url, uname)
		// time.Sleep(time.Second * 1)
	}
	return
}

func checkUpstreamHealth(url, uname string) {
	upstreamCheckUrl := url + "/upstreams/" + uname + "/health"
	out := httpopt.GetProcess("GET", upstreamCheckUrl, nil)
	checkResult := &UpstreamStatus{}
	err := json.Unmarshal([]byte(out), checkResult)
	if err != nil {
		logs.Error("checkUpstreamHealth json2struct error")
	}
	for _, data := range checkResult.Data {
		if data.Health != "HEALTHY" {
			serviceName := strings.TrimLeft(uname, "-group")
			alarm_content := fmt.Sprintf("Service name: %s; target: %s; healthStatus: %s\n", serviceName, data.Target, data.Health)
			// logs.Debug(alarm_content)
			logs.Error(alarm_content)
			logs.Error("alarm_content:", config.AppConfig.PhoneNumber, alarm_content)
			alarm_result := alarm.SendSms(config.AppConfig.PhoneNumber, alarm_content)
			logs.Error("alarm_result: %s", alarm_result)
		}
	}
	return
}
