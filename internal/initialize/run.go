package initialize

import (
	"fmt"
	"go-crm-api-shopdev/global"
)

func Run() {
	LoadConfig()
	fmt.Println("Loading config", global.Config.Mysql.Username)
	InitLogger()
	InitMySql()
	InitRedis()

	r := InitRouter()
	r.Run(":8002")
}
