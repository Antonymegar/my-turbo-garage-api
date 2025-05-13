package main

import (
	"log"
	"myturbogarage/config"
	"myturbogarage/helpers"
	"myturbogarage/routes"
	"time"

	"github.com/getsentry/sentry-go"
)

func init() {
	config.LoadEnvVariables()
	config.ConnectDB()
}

func main() {
	port := helpers.GetEnv("PORT", "8080")
	log.Println("Server running on port: " + port)
	server := routes.RouteApp()
	if err := server.Run(":" + port); err != nil {
		sentry.CaptureMessage(err.Error())
		panic(err)
	}
	// Flush buffered events before the program terminates.
	sentry.Flush(2 * time.Second)

}
