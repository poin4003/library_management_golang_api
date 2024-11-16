package main

import (
	"github.com/gin-contrib/cors"
	"library_management/db"
	"library_management/routes"
)

func main() {
	db.InitPostgresql()

	r := routes.InitRouter()

	r.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:  []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders: []string{"Content-Length"},
	}))

	r.Run(":8080")
}
