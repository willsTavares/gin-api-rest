package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/willsTavares/api-go-gin/controllers"
)

func SetupTestRoutes() *gin.Engine {
	routes := gin.Default()
	return routes
}

func TestVerifyStatusCodeGreet(t *testing.T) {
	r := SetupTestRoutes()
	r.GET("/:name", controllers.Greet)
	req, _ := http.NewRequest("GET", "/will", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)
	if response.Code != http.StatusOK {
		t.Fatalf("Expected status code %d but got %d", http.StatusOK, response.Code)
	}
}
