package global

import (
	"gin-digital-garden/config"

	"go.uber.org/zap"

	"github.com/spf13/viper"
)

var (
	GLOBAL_CONFIG *config.Configuration
	GLOBAL_VP     *viper.Viper
	GLOBAL_LOG    *zap.Logger
)
