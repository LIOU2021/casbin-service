package config

import (
	"casbin-service/logger"
	"os"
	"sync"

	"gopkg.in/yaml.v3"
)

type ConfigYml struct {
	Logger LoggerYml `yaml:"logger"`
	Redis  RedisYml  `yaml:"redis"`
	Mysql  MysqlYml  `yaml:"mysql"`
}

type LoggerYml struct {
	MaxSize    int  `yaml:"maxSize"`    // 单一档案最大几M
	MaxBackups int  `yaml:"maxBackups"` // 最多保留几份
	MaxAge     int  `yaml:"maxAge"`     // 最多保留几天
	Compress   bool `yaml:"compress"`   // 压缩成gz
}

type RedisYml struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
}

type MysqlYml struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	DBname   string `yaml:"dbname"`
}

var Config *ConfigYml
var ConfigOnce sync.Once

func Init() {
	ConfigOnce.Do(func() {
		data, err := os.ReadFile("../config.yml")
		if err != nil {
			logger.Errorf("config init read file fail | err: %v", err)
			os.Exit(1)
		}

		err = yaml.Unmarshal(data, Config)
		if err != nil {
			logger.Errorf("config init unmarshal fail | err: %v", err)
			os.Exit(1)
		}
	})
}
