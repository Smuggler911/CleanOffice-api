package main

import (
	"CleanOffice/internal/models"
	"CleanOffice/internal/repository"
)

func init() {
	repository.ConnectDB()
}
func main() {
	err := repository.DB.AutoMigrate(
		&models.User{}, &models.Offer{}, &models.Application{}, &models.Review{}, &models.CostRange{})
	if err != nil {
		return
	}
}
