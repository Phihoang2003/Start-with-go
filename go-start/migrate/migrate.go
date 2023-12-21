package main

import (
	"go/api/go-start/infrastructure"
	"go/api/go-start/models"
)

func init() {
	infrastructure.ConnectToDB()
}

func main() {
	infrastructure.DB.AutoMigrate(&models.Product{})
}
