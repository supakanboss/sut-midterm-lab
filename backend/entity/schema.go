package entity

import (
	"time"

	"gorm.io/gorm"
)

type GENDER struct {
	gorm.Model
	Gender_Type string

	Admin   []ADMIN   `gorm:"foreignKey:GenderID"`
	Student []STUDENT `gorm:"foreignKey:GenderID"`
}
type PROVINCE struct {
	gorm.Model
	Province_Name string

	Admin   []ADMIN   `gorm:"foreignKey:ProvinceID"`
	Student []STUDENT `gorm:"foreignKey:ProvinceID"`
}
type DEGREE struct {
	gorm.Model
	Degree_Name string

	Course  []COURSE  `gorm:"foreignKey:DegreeID"`
	Student []STUDENT `gorm:"foreignKey:DegreeID"`
}
type PREFIX struct {
	gorm.Model
	Prefix_Name string

	Admin   []ADMIN   `gorm:"foreignKey:PrefixID"`
	Branch  []BRANCH  `gorm:"foreignKey:PrefixID"`
	Course  []COURSE  `gorm:"foreignKey:PrefixID"`
	Student []STUDENT `gorm:"foreignKey:PrefixID"`
}
type INSTITUTE struct {
	gorm.Model
	Institute_Name string

	Branch  []BRANCH  `gorm:"foreignKey:InstituteID"`
	Course  []COURSE  `gorm:"foreignKey:InstituteID"`
	Student []STUDENT `gorm:"foreignKey:InstituteID"`
}
type ADMIN struct {
	gorm.Model
	Admin_Name     string
	Admin_Email    string `gorm:"uniqueIndex"`
	Admin_Password string
	Admin_Tel      string
	Admin_Address  string

	GenderID   *uint
	PrefixID   *uint
	ProvinceID *uint

	Gender   GENDER
	Prefix   PREFIX
	Province PROVINCE

	Branch  []BRANCH  `gorm:"foreignKey:AdminID"`
	Course  []COURSE  `gorm:"foreignKey:AdminID"`
	Student []STUDENT `gorm:"foreignKey:AdminID"`
}
type BRANCH struct {
	gorm.Model
	Branch_Name    string
	Branch_Teacher string
	Branch_Info    string

	PrefixID    *uint
	InstituteID *uint
	AdminID     *uint

	Prefix    PREFIX
	Institute INSTITUTE
	Admin     ADMIN

	Course  []COURSE  `gorm:"foreignKey:BranchID"`
	Student []STUDENT `gorm:"foreignKey:BranchID"`
}

type COURSE struct {
	gorm.Model
	Course_Name    string
	Course_Teacher string
	Course_Credit  uint `valid:"range(120|200)"`
	Course_Detail  string `valid:"minstringlength(10)"`
	Course_Year    uint `valid:"range(2560|9999)"`

	DegreeID    *uint
	PrefixID    *uint
	InstituteID *uint
	BranchID    *uint
	AdminID     *uint

	Degree    DEGREE
	Prefix    PREFIX
	Institute INSTITUTE
	Branch    BRANCH
	Admin     ADMIN

	Student []STUDENT `gorm:"foreignKey:CourseID"`
}
type STUDENT struct {
	gorm.Model
	Student_Year_Of_Entry time.Time
	Student_Number        string `gorm:"uniqueIndex" valid:"required ,matches(^[BMD]\\d{7}$)"`
	Student_Name          string
	Student_Birthday      time.Time
	Student_Tel           string
	Student_Identity_Card string `gorm:"uniqueIndex" valid:"required ,matches(^\\d{13}$)"`
	Student_Nationality   string
	Student_Religion      string
	Student_Address       string `valid:"minstringlength(10)"`
	Student_Fathers_Name  string
	Student_Mothers_Name  string

	GenderID    *uint
	DegreeID    *uint
	PrefixID    *uint
	InstituteID *uint
	ProvinceID  *uint
	BranchID    *uint
	CourseID    *uint
	AdminID     *uint

	Gender    GENDER
	Degree    DEGREE
	Prefix    PREFIX
	Institute INSTITUTE
	Province  PROVINCE
	Branch    BRANCH
	Course    COURSE
	Admin     ADMIN
}
