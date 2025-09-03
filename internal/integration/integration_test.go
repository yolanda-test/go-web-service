package integration

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/yolanda-test/go-web-service/internal/handlers"
	"github.com/yolanda-test/go-web-service/internal/models"
)

func setupRouter() *gin.Engine {
	store := models.NewInMemoryStore()
	handler := handlers.NewHandler(store)

	router := gin.Default()
	handler.SetupRoutes(router)
	return router
}

func TestIntegration(t *testing.T) {
	router := setupRouter()

	// Example of an integration test
	t.Run("GET /example", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/example", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("Expected status OK; got %v", w.Code)
		}
	})
}
