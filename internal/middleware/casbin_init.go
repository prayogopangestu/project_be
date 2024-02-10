package middleware

import (
	"fmt"
	"log"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type CasbinInit interface {
	InitEnforcer() string
}

type casbinConnection struct {
	connection *gorm.DB
}

func NewCasbinInit(db *gorm.DB) CasbinInit {
	return &casbinConnection{
		connection: db,
	}
}

func (db *casbinConnection) InitEnforcer() string {
	a, err := gormadapter.NewAdapterByDB(db.connection)
	if err != nil {
		fmt.Print("error at casbin")
	}
	e, _ := casbin.NewEnforcer("config/etc/casbin/model.conf", a)

	if err != nil {
		log.Fatalf("error: enforcer: %s", err)
	}

	// Frontend
	// e.AddPermissionForUser("1", "rawat_jalan", "store")
	// e.AddPermissionForUser("1", "rawat_jalan", "read")

	// e.AddPermissionForUser("2", "rawat_jalan", "read")
	// e.AddPermissionForUser("2", "rawat_jalan", "store")

	// e.AddPermissionForUser("3", "poliklinik", "read")
	// e.AddPermissionForUser("1", "poliklinik", "read")

	// e.AddPermissionForUser("4", "farmasi", "read")
	// e.AddPermissionForUser("1", "farmasi", "read")

	// //backend
	// e.AddRoleForUser("superadmin@gmail.com", "1")
	// e.AddRoleForUser("superadmin@gmail.com", "2")
	// e.AddRoleForUser("superadmin@gmail.com", "3")
	// e.AddRoleForUser("superadmin@gmail.com", "4")
	// e.AddRoleForUser("fataa@gmail.com", "2")
	// e.AddRoleForUser("nawaf@gmail.com", "2")
	// e.AddRoleForUser("poli@gmail.com", "3")
	// e.AddRoleForUser("farmasi@gmail.com", "4")

	e.SavePolicy()
	return "OK"

}
