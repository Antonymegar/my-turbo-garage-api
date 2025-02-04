package main

import (
  "myturbogarage/models"
  "myturbogarage/config"
)

func init() {
	config.LoadEnvVariables()
	config.ConnectDB()
}

func main() {
	config.DB.AutoMigrate(&models.User{})
}