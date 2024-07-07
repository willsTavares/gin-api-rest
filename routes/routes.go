package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/willsTavares/api-go-gin/controllers"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/students", controllers.ShowAllStudents)
	r.GET("/:name", controllers.Greet)
	r.POST("/students", controllers.AddNewStudent)
	r.DELETE("/students/:id", controllers.DeleteStudentById)
	r.PATCH("/students/:id", controllers.EditStudentById)
	r.GET("/students/:id", controllers.SearchStudentById)
	r.GET("/students/cpf/:cpf", controllers.SearchStudentByCPF)
	r.GET("/students/rg/:rg", controllers.SearchStudentByRG)
	r.Run(":5000")
}
