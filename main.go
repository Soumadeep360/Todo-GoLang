package main

import (
	"log"
	"todo-app/database"
	"todo-app/routes"
)

func main() {
	database.ConnectDB()
	r := routes.SetupRouter()

	log.Println("Server running on port 8080")
	r.Run(":8080")
}
