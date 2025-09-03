package handlers

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

type Handler struct {
    store Store
}

type Store interface {
    // Define methods that the store should implement
}

func NewHandler(store Store) *Handler {
    return &Handler{store: store}
}

func (h *Handler) SetupRoutes(router *gin.Engine) {
    router.GET("/example", h.ExampleHandler)
}

func (h *Handler) ExampleHandler(c *gin.Context) {
    // Example handler logic
    c.JSON(http.StatusOK, gin.H{"message": "Hello, World!"})
}