package initialize

import "go.uber.org/zap"

func InitLogger() {
	logger, _ := zap.NewDevelopment()
	//设置全局logger
	zap.ReplaceGlobals(logger)
}
