package main

import (
	"myturbogarage/config"
	"myturbogarage/models"

	"github.com/getsentry/sentry-go"
)

func init() {
	config.LoadEnvVariables()
	config.ConnectDB()
}

func main() {
	if err := config.DB.AutoMigrate(models.GetSchema()...); err != nil {
		sentry.CaptureException(err)
		panic(err)
	}
}
