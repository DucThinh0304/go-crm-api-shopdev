package initialize

import (
	"fmt"
	"go-crm-api-shopdev/global"

	"go.uber.org/zap"
)

func Run() {
	LoadConfig()
	fmt.Println("Loading config", global.Config.Mysql.Username)
	InitLogger()
	global.Logger.Info("Config log", zap.String("ok", "abc"))
	InitMySql()
	InitRedis()

	r := InitRouter()
	r.Run(":8002")
}
