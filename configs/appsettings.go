package configs

import (
	"sync"

	"github.com/spf13/viper"
	"gorm.io/gorm/logger"
)

var (
	logger_once sync.Once
	Logger      logger.Interface
)

func GetLogger() logger.Interface {
	logger_once.Do(func() {
		Logger = logger.Default.LogMode(logger.Info)
	})
	return Logger
}

func SetConfigurations(fileName, fileType, path string) error {
	viper.SetConfigName("appsettings")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs/")
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	return nil
}

func GetString(key string) string {
	return viper.GetString(key)
}

func GetInt(key string) int {
	return viper.GetInt(key)
}
