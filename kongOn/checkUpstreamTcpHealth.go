package kongOn

import (
	"encoding/json"
	"fmt"

	"monitor_kong_service/alarm"
	"monitor_kong_service/httpopt"

	// "strings"

	// "time"
	"monitor_kong_service/config"

	"github.com/astaxie/beego/logs"
)

type UpstreamStatusTcp struct {
	Data []struct {
		CreatedAt  int    `json:"created_at"`
		ID         string `json:"id"`
		Target     string `json:"target"`
		UpstreamID string `json:"upstream_id"`
		Weight     int    `json:"weight"`
	} `json:"data"`
	Total int `json:"total"`
}

func CheckUpstreamHealthTcp(url string, allUpstreamList []string) {
	// var needCheckList chan []string
	needCheckList := make(chan []string, 10000)
	getTargetChan := make(chan bool, 10000)
	workChan := make(chan bool, 10000)

	pollnum := config.AppConfig.PollNum
	tagnum := 0
	for _, uname := range allUpstreamList {
		go checkUpstreamHealthTcp(url, uname, needCheckList, getTargetChan)
		tagnum++
	}

	for i := 0; i < tagnum; i++ {
		_ = <-getTargetChan
	}
	close(needCheckList)

	for i := 0; i < pollnum; i++ {
		go checkTargetTcp(needCheckList, workChan)
	}

	for i := 0; i < pollnum; i++ {
		a := <-workChan
		logs.Debug("workChanDone:", a)
	}

	return
}

func checkTargetTcp(needCheckList chan []string, workChan chan bool) {
	for v := range needCheckList {
		tcpresult := httpopt.GetProcessTcp(v[1])
		if tcpresult != 0 {
			alarm_content := fmt.Sprintf("Service name: %s; target: %s; healthStatus: %s\n", v[0], v[1], "UNHEALTH.")
			alarm.SendSms(config.AppConfig.PhoneNumber, alarm_content)
			logs.Error(alarm_content)
		} else {
			logs.Debug("Service name: %s; target: %s; healthStatus: %s\n", v[0], v[1], "HEALTH.")
		}
	}
	workChan <- true
}

func checkUpstreamHealthTcp(url, uname string, needCheckList chan []string, getTargetChan chan bool) {
	upstreamGetTargetsUrl := url + "/upstreams/" + uname + "/targets/active/"
	out := httpopt.GetProcess("GET", upstreamGetTargetsUrl, nil)
	checkResult := &UpstreamStatusTcp{}
	err := json.Unmarshal([]byte(out), checkResult)
	if err != nil {
		logs.Error("checkUpstreamHealthTcp json2struct error")
	}
	targetList := make(map[string]bool)
	for _, data := range checkResult.Data {
		if _, ok := targetList[data.Target]; !ok {
			targetList[data.Target] = true
			msg := []string{uname, data.Target}
			needCheckList <- msg
		}
	}
	getTargetChan <- true
	return
}
