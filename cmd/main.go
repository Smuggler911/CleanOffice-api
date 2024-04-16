package main

import (
	"CleanOffice/config"
	"CleanOffice/internal/api"
	"CleanOffice/internal/repository"
	"CleanOffice/server"
	"log"
)

func init() {
	repository.ConnectDB()
}

func main() {

	conf, _ := config.LoadConfig()
	handlers := new(api.Handler)
	srv := new(server.Server)
	if err := srv.Run(conf.Port, handlers.InitRoutes()); err != nil {
		log.Fatalln("Error start server: " + err.Error())
	}

}
