package kongOn

import (
	"encoding/json"
	"monitor_kong_service/httpopt"

	"github.com/astaxie/beego/logs"
)

type AllUpstreamData struct {
	Data []struct {
		CreatedAt    int    `json:"created_at"`
		HashFallback string `json:"hash_fallback"`
		HashOn       string `json:"hash_on"`
		Healthchecks struct {
			Active struct {
				Concurrency int `json:"concurrency"`
				Healthy     struct {
					HTTPStatuses []int `json:"http_statuses"`
					Interval     int   `json:"interval"`
					Successes    int   `json:"successes"`
				} `json:"healthy"`
				HTTPPath  string `json:"http_path"`
				Timeout   int    `json:"timeout"`
				Unhealthy struct {
					HTTPFailures int   `json:"http_failures"`
					HTTPStatuses []int `json:"http_statuses"`
					Interval     int   `json:"interval"`
					TcpFailures  int   `json:"tcp_failures"`
					Timeouts     int   `json:"timeouts"`
				} `json:"unhealthy"`
			} `json:"active"`
			Passive struct {
				Healthy struct {
					HTTPStatuses []int `json:"http_statuses"`
					Successes    int   `json:"successes"`
				} `json:"healthy"`
				Unhealthy struct {
					HTTPFailures int   `json:"http_failures"`
					HTTPStatuses []int `json:"http_statuses"`
					TcpFailures  int   `json:"tcp_failures"`
					Timeouts     int   `json:"timeouts"`
				} `json:"unhealthy"`
			} `json:"passive"`
		} `json:"healthchecks"`
		ID    string `json:"id"`
		Name  string `json:"name"`
		Slots int    `json:"slots"`
	} `json:"data"`
	Total int `json:"total"`
}

func GetAllUpstreamData(url string) []string {
	allUpstreamList := []string{}
	url_with_upstrem := url + "/upstreams/"
	out := httpopt.GetProcess("GET", url_with_upstrem, nil)
	allUpstreamData := &AllUpstreamData{}
	err := json.Unmarshal([]byte(out), allUpstreamData)
	if err != nil {
		logs.Error("getAllUpstreamData json2struct error")
	}
	for _, upstreamdata := range allUpstreamData.Data {
		allUpstreamList = append(allUpstreamList, upstreamdata.Name)
	}
	return allUpstreamList
}
