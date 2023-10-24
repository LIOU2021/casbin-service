package core

import (
	"casbin-service/config"
	"fmt"
	"log"
	"sync"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
)

var (
	casbinOnce    sync.Once
	Enforcer      *casbin.Enforcer
	CasbinAdapter *gormadapter.Adapter
)

func CasbinInit() {
	casbinOnce.Do(func() {
		source := fmt.Sprintf("%s:%s@tcp(%s:%s)/", config.Config.Mysql.User, config.Config.Mysql.Password, config.Config.Mysql.Host, config.Config.Mysql.Port)
		CasbinAdapter, err := gormadapter.NewAdapter("mysql", source)
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
