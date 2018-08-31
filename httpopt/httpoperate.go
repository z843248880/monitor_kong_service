package httpopt

import (
	"context"
	// "fmt"
	"io"
	"io/ioutil"
	"monitor_kong_service/config"
	"net/http"
	"time"

	"github.com/astaxie/beego/logs"
)

type Result struct {
	r   *http.Response
	err error
}

func GetProcess(method, url string, body io.Reader) (httpResult []byte) {
	timeout_value := config.AppConfig.HealthCheckTimeoutHttp   
	httpResult = []byte{}
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout_value) * time.Second)
	defer cancel()
	tr := &http.Transport{}
	client := &http.Client{Transport: tr}
	c := make(chan Result, 3)
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		logs.Error("http request failed, err:", err)
		return
	}

	//这行header必须有，有的短信机识别接收post请求时此header。
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	go func() {
		resp, err := client.Do(req)
		pack := Result{r: resp, err: err}
		c <- pack
	}()
	select {
	case <-ctx.Done():
		tr.CancelRequest(req)
		res := <-c
		logs.Error("HTTP Timeout! err:", res.err)

	case res := <-c:
		if res.r != nil {
			defer res.r.Body.Close()
			out, _ := ioutil.ReadAll(res.r.Body)
			logs.Debug("GET url: %s success.", url)
			httpResult = out
		}
	}
	return
}
