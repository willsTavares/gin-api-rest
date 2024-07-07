package main

import (
	"github.com/willsTavares/api-go-gin/database"
	"github.com/willsTavares/api-go-gin/routes"
)

func main() {
	database.ConnectToDatabase()
	routes.HandleRequests()
}
