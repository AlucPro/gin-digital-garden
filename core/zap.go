package core

import (
	"fmt"
	"gin-digital-garden/core/internal"
	"gin-digital-garden/global"
	"gin-digital-garden/utils"

	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// InitializeZap Zap 获取 zap.Logger
func InitializeZap() (logger *zap.Logger) {
	if ok, _ := utils.PathExists(global.GLOBAL_CONFIG.Zap.Director); !ok { // 判断是否有Director文件夹
		fmt.Printf("create %v directory\n", global.GLOBAL_CONFIG.Zap.Director)
		_ = os.Mkdir(global.GLOBAL_CONFIG.Zap.Director, os.ModePerm)
	}

	cores := internal.Zap.GetZapCores()
	logger = zap.New(zapcore.NewTee(cores...))

	if global.GLOBAL_CONFIG.Zap.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}

	fmt.Println("====2-zap====: zap log init success")
	return logger
}
