package alarm

import (
	"monitor_kong_service/httpopt"
	"strings"
	"monitor_kong_service/config"
)

func SendSms(phone, smsContent string) (out []byte) {
	hbody := "tos=" + phone + "&content=" + smsContent
	out = httpopt.GetProcess("POST", config.AppConfig.AlarmServer, strings.NewReader(hbody))
	return
}
