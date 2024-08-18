package main

import (
	"os"

	"super_calendar/initialize"
	"super_calendar/util"
)

func main() {
	initialize.Database(".env")

	util.InitializePayloadValidator()

	app := initialize.App()

	port := os.Getenv("PORT")
	app.Listen(":" + port)
}
