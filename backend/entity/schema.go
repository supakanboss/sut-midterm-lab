package entity

import (
	"time"

	"gorm.io/gorm"
)

type GENDER struct {
	gorm.Model
	Gender_Type string

	Student []STUDENT `gorm:"foreignKey:Gender_ID"`
	Admin   []ADMIN   `gorm:"foreignKey:Gender_ID"`
}
type DEGREE struct {
	gorm.Model
	Degree_Name string

	Student []STUDENT `gorm:"foreignKey:Degree_ID"`
	Course  []COURSE  `gorm:"foreignKey:Degree_ID"`
}
type PREFIX struct {
	gorm.Model
	Prefix_Name string

	Student []STUDENT `gorm:"foreignKey:Prefix_ID"`
	Admin   []ADMIN   `gorm:"foreignKey:Prefix_ID"`
	Branch  []BRANCH  `gorm:"foreignKey:Prefix_ID"`
	Course  []COURSE  `gorm:"foreignKey:Prefix_ID"`
}
type INSTITUTE struct {
	gorm.Model
	Institute_Name string

	Student []STUDENT `gorm:"foreignKey:Institute_ID"`
	Branch  []BRANCH  `gorm:"foreignKey:Institute_ID"`
	Course  []COURSE  `gorm:"foreignKey:Institute_ID"`
}
type PROVINCE struct {
	gorm.Model
	Province_Name string

	Student []STUDENT `gorm:"foreignKey:Province_ID"`
	Admin   []ADMIN   `gorm:"foreignKey:Province_ID"`
}
type BRANCH struct {
	gorm.Model
	Branch_Name    string
	Branch_Teacher string
	Branch_Info    string

	Prefix_ID    *uint
	Institute_ID *uint

	Prefix    PREFIX
	Institute INSTITUTE
}
type COURSE struct {
	gorm.Model
	Course_Name    string
	Course_Teacher string
	Course_Credit  uint
	Course_Detail  string
	Course_Year    uint

	Degree_ID    *uint
	Prefix_ID    *uint
	Institute_ID *uint

	Degree    DEGREE
	Prefix    PREFIX
	Institute INSTITUTE
}
type ADMIN struct {
	gorm.Model
	Admin_Name     string
	Admin_Email    string `gorm:"uniqueIndex"`
	Admin_Password string
	Admin_Tel      uint
	Admin_Address  uint

	Gender_ID   *uint
	Prefix_ID   *uint
	Province_ID *uint

	Gender   GENDER
	Prefix   PREFIX
	Province PROVINCE
}
type STUDENT struct {
	gorm.Model
	Student_Year_Of_Entry time.Time
	Student_Number        string `gorm:"uniqueIndex"`
	Student_Name          string
	Student_Birthday      time.Time
	Student_Tel           string
	Student_Identity_Card string `gorm:"uniqueIndex"`
	Student_Nationality   string
	Student_Religion      string
	Student_Address       string
	Student_Fathers_Name  string
	Student_Mothers_Name  string

	Gender_ID    *uint
	Degree_ID    *uint
	Prefix_ID    *uint
	Institute_ID *uint
	Province_ID  *uint

	Gender    GENDER
	Degree    DEGREE
	Prefix    PREFIX
	Institute INSTITUTE
	Province  PROVINCE
}
