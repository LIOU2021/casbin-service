package core

import (
	"casbin-service/config"
	"fmt"
	"log"
	"sync"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	casbinOnce    sync.Once
	Enforcer      *casbin.Enforcer
	CasbinAdapter *gormadapter.Adapter
)

func CasbinInit() {
	casbinOnce.Do(func() {
		source := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.Config.Mysql.User, config.Config.Mysql.Password, config.Config.Mysql.Host, config.Config.Mysql.Port, config.Config.Mysql.DBname)

		// 預設會跑gorm的auto migration
		// CasbinAdapter, err := gormadapter.NewAdapter("mysql", source)

		db, err := gorm.Open(mysql.Open(source), &gorm.Config{})
		if err != nil {
			log.Fatalf("casbin init gorm open failed | err: %v\n", err)
		}
		gormadapter.TurnOffAutoMigrate(db)
		CasbinAdapter, err := gormadapter.NewAdapterByDBWithCustomTable(db, nil, "casbin_rule")
		if err != nil {
			log.Fatalf("casbin init NewAdapter failed | err: %v\n", err)
		}

		Enforcer, err = casbin.NewEnforcer("./model.conf", CasbinAdapter)
		if err != nil {
			log.Fatalf("casbin init NewEnforcer failed | err: %v\n", err)
		}

		Enforcer.LoadPolicy()
	})
}

// reload data from file/database
func CasbinReload() {
	Enforcer.LoadPolicy()
}
