package casbin

import (
	"github.com/Aoang/haiju/pkg/database"
	"github.com/Aoang/haiju/pkg/logger"
	"github.com/casbin/casbin/v2"
	apter "github.com/casbin/gorm-adapter/v2"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

var enforcer *casbin.Enforcer

var log = logger.NewLogger(os.Stdout)

func Open() {
	a, err := apter.NewAdapterByDB(database.Use())
	if err != nil {
		log.Fatal(err)
	}
	enforcer, err = casbin.NewEnforcer("config/rbac_model.conf", a)
	if err != nil {
		log.Fatal(err)
	}
}

func Use() *casbin.Enforcer {
	return enforcer
}
