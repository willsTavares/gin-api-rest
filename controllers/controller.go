package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/willsTavares/api-go-gin/database"
	"github.com/willsTavares/api-go-gin/models"
)

func Greet(c *gin.Context) {
	nome := c.Param("name")
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello, " + nome,
	})
}

func ShowAllStudents(c *gin.Context) {
	var students []models.Student
	database.DB.Find(&students)
	c.JSON(http.StatusOK, students)
}

func AddNewStudent(c *gin.Context) {
	var student models.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	if err := models.ValidateStudent(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	database.DB.Create(&student)
	c.JSON(http.StatusCreated, student)
}

func DeleteStudentById(c *gin.Context) {
	id := c.Param("id")
	var student models.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	database.DB.Delete(&student, id)
	c.JSON(http.StatusOK, gin.H{
		"message": "Student deleted successfully",
	})
}

func EditStudentById(c *gin.Context) {
	id := c.Param("id")
	var student models.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	if err := models.ValidateStudent(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	database.DB.First(&student, id)
	database.DB.Model(&student).UpdateColumns(student)
	c.JSON(http.StatusOK, gin.H{
		"message": "Student updated successfully",
	})
}

func SearchStudentById(c *gin.Context) {
	id := c.Param("id")
	var student models.Student
	database.DB.First(&student, id)
	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Student not found"})
		return
	}
	c.JSON(http.StatusOK, student)
}

func SearchStudentByCPF(c *gin.Context) {
	cpf := c.Param("cpf")
	var student models.Student
	database.DB.Where(&models.Student{CPF: cpf}).First(&student)
	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Student not found"})
		return
	}
	c.JSON(http.StatusOK, student)
}

func SearchStudentByRG(c *gin.Context) {
	rg := c.Param("rg")
	var student models.Student
	database.DB.Where(&models.Student{RG: rg}).First(&student)
	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Student not found"})
		return
	}
	c.JSON(http.StatusOK, student)
}
