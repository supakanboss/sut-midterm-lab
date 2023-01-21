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
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

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

	//combobox Institute
	r.GET("/Institute", controller.ListInstitute)
	//combobox Branch
	r.GET("/Branch", controller.ListBranch)
	//combobox Course
	r.GET("/Course", controller.ListCourse)
	//combobox Degree
	r.GET("/Degree", controller.ListDegree)
	//combobox Prefix
	r.GET("/Prefix", controller.ListPrefix)
	//combobox Gender
	r.GET("/Gender", controller.ListGender)
	//combobox Province
	r.GET("/Province", controller.ListProvince)

	r.Run()
}
