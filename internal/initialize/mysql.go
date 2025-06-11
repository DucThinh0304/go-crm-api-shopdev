package initialize

import (
	"fmt"
	"go-crm-api-shopdev/global"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func checkErrorPanic(err error, errString string) {
	if err != nil {
		global.Logger.Error(errString, zap.Error(err))
		panic(err)
	}
}

func InitMySql() {
	config := global.Config.Mysql
	dsn := "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	s := fmt.Sprintf(dsn, config.Username, config.Password, config.Host, config.Port, config.Dbname)
	db, err := gorm.Open(mysql.Open(s), &gorm.Config{})
	checkErrorPanic(err, "Init SQL failed")
}
