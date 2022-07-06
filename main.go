package main

import (
	"altman/core"
	"altman/global"
	"altman/initialize"
	"altman/routers"
)

func main() {
	core.Viper()
	global.DB = initialize.InitDB()
	if global.DB != nil {
		//initialize.RegisterTables(global.DB)
		db, _ := global.DB.DB()
		defer db.Close()
	}
	routers.InitRouters()
}
