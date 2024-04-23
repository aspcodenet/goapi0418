package data

import (
	"github.com/glebarez/sqlite"
	//"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase(filename string) {
	DB, _ = gorm.Open(sqlite.Open(filename), &gorm.Config{})
	DB.AutoMigrate(&Employee{}) // SYNKAR databas med kod KOD = sanningen
	var antal int64
	DB.Model(&Employee{}).Count(&antal) // Seed
	if antal == 0 {
		DB.Create(&Employee{Age: 50, Namn: "Stefan", City: "Test"})
		DB.Create(&Employee{Age: 14, Namn: "Oliver", City: "Test"})
		DB.Create(&Employee{Age: 20, Namn: "Josefine", City: "Test"})
	}
}
