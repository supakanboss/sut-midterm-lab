package main

import (
	"github.com/gin-gonic/gin"
	"github.com/supakanboss/sut-midterm-lab/controller"
	"github.com/supakanboss/sut-midterm-lab/entity"
)

func CORSMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

func main() {
	entity.SetupDatabase()
	r := gin.Default()
	r.Use(CORSMiddleware())

	/////////////////////////////////////////////////////////////

	//combobox Institute
	r.GET("/institute", controller.ListInstitute)
	//combobox Branch
	r.GET("/branch", controller.ListBranch)
	//combobox Course
	r.GET("/course", controller.ListCourse)
	//combobox Degree
	r.GET("/degree", controller.ListDegree)
	//combobox Prefix
	r.GET("/prefix", controller.ListPrefix)
	//combobox Gender
	r.GET("/gender", controller.ListGender)
	//combobox Province
	r.GET("/province", controller.ListProvince)

	/////////////////////////////////////////////////////////////

	//รับข้อมูลเข้าตาราง Student
	r.POST("/create_Student", controller.CreateStudent)
	//แสดงข้อมูลตาราง Student
	r.GET("/student_table", controller.ListStudentTable)
	// ดึงข้อมูล student by id
	r.GET("/student/:id", controller.ListStudentByID)
	// แก้ไขข้อมูล student
	r.PATCH("/update_student", controller.UpdateStudent)
	// ลบข้อมูล student by id
	r.DELETE("/delete_student/:id", controller.DeleteStudentByID)

	/////////////////////////////////////////////////////////////

	//รับข้อมูลเข้าตาราง Course
	r.POST("/create_course", controller.CreateCourse)
	//แสดงข้อมูลตาราง Course
	r.GET("/course_table", controller.ListCourseTable)
	// ดึงข้อมูล course by id
	r.GET("/course/:id", controller.ListStudentByID)
	// แก้ไขข้อมูล course
	r.PATCH("/update_course", controller.UpdateStudent)
	// ลบข้อมูล course by id
	r.DELETE("/delete_course/:id", controller.DeleteStudentByID)

	/////////////////////////////////////////////////////////////

	r.Run()
}
