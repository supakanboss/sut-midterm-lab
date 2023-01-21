package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/supakanboss/sut-midterm-lab/entity"
)

func CreateStudent(c *gin.Context) {

	var Admin entity.ADMIN
	var Institute entity.INSTITUTE
	var Branch entity.BRANCH
	var Course entity.COURSE
	var Degree entity.DEGREE
	var Prefix entity.PREFIX
	var Gender entity.GENDER
	var Province entity.PROVINCE
	var Student entity.STUDENT

	if err := c.ShouldBindJSON(&Student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"ShouldBindJSON_Student_error": err.Error()})
		return
	}

	// 13: ค้นหา institute ด้วย id
	if tx := entity.DB().Where("id = ?", Student.InstituteID).First(&Institute); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Institute not found"})
		return
	}

	// 14: ค้นหา branch ด้วย id
	if tx := entity.DB().Where("id = ?", Student.BranchID).First(&Branch); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Branch not found"})
		return
	}

	// 15: ค้นหา course ด้วย id
	if tx := entity.DB().Where("id = ?", Student.CourseID).First(&Course); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Course not found"})
		return
	}

	// 16: ค้นหา degree ด้วย id
	if tx := entity.DB().Where("id = ?", Student.DegreeID).First(&Degree); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Degree not found"})
		return
	}

	// 17: ค้นหา prefix ด้วย id
	if tx := entity.DB().Where("id = ?", Student.PrefixID).First(&Prefix); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Prefix not found"})
		return
	}

	// 18: ค้นหา gender ด้วย id
	if tx := entity.DB().Where("id = ?", Student.GenderID).First(&Gender); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Gender not found"})
		return
	}

	// 19: ค้นหา province ด้วย id
	if tx := entity.DB().Where("id = ?", Student.ProvinceID).First(&Province); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Province not found"})
		return
	}

	// 20: สร้าง entity Student
	Data_Student := entity.STUDENT{
		Student_Year_Of_Entry: Student.Student_Year_Of_Entry,
		Student_Number:        Student.Student_Number,
		Student_Name:          Student.Student_Name,
		Student_Birthday:      Student.Student_Birthday,
		Student_Tel:           Student.Student_Tel,
		Student_Identity_Card: Student.Student_Identity_Card,
		Student_Nationality:   Student.Student_Nationality,
		Student_Religion:      Student.Student_Religion,
		Student_Address:       Student.Student_Address,
		Student_Fathers_Name:  Student.Student_Fathers_Name,
		Student_Mothers_Name:  Student.Student_Mothers_Name,

		Gender:    Gender,
		Degree:    Degree,
		Prefix:    Prefix,
		Institute: Institute,
		Province:  Province,
		Branch:    Branch,
		Course:    Course,
		AdminID:   Student.AdminID,
		Admin:     Admin,
	}

	// 21: บันทึก
	if err := entity.DB().Create(&Data_Student).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Data_Student_error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": Data_Student})
}

// ดึงข้อมูล Student มาแสดง
func ListStudentTable(c *gin.Context) {

	var Student []entity.STUDENT

	if err := entity.DB().Preload("Gender").Preload("Degree").Preload("Prefix").Preload("Institute").Preload("Province").Preload("Branch").Preload("Course").Preload("Admin").Raw("SELECT * FROM students").Scan(&Student).Find(&Student).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"ListStudent_error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": Student})
}
