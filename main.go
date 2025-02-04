package main

import (
 "github.com/gin-gonic/gin"
 "myturbogarage/config"
)

func init() {
 config.LoadEnvVariables()
 config.ConnectDB()
}

func main() {

 r := gin.Default()

 r.Run()
}