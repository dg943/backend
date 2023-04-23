package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dg943/MyProject/backend/configs"
	"github.com/dg943/MyProject/backend/migrations"
	"github.com/dg943/MyProject/backend/servers"
	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	server *http.Server
)

func init() {
	var err error
	configs.SetConfigurations("appsettings", "yaml", "./configs/")
	db, err = configs.GetDB()
	if err != nil {
		log.Fatal(err)
	}
	migrations.Run(db)
	server = servers.NewServer(db)
}

func main() {
	fmt.Println("Starting the server ")
	log.Fatal(server.ListenAndServe())
}
