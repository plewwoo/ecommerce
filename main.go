package main

import (
	"ecommerce/common"
	"ecommerce/config"
	"ecommerce/router"
	"log"
)

func main() {
	common.SetThailandTimezone()

	common.InitEnv()
	common.ConnectDatabaseViper()

	config.AutoMigrate(common.Database)

	routes := router.SetupRouter()

	if err := routes.Listen(":8000"); err != nil {
		log.Fatal(err)
	}
}
