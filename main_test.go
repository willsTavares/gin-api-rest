package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/willsTavares/api-go-gin/controllers"
	"github.com/willsTavares/api-go-gin/database"
	"github.com/willsTavares/api-go-gin/models"
)

var ID int

func SetupTestRoutes() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	routes := gin.Default()
	return routes
}

func CreateStudentMock() {
	student := models.Student{
		Name: "Test",
		CPF:  "12345678901",
		RG:   "123456789",
	}
	database.DB.Create(&student)
	ID = int(student.ID)
}

func DeleteStudentMock() {
	var student models.Student
	database.DB.Delete(&student, ID)
}

func TestVerifyStatusCodeGreet(t *testing.T) {
	r := SetupTestRoutes()
	r.GET("/:name", controllers.Greet)
	req, _ := http.NewRequest("GET", "/will", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)
	assert.Equal(t, http.StatusOK, response.Code, "OK response is expected")
}

func TestListAllStudentsHandler(t *testing.T) {
	database.ConnectToDatabase()
	CreateStudentMock()
	defer DeleteStudentMock()
	r := SetupTestRoutes()
	r.GET("/students", controllers.ListAllStudents)
	req, _ := http.NewRequest("GET", "/students", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)
	assert.Equal(t, http.StatusOK, response.Code, "OK response is expected")
}

func TeastSearchStudentByIdHandler(t *testing.T) {
	database.ConnectToDatabase()
	CreateStudentMock()
	defer DeleteStudentMock()
	r := SetupTestRoutes()
	r.GET("/students/:id", controllers.SearchStudentById)
	req, _ := http.NewRequest("GET", "/students/1", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)
	assert.Equal(t, http.StatusOK, response.Code, "OK response is expected")
}

func TestSearchStudentByCPFHandler(t *testing.T) {
	database.ConnectToDatabase()
	CreateStudentMock()
	defer DeleteStudentMock()
	r := SetupTestRoutes()
	r.GET("/students/:cpf", controllers.SearchStudentByCPF)
	req, _ := http.NewRequest("GET", "/students/12345678901", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)
	assert.Equal(t, http.StatusOK, response.Code, "OK response is expected")
}

func TestSearchStudentByRGHandler(t *testing.T) {
	database.ConnectToDatabase()
	CreateStudentMock()
	defer DeleteStudentMock()
	r := SetupTestRoutes()
	r.GET("/students/:rg", controllers.SearchStudentByRG)
	req, _ := http.NewRequest("GET", "/students/123456789", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)
	assert.Equal(t, http.StatusOK, response.Code, "OK response is expected")
}
