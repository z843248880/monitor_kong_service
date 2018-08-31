package kongOn

import (
	"monitor_kong_service/initp"
	"time"

	"github.com/astaxie/beego/logs"
)

//此时发现一个bug，get /upstreams/获取负载名字时，最多返回一百条记录
//https://github.com/Kong/kong/issues/3736
//因为weops-kong上已经超过100个负载节点了，所以选择直接从postresql里取负载的名称列表.

// type AllUpstreamDataTcp struct {
// 	Data []struct {
// 		CreatedAt int    `json:"created_at"`
// 		ID        string `json:"id"`
// 		Name      string `json:"name"`
// 		Orderlist []int  `json:"orderlist"`
// 		Slots     int    `json:"slots"`
// 	} `json:"data"`
// 	Next   string `json:"next"`
// 	Offset string `json:"offset"`
// 	Total  int    `json:"total"`
// }

// func GetAllUpstreamDataTcp(url string) []string {
// 	allUpstreamList := []string{}
// 	url_with_upstrem := url + "/upstreams/"
// 	out := httpopt.GetProcess("GET", url_with_upstrem, nil)
// 	allUpstreamDataTcp := &AllUpstreamDataTcp{}
// 	err := json.Unmarshal([]byte(out), allUpstreamDataTcp)
// 	if err != nil {
// 		logs.Error("GetAllUpstreamDataTcp json2struct error")
// 	}
// 	for _, upstreamdata := range allUpstreamDataTcp.Data {
// 		allUpstreamList = append(allUpstreamList, upstreamdata.Name)
// 	}
// 	return allUpstreamList
// }

type UpstreamsData struct {
	name       string
	slots      int
	orderlist  string
	created_at time.Time
}

func GetAllUpstreamDataTcp(url string) []string {
	var upstreamsDataList []UpstreamsData
	// err := initp.Db.QueryRow("select name from upstreams").Scan(&upstreamsData.name)
	statement := "select name from upstreams"
	rows, err := initp.Db.Query(statement)
	if err != nil {
		logs.Error("GetAllUpstreamDataTcp get upstreams from db err: %v", err)
	}

	for rows.Next() {
		var name string
		err = rows.Scan(&name)
		if err != nil {
			logs.Error("get upstreams each name err: %s", err)
		}
		var upstreamsData UpstreamsData
		upstreamsData.name = name

		upstreamsDataList = append(upstreamsDataList, upstreamsData)
	}
	var upstreamList []string
	for _, v := range upstreamsDataList {
		upstreamList = append(upstreamList, v.name)
	}
	return upstreamList
}
