package config

import (
	"casbin-service/helper"
	"errors"
	"log"
	"os"
	"sync"

	"gopkg.in/yaml.v3"
)

type ConfigYml struct {
	Server ServerYml `yaml:"server"`
	Logger LoggerYml `yaml:"logger"`
	Redis  RedisYml  `yaml:"redis"`
	Mysql  MysqlYml  `yaml:"mysql"`
}

type ServerYml struct {
	Port string `yaml:"port"`
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
		data, err := os.ReadFile("./config.yml")
		if err != nil {
			log.Fatalf("config init read file fail | err: %v\n", err)
		}
		Config = &ConfigYml{}
		err = yaml.Unmarshal(data, Config)
		if err != nil {
			log.Fatalf("config init unmarshal fail | err: %v\n", err)
		}

		if err := Check(); err != nil {
			log.Fatalf("config init check fail | err: %v\n", err)
		}
	})
}

// 檢查logger配置是否能正常讀取
func Check() (err error) {
	if helper.IsEmpty(Config.Server.Port) {
		return errors.New("Config.Server.Port was empty")
	}

	if helper.IsEmpty(Config.Logger.MaxSize) {
		return errors.New("Config.Logger.MaxSize was empty")
	}

	if helper.IsEmpty(Config.Logger.MaxBackups) {
		return errors.New("Config.Logger.MaxBackups was empty")
	}

	if helper.IsEmpty(Config.Logger.MaxAge) {
		return errors.New("Config.Logger.MaxAge was empty")
	}

	out, err := yaml.Marshal(Config)
	if err != nil {
		return
	}

	if helper.IsEmpty(string(out)) {
		return errors.New("config.yaml marshal was empty")
	}

	return
}
