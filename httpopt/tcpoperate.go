package httpopt

import (
	"context"
	"net"
	"time"

	"monitor_kong_service/config"

	"github.com/astaxie/beego/logs"
)

type tcpResult struct {
	conn net.Conn
	err  error
}

func GetProcessTcp(url string) (tcpresult int) {
	timeout_value := config.AppConfig.HealthCheckTimeoutTcp
	tcpresult = 0
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout_value)*time.Second)
	defer cancel()
	c := make(chan tcpResult, 3)

	go func() {
		conn, err := net.Dial("tcp", url)
		if err != nil {
			tcpresult = 1
		}
		pack := tcpResult{conn: conn, err: err}
		c <- pack
	}()
	select {
	case <-ctx.Done():
		tcpresult = 1
		logs.Error("TCP url: %s Timeout!\n", url)
		return
	case res := <-c:
		if res.conn != nil {
			defer res.conn.Close()
		}
		if res.err == nil {
			logs.Debug("TCP url: %s success.", url)
		} else {
			tcpresult = 1
			logs.Error("TCP url: %s failed.", url)
		}

		return
	}
}
