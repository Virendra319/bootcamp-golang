package Controllers

import (
	"awesomeProject/Models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllStudents(c *gin.Context) {
	var students []Models.Student
	err := Models.GetAllStudents(&students)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(200, students)
	}
}

func AddStudent(c *gin.Context) {
	var student Models.Student
	c.BindJSON(&student)
	err := Models.AddNewStudent(&student)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(200, student)
	}
}

func GetStudentById(c *gin.Context) {
	id := c.Params.ByName("id")
	var student Models.Student
	err := Models.GetStudentById(&student, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(200, student)
	}
}

func UpdateStudent(c *gin.Context) {
	var student Models.Student
	id := c.Params.ByName("id")
	err := Models.GetStudentById(&student, id)
	if err != nil {
		c.JSON(http.StatusNotFound, student)
	}
	c.BindJSON(&student)
	err = Models.UpdateStudent(&student)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, student)
	}
}

func DeleteStudent(c *gin.Context) {
	var student Models.Student
	id := c.Params.ByName("id")
	err := Models.DeleteStudent(&student, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}
