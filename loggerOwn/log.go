package loggerOwn

import (
	"encoding/json"
	"fmt"
	"monitor_kong_service/config"

	"github.com/astaxie/beego/logs"
)

func convertLogLevel(level string) int {
	switch level {
	case "debug":
		return logs.LevelDebug
	case "warn":
		return logs.LevelWarn
	case "info":
		return logs.LevelInfo
	case "trace":
		return logs.LevelTrace
	case "error":
		return logs.LevelError
	}
	return logs.LevelDebug
}

func InitLogger() (err error) {
	configinner := make(map[string]interface{})
	configinner["filename"] = config.AppConfig.LogPath
	configinner["level"] = convertLogLevel(config.AppConfig.LogLevel)

	configStr, err := json.Marshal(configinner)
	if err != nil {
		fmt.Println("init logger failed, err:", err)
		return
	}

	logs.SetLogger(logs.AdapterFile, string(configStr))
	return
}
