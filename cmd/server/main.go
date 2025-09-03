package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/yolanda-test/go-web-service/internal/handlers"
	"github.com/yolanda-test/go-web-service/internal/models"
)

func main() {
	store := models.NewInMemoryStore()
	handler := handlers.NewHandler(store)

	router := gin.Default()
	handler.SetupRoutes(router)

	log.Println("Server starting on :8080...")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Could not start server: %v\n", err)
	}
}
