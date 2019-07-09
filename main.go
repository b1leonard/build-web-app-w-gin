package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {

	router = gin.Default()

	router.LoadHTMLGlob("templates/*")

	initializeRoutes()

	router.Run(":80")

}

func port() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8082"
	}
	return ":" + port
}
