package main

import (
	"github.com/b-hurskyi/go-back-learn/db"
	"github.com/b-hurskyi/go-back-learn/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":5555")
}
