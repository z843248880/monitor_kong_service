package config

import (
	"fmt"

	"github.com/astaxie/beego/config"
)

var (
	AppConfig *Config
)

type Config struct {
	ListenAddr             string
	ListenPort             int
	HealthCheckTimeoutTcp  int
	HealthCheckTimeoutHttp int
	PollNum                int
	AlarmServer            string
	CheckInternal          int
	LogLevel               string
	LogPath                string
	PhoneNumber            string
	KongAddr               string
	CheckMethod            string
	Phost                  string
	Pport                  int
	Puser                  string
	Ppassword              string
	Pdbname                string
}

func LoadConf(confType, filename string) (err error) {
	conf, err := config.NewConfig(confType, filename)
	if err != nil {
		fmt.Println("new config failed, err:", err)
		return
	}

	AppConfig = &Config{}

	AppConfig.ListenAddr = conf.String("server::listen_addr")
	if len(AppConfig.ListenAddr) == 0 {
		AppConfig.ListenAddr = "0.0.0.0"
	}

	AppConfig.ListenPort, err = conf.Int("server::listen_port")
	if err != nil {
		AppConfig.ListenPort = 21024
	}

	AppConfig.HealthCheckTimeoutTcp, err = conf.Int("server::healthcheck_timeout_tcp")
	if err != nil {
		AppConfig.HealthCheckTimeoutTcp = 7
	}

	AppConfig.HealthCheckTimeoutHttp, err = conf.Int("server::healthcheck_timeout_http")
	if err != nil {
		AppConfig.HealthCheckTimeoutHttp = 7
	}

	AppConfig.PollNum, err = conf.Int("server::poll_num")
	if err != nil {
		AppConfig.PollNum = 8
	}

	AppConfig.CheckInternal, err = conf.Int("server::check_internal")
	if err != nil {
		AppConfig.CheckInternal = 300
	}

	AppConfig.AlarmServer = conf.String("server::alarm_server")
	if len(AppConfig.AlarmServer) == 0 {
		err = fmt.Errorf("invalid alarm server")
	}

	AppConfig.LogLevel = conf.String("logs::log_level")
	if len(AppConfig.LogLevel) == 0 {
		AppConfig.LogLevel = "debug"
	}
	AppConfig.LogPath = conf.String("logs::log_path")
	if len(AppConfig.LogPath) == 0 {
		AppConfig.LogPath = "./logs"
	}

	AppConfig.PhoneNumber = conf.String("user::phone_number")
	if len(AppConfig.PhoneNumber) == 0 {
		err = fmt.Errorf("invalid phone number")
	}

	AppConfig.KongAddr = conf.String("kong::server_addr")
	if len(AppConfig.KongAddr) == 0 {
		err = fmt.Errorf("invalid kong server")
	}

	AppConfig.CheckMethod = conf.String("kong::health_check_method")
	if len(AppConfig.CheckMethod) == 0 {
		AppConfig.CheckMethod = "http"
	}

	AppConfig.Phost = conf.String("postgresql::phost")
	if AppConfig.CheckMethod == "tcp" {
		if len(AppConfig.Phost) == 0 {
			err = fmt.Errorf("invalid postgresql host")
		}
	}

	AppConfig.Pport, err = conf.Int("postgresql::pport")
	if err != nil {
		if AppConfig.CheckMethod == "tcp" {
			err = fmt.Errorf("invalid postgresql pport")
		}
	}

	AppConfig.Puser = conf.String("postgresql::puser")
	if AppConfig.CheckMethod == "tcp" {
		if len(AppConfig.Puser) == 0 {
			err = fmt.Errorf("invalid postgresql puser")
		}
	}

	AppConfig.Ppassword = conf.String("postgresql::ppassword")
	if AppConfig.CheckMethod == "tcp" {
		if len(AppConfig.Ppassword) == 0 {
			err = fmt.Errorf("invalid postgresql ppassword")
		}
	}

	AppConfig.Pdbname = conf.String("postgresql::pdbname")
	if AppConfig.CheckMethod == "tcp" {
		if len(AppConfig.Pdbname) == 0 {
			err = fmt.Errorf("invalid postgresql pdbname")
		}
	}

	return
}
