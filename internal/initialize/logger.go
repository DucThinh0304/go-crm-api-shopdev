package initialize

import (
	"go-crm-api-shopdev/global"
	"go-crm-api-shopdev/pkg/logger"
)

func InitLogger() {
	global.Logger = logger.NewLogger(global.Config.Logger)
}
