package main

import (
	"github.com/NiuStar/NiuServer/STNet"
	"github.com/casbin/casbin"
	"github.com/gin-contrib/authz"
	_ "github.com/go-sql-driver/mysql"
	"NiuServerTemplate/Controller"
)

func main() {
	//a, _ := xormadapter.NewAdapter("mysql", "root:123456@tcp(127.0.0.1:3306)/")
	// load the casbin model and policy from files, database is also supported.
	//e, err := casbin.NewEnforcer("./Config/auth/authz_model.conf", a)
	e, err := casbin.NewEnforcer("./Config/auth/authz_model.conf", "./Config/auth/authz_policy.csv")

	if err != nil {
		panic(err)
	}
	// Load the policy from DB.
	e.LoadPolicy()

	Controller.Initiate()

	server := STNet.DefaultServer()
	router := server.GetEngine()
	router.Use(authz.NewAuthorizer(e))
	server.Run("test", ":8002")
}
