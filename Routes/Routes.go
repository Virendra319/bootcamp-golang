package Routes

import (
	"awesomeProject/Controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	grp1 := r.Group("/students")
	{
		grp1.GET("get/all", Controllers.GetAllStudents)
		grp1.POST("add", Controllers.AddStudent)
		grp1.GET("get/:id", Controllers.GetStudentById)
		grp1.PUT("update/:id", Controllers.UpdateStudent)
		grp1.DELETE("delete/:id", Controllers.DeleteStudent)
	}
	return r
}
