package global

import (
	"gin-digital-garden/config"

	"go.uber.org/zap"
	"gorm.io/gorm"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

var (
	GLOBAL_CONFIG *config.Configuration
	GLOBAL_VP     *viper.Viper
	GLOBAL_LOG    *zap.Logger
	GLOBAL_DB     *gorm.DB
	GLOBAL_REDIS  *redis.Client
)
