package main

import (
	"go-mongo-lab/config"
	"go-mongo-lab/handlers"
	"go-mongo-lab/routes"
	"log"
)

func main() {
	db := config.ConnectDB()
	handler := handlers.NewUserHandler(db)
	router := routes.SetupRouter(handler)

	log.Println("Servidor rodando em http://localhost:8088")
	router.Run(":8088")
}
