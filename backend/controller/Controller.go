package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/supakanboss/sut-midterm-lab/entity"
)

// 4: GET Institute เตรียมข้อมูลให้ combobox
func ListInstitute(c *gin.Context) {
	var Institute []entity.INSTITUTE

	if err := entity.DB().Raw("SELECT * FROM institutes").Scan(&Institute).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"ListInstitute_error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": Institute})
}

// 5: GET Branch เตรียมข้อมูลให้ combobox
func ListBranch(c *gin.Context) {
	var Branch []entity.BRANCH

	if err := entity.DB().Raw("SELECT * FROM branches").Scan(&Branch).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"ListBranch_error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": Branch})
}

// 6: GET Course เตรียมข้อมูลให้ combobox
func ListCourse(c *gin.Context) {
	var Course []entity.COURSE

	if err := entity.DB().Raw("SELECT * FROM courses").Scan(&Course).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"ListCourse_error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": Course})
}

// 7: GET Degree เตรียมข้อมูลให้ combobox
func ListDegree(c *gin.Context) {
	var Degree []entity.DEGREE

	if err := entity.DB().Raw("SELECT * FROM degrees").Scan(&Degree).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"ListDegree_error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": Degree})
}

// 8: GET Prefix เตรียมข้อมูลให้ combobox
func ListPrefix(c *gin.Context) {
	var Prefix []entity.PREFIX

	if err := entity.DB().Raw("SELECT * FROM prefixes").Scan(&Prefix).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"ListPrefix_error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": Prefix})
}

// 9: GET Gender เตรียมข้อมูลให้ combobox
func ListGender(c *gin.Context) {
	var Gender []entity.GENDER

	if err := entity.DB().Raw("SELECT * FROM genders").Scan(&Gender).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"ListGender_error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": Gender})
}

// 10: GET Province เตรียมข้อมูลให้ combobox
func ListProvince(c *gin.Context) {
	var Province []entity.PROVINCE

	if err := entity.DB().Raw("SELECT * FROM provinces").Scan(&Province).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"ListProvince_error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": Province})
}

///////////////////////////////// login //////////////////////////////////////////////////////////////////////

// GET /student/:id
func GetStudent(c *gin.Context) {
	var student entity.STUDENT
	id := c.Param("id")

	if err := entity.DB().Raw("SELECT * FROM students WHERE id = ?", id).Scan(&student).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": student})
}

// GET /staff/:id
func GetAdmin(c *gin.Context) {
	var admin entity.ADMIN
	id := c.Param("id")

	if err := entity.DB().Raw("SELECT * FROM admins WHERE id = ?", id).Scan(&admin).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": admin})
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////