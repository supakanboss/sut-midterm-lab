package entity

import (
	//"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}

func SetupDatabase() {
	///////////ทดลองให้เวลาเกิดเป็นเวลาปัจุบัน///////////
	//t := time.Now()
	//////////////////////////////////////////////

	database, err := gorm.Open(sqlite.Open("sut-midterm-lab.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	//////////////////////////////////////////////
	db = database
	//////////////////////////////////////////////
	database.AutoMigrate(
		&GENDER{},
		&PROVINCE{},
		&DEGREE{},
		&INSTITUTE{},
		&PREFIX{},
	)
	///////////////ข้อมูลใน entity GENDER///////////
	GenderMale := GENDER{
		Gender_Type: "ชาย",
	}
	db.Model(&GENDER{}).Create(&GenderMale)

	GenderFemale := GENDER{
		Gender_Type: "หญิง",
	}
	db.Model(&GENDER{}).Create(&GenderFemale)

	GenderOther := GENDER{
		Gender_Type: "อื่นๆ",
	}
	db.Model(&GENDER{}).Create(&GenderOther)
	///////////////ข้อมูลใน entity PROVINCE/////////
	Province1 := PROVINCE{
		Province_Name: "กระบี่",
	}
	db.Model(&PROVINCE{}).Create(&Province1)

	Province2 := PROVINCE{
		Province_Name: "กรุงเทพมหานคร",
	}
	db.Model(&PROVINCE{}).Create(&Province2)

	Province3 := PROVINCE{
		Province_Name: "กาญจนบุรี",
	}
	db.Model(&PROVINCE{}).Create(&Province3)

	Province4 := PROVINCE{
		Province_Name: "กาฬสินธุ์",
	}
	db.Model(&PROVINCE{}).Create(&Province4)

	Province5 := PROVINCE{
		Province_Name: "กำแพงเพชร",
	}
	db.Model(&PROVINCE{}).Create(&Province5)

	Province6 := PROVINCE{
		Province_Name: "ขอนแก่น",
	}
	db.Model(&PROVINCE{}).Create(&Province6)

	Province7 := PROVINCE{
		Province_Name: "จันทบุรี",
	}
	db.Model(&PROVINCE{}).Create(&Province7)

	Province8 := PROVINCE{
		Province_Name: "ฉะเชิงเทรา",
	}
	db.Model(&PROVINCE{}).Create(&Province8)

	Province9 := PROVINCE{
		Province_Name: "ชลบุรี",
	}
	db.Model(&PROVINCE{}).Create(&Province9)

	Province10 := PROVINCE{
		Province_Name: "ชัยนาท",
	}
	db.Model(&PROVINCE{}).Create(&Province10)

	Province11 := PROVINCE{
		Province_Name: "ชัยภูมิ",
	}
	db.Model(&PROVINCE{}).Create(&Province11)

	Province12 := PROVINCE{
		Province_Name: "ชุมพร",
	}
	db.Model(&PROVINCE{}).Create(&Province12)

	Province13 := PROVINCE{
		Province_Name: "เชียงราย",
	}
	db.Model(&PROVINCE{}).Create(&Province13)

	Province14 := PROVINCE{
		Province_Name: "เชียงใหม่",
	}
	db.Model(&PROVINCE{}).Create(&Province14)

	Province15 := PROVINCE{
		Province_Name: "ตรัง",
	}
	db.Model(&PROVINCE{}).Create(&Province15)

	Province16 := PROVINCE{
		Province_Name: "ตราด",
	}
	db.Model(&PROVINCE{}).Create(&Province16)

	Province17 := PROVINCE{
		Province_Name: "ตาก",
	}
	db.Model(&PROVINCE{}).Create(&Province17)

	Province18 := PROVINCE{
		Province_Name: "นครนายก",
	}
	db.Model(&PROVINCE{}).Create(&Province18)

	Province19 := PROVINCE{
		Province_Name: "นครปฐม",
	}
	db.Model(&PROVINCE{}).Create(&Province19)

	Province20 := PROVINCE{
		Province_Name: "นครพนม",
	}
	db.Model(&PROVINCE{}).Create(&Province20)

	Province21 := PROVINCE{
		Province_Name: "นครราชสีมา",
	}
	db.Model(&PROVINCE{}).Create(&Province21)

	Province22 := PROVINCE{
		Province_Name: "นครศรีธรรมราช",
	}
	db.Model(&PROVINCE{}).Create(&Province22)

	Province23 := PROVINCE{
		Province_Name: "นครสวรรค์",
	}
	db.Model(&PROVINCE{}).Create(&Province23)

	Province24 := PROVINCE{
		Province_Name: "นนทบุรี",
	}
	db.Model(&PROVINCE{}).Create(&Province24)

	Province25 := PROVINCE{
		Province_Name: "นราธิวาส",
	}
	db.Model(&PROVINCE{}).Create(&Province25)

	Province26 := PROVINCE{
		Province_Name: "น่าน",
	}
	db.Model(&PROVINCE{}).Create(&Province26)

	Province27 := PROVINCE{
		Province_Name: "บึงกาฬ",
	}
	db.Model(&PROVINCE{}).Create(&Province27)

	Province28 := PROVINCE{
		Province_Name: "บุรีรัมย์",
	}
	db.Model(&PROVINCE{}).Create(&Province28)

	Province29 := PROVINCE{
		Province_Name: "ปทุมธานี",
	}
	db.Model(&PROVINCE{}).Create(&Province29)

	Province30 := PROVINCE{
		Province_Name: "ประจวบคีรีขันธ์",
	}
	db.Model(&PROVINCE{}).Create(&Province30)

	Province31 := PROVINCE{
		Province_Name: "ปราจีนบุรี",
	}
	db.Model(&PROVINCE{}).Create(&Province31)

	Province32 := PROVINCE{
		Province_Name: "ปัตตานี",
	}
	db.Model(&PROVINCE{}).Create(&Province32)

	Province33 := PROVINCE{
		Province_Name: "พะเยา",
	}
	db.Model(&PROVINCE{}).Create(&Province33)

	Province34 := PROVINCE{
		Province_Name: "พระนครศรีอยุธยา",
	}
	db.Model(&PROVINCE{}).Create(&Province34)

	Province35 := PROVINCE{
		Province_Name: "พังงา",
	}
	db.Model(&PROVINCE{}).Create(&Province35)

	Province36 := PROVINCE{
		Province_Name: "พัทลุง",
	}
	db.Model(&PROVINCE{}).Create(&Province36)

	Province37 := PROVINCE{
		Province_Name: "พิจิตร",
	}
	db.Model(&PROVINCE{}).Create(&Province37)

	Province38 := PROVINCE{
		Province_Name: "พิษณุโลก",
	}
	db.Model(&PROVINCE{}).Create(&Province38)

	Province39 := PROVINCE{
		Province_Name: "เพชรบุรี",
	}
	db.Model(&PROVINCE{}).Create(&Province39)

	Province40 := PROVINCE{
		Province_Name: "เพชรบูรณ์",
	}
	db.Model(&PROVINCE{}).Create(&Province40)

	Province41 := PROVINCE{
		Province_Name: "แพร่",
	}
	db.Model(&PROVINCE{}).Create(&Province41)

	Province42 := PROVINCE{
		Province_Name: "ภูเก็ต",
	}
	db.Model(&PROVINCE{}).Create(&Province42)

	Province43 := PROVINCE{
		Province_Name: "มหาสารคาม",
	}
	db.Model(&PROVINCE{}).Create(&Province43)

	Province44 := PROVINCE{
		Province_Name: "มุกดาหาร",
	}
	db.Model(&PROVINCE{}).Create(&Province44)

	Province45 := PROVINCE{
		Province_Name: "แม่ฮ่องสอน",
	}
	db.Model(&PROVINCE{}).Create(&Province45)

	Province46 := PROVINCE{
		Province_Name: "ยโสธร",
	}
	db.Model(&PROVINCE{}).Create(&Province46)

	Province47 := PROVINCE{
		Province_Name: "ยะลา",
	}
	db.Model(&PROVINCE{}).Create(&Province47)

	Province48 := PROVINCE{
		Province_Name: "ร้อยเอ็ด",
	}
	db.Model(&PROVINCE{}).Create(&Province48)

	Province49 := PROVINCE{
		Province_Name: "ระนอง",
	}
	db.Model(&PROVINCE{}).Create(&Province49)

	Province50 := PROVINCE{
		Province_Name: "ระยอง",
	}
	db.Model(&PROVINCE{}).Create(&Province50)

	Province51 := PROVINCE{
		Province_Name: "ราชบุรี",
	}
	db.Model(&PROVINCE{}).Create(&Province51)

	Province52 := PROVINCE{
		Province_Name: "ลพบุรี",
	}
	db.Model(&PROVINCE{}).Create(&Province52)

	Province53 := PROVINCE{
		Province_Name: "ลำปาง",
	}
	db.Model(&PROVINCE{}).Create(&Province53)

	Province54 := PROVINCE{
		Province_Name: "ลำพูน",
	}
	db.Model(&PROVINCE{}).Create(&Province54)

	Province55 := PROVINCE{
		Province_Name: "เลย",
	}
	db.Model(&PROVINCE{}).Create(&Province55)

	Province56 := PROVINCE{
		Province_Name: "ศรีสะเกษ",
	}
	db.Model(&PROVINCE{}).Create(&Province56)

	Province57 := PROVINCE{
		Province_Name: "สกลนคร",
	}
	db.Model(&PROVINCE{}).Create(&Province57)

	Province58 := PROVINCE{
		Province_Name: "สงขลา",
	}
	db.Model(&PROVINCE{}).Create(&Province58)

	Province59 := PROVINCE{
		Province_Name: "สตูล",
	}
	db.Model(&PROVINCE{}).Create(&Province59)

	Province60 := PROVINCE{
		Province_Name: "สมุทรปราการ",
	}
	db.Model(&PROVINCE{}).Create(&Province60)

	Province61 := PROVINCE{
		Province_Name: "สมุทรสงคราม",
	}
	db.Model(&PROVINCE{}).Create(&Province61)

	Province62 := PROVINCE{
		Province_Name: "สมุทรสาคร",
	}
	db.Model(&PROVINCE{}).Create(&Province62)

	Province63 := PROVINCE{
		Province_Name: "สระแก้ว",
	}
	db.Model(&PROVINCE{}).Create(&Province63)

	Province64 := PROVINCE{
		Province_Name: "สระบุรี",
	}
	db.Model(&PROVINCE{}).Create(&Province64)

	Province65 := PROVINCE{
		Province_Name: "สิงห์บุรี",
	}
	db.Model(&PROVINCE{}).Create(&Province65)

	Province66 := PROVINCE{
		Province_Name: "สุโขทัย",
	}
	db.Model(&PROVINCE{}).Create(&Province66)

	Province67 := PROVINCE{
		Province_Name: "สุพรรณบุรี",
	}
	db.Model(&PROVINCE{}).Create(&Province67)

	Province68 := PROVINCE{
		Province_Name: "สุราษฎร์ธานี",
	}
	db.Model(&PROVINCE{}).Create(&Province68)

	Province69 := PROVINCE{
		Province_Name: "สุรินทร์",
	}
	db.Model(&PROVINCE{}).Create(&Province69)

	Province70 := PROVINCE{
		Province_Name: "หนองคาย",
	}
	db.Model(&PROVINCE{}).Create(&Province70)

	Province71 := PROVINCE{
		Province_Name: "หนองบัวลำภู",
	}
	db.Model(&PROVINCE{}).Create(&Province71)

	Province72 := PROVINCE{
		Province_Name: "อ่างทอง",
	}
	db.Model(&PROVINCE{}).Create(&Province72)

	Province73 := PROVINCE{
		Province_Name: "อำนาจเจริญ",
	}
	db.Model(&PROVINCE{}).Create(&Province73)

	Province74 := PROVINCE{
		Province_Name: "อุดรธานี",
	}
	db.Model(&PROVINCE{}).Create(&Province74)

	Province75 := PROVINCE{
		Province_Name: "อุตรดิตถ์",
	}
	db.Model(&PROVINCE{}).Create(&Province75)

	Province76 := PROVINCE{
		Province_Name: "อุทัยธานี",
	}
	db.Model(&PROVINCE{}).Create(&Province76)

	Province77 := PROVINCE{
		Province_Name: "อุบลราชธานี",
	}
	db.Model(&PROVINCE{}).Create(&Province77)
	///////////////ข้อมูลใน entity DEGREE///////////
	Degree1 := DEGREE{
		Degree_Name: "ปริญญาตรี",
	}
	db.Model(&DEGREE{}).Create(&Degree1)

	Degree2 := DEGREE{
		Degree_Name: "ปริญญาโท",
	}
	db.Model(&DEGREE{}).Create(&Degree2)

	Degree3 := DEGREE{
		Degree_Name: "ปริญญาเอก",
	}
	db.Model(&DEGREE{}).Create(&Degree3)
	///////////////ข้อมูลใน entity INSTITUTE///////////
	Institute1 := INSTITUTE{
		Institute_Name: "สำนักวิชาวิศวกรรมศาสตร์",
	}
	db.Model(&INSTITUTE{}).Create(&Institute1)

	Institute2 := INSTITUTE{
		Institute_Name: "สำนักวิชาวิทยาศาสตร์",
	}
	db.Model(&INSTITUTE{}).Create(&Institute2)
	///////////////ข้อมูลใน entity PREFIX///////////
	Prefix1 := PREFIX{
		Prefix_Name: "นาย",
	}
	db.Model(&PREFIX{}).Create(&Prefix1)

	Prefix2 := PREFIX{
		Prefix_Name: "นาง",
	}
	db.Model(&PREFIX{}).Create(&Prefix2)

	Prefix3 := PREFIX{
		Prefix_Name: "นางสาว",
	}
	db.Model(&PREFIX{}).Create(&Prefix3)

	Prefix4 := PREFIX{
		Prefix_Name: "ดร.",
	}
	db.Model(&PREFIX{}).Create(&Prefix4)

	Prefix5 := PREFIX{
		Prefix_Name: "รองศาสตราจารย์ดร.",
	}
	db.Model(&PREFIX{}).Create(&Prefix5)

	Prefix6 := PREFIX{
		Prefix_Name: "ผู้ช่วยศาสตราจารย์ดร.",
	}
	db.Model(&PREFIX{}).Create(&Prefix6)

	
}
