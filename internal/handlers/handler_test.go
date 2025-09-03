package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/yolanda-test/shorty/internal/models"
)

func TestHandler_GetShortenedURL(t *testing.T) {
	store := models.NewInMemoryStore()
	handler := NewHandler(store)

	router := gin.Default()
	handler.SetupRoutes(router)

	t.Run("valid request", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/shortened-url/123", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		// Add more assertions based on expected response
	})

	t.Run("invalid request", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/shortened-url/invalid", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotFound, w.Code)
		// Add more assertions based on expected response
	})
}
