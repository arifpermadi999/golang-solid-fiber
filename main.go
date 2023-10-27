package main

import (
	"log"

	"github.com/arifpermadi999/core-echo-golang/db"
	"github.com/arifpermadi999/core-echo-golang/routes"
)

func main() {
	database := db.Init()
	db.Migration(database)
	app := routes.Init(database)
	log.Fatal(app.Listen(":3000"))
}
