package main

import (
	"github.com/gin-gonic/gin"

	"example/src/config"
	"example/src/router"
)

func main() {
	r := gin.Default()

	config.ConnectMongoDB("mongodb://localhost:27017")
	// router.UserRouter(r)
	router.CustomerRoute(r)
	r.Run()
}
