package repository

import (
	"fmt"
	"lingotalk-exam/resources/texts"

	"github.com/jinzhu/gorm"

	// Register gorm postgres
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"lingotalk-exam/config"
	"lingotalk-exam/model"
)

// GormDB Export *gorm.DB
var GormDB *gorm.DB
var isAutoMigrate bool

func init() {
	isAutoMigrate = true
	OpenDB()
	AutoMigrateDB()
}

// OpenDB func
func OpenDB() {
	var connectionString = fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		config.JSONConfig.DB.Host,
		config.JSONConfig.DB.Port,
		config.JSONConfig.DB.User,
		config.JSONConfig.DB.DBName,
		config.JSONConfig.DB.Password)
	var err error
	GormDB, err = gorm.Open("postgres", connectionString)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(texts.EN_DIC)
	}
}

// AutoMigrateDB func
func AutoMigrateDB() {
	// read from env
	// autoMigrate := os.Getenv("AM")
	// if autoMigrate == "true" {
	// 	GormDB.AutoMigrate(&model.Stuff{})
	// }

	if isAutoMigrate {
		GormDB.AutoMigrate(&model.Stuff{})
	}
}
