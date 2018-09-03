package initp

import (
	"database/sql"
	"fmt"
	"monitor_kong_service/config"
	"monitor_kong_service/loggerOwn"

	"github.com/astaxie/beego/logs"
	_ "github.com/lib/pq"
)

var Db *sql.DB

func InitDb(host, user, password, dbname string, port int) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.AppConfig.Phost, config.AppConfig.Pport, config.AppConfig.Puser, config.AppConfig.Ppassword, config.AppConfig.Pdbname)
	var err error
	Db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
}

//func ReturnDb() *sql.DB {
//	return Db
//}

func init() {
	filename := "./conf/ch.conf"
	err := config.LoadConf("ini", filename)
	if err != nil {
		fmt.Printf("load conf failed, err: %v\n", err)
		panic("load conf failed")
	}

	err = loggerOwn.InitLogger()
	if err != nil {
		fmt.Printf("init logger failed, err: %v\n", err)
		panic("init logger failed")
	}

	logs.Debug("initialize success")
}
