package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/supakanboss/sut-midterm-lab/entity"
)

func CreateCourse(c *gin.Context) {

	var Admin entity.ADMIN
	var Institute entity.INSTITUTE
	var Branch entity.BRANCH
	var Degree entity.DEGREE
	var Prefix entity.PREFIX
	var Course entity.COURSE

	if err := c.ShouldBindJSON(&Course); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"ShouldBindJSON_Course_error": err.Error()})
		return
	}

	// 10: ค้นหา prefix ด้วย id
	if tx := entity.DB().Where("id = ?", Course.PrefixID).First(&Prefix); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Prefix not found"})
		return
	}

	// 11: ค้นหา institute ด้วย id
	if tx := entity.DB().Where("id = ?", Course.InstituteID).First(&Institute); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Institute not found"})
		return
	}

	// 12: ค้นหา branch ด้วย id
	if tx := entity.DB().Where("id = ?", Course.BranchID).First(&Branch); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Branch not found"})
		return
	}

	// 13: ค้นหา degree ด้วย id
	if tx := entity.DB().Where("id = ?", Course.DegreeID).First(&Degree); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Degree not found"})
		return
	}

	// 14: สร้าง entity Course
	Data_Course := entity.COURSE{
		Course_Name:    Course.Course_Name,
		Course_Teacher: Course.Course_Teacher,
		Course_Credit:  Course.Course_Credit,
		Course_Detail:  Course.Course_Detail,
		Course_Year:    Course.Course_Year,

		Degree:    Degree,
		Prefix:    Prefix,
		Institute: Institute,
		Branch:    Branch,
		AdminID:   Course.AdminID,
		Admin:     Admin,
	}

	// 21: บันทึก
	if err := entity.DB().Create(&Data_Course).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Data_Course_error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": Data_Course})
}

// ดึงข้อมูล Student มาแสดง
func ListCourseTable(c *gin.Context) {

	var Course []entity.COURSE

	if err := entity.DB().Preload("Degree").Preload("Prefix").Preload("Institute").Preload("Branch").Preload("Admin").Raw("SELECT * FROM courses").Scan(&Course).Find(&Course).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"ListCourseTable_error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": Course})
}
