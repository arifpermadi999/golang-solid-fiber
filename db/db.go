package db

import (
	"log"
	"os"

	"github.com/arifpermadi999/core-echo-golang/helpers"
	"github.com/arifpermadi999/core-echo-golang/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() *gorm.DB {
	helpers.GetEnv()
	app := os.Getenv("APP")
	connectionString := connectionString(app)
	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Panic(err)
	}
	DB = db
	return DB
}
func connectionString(app string) string {
	var connectionString string
	if app == "local" {
		connectionString = os.Getenv("DB_USERNAME_LOCAL") + ":" + os.Getenv("DB_PASSWORD_LOCAL") + "@tcp(" + os.Getenv("DB_HOST_LOCAL") + ":" + os.Getenv("DB_PORT_LOCAL") + ")/" + os.Getenv("DB_NAME_LOCAL") + "?charset=utf8mb4&parseTime=True&loc=Local"
	} else if app == "dev" {
		connectionString = os.Getenv("DB_USERNAME_DEV") + ":" + os.Getenv("DB_PASSWORD_DEV") + "@tcp(" + os.Getenv("DB_HOST_DEV") + ":" + os.Getenv("DB_PORT_DEV") + ")/" + os.Getenv("DB_NAME_DEV") + "?charset=utf8mb4&parseTime=True&loc=Local"
	} else {
		connectionString = os.Getenv("DB_USERNAME_PROD") + ":" + os.Getenv("DB_PASSWORD_PROD") + "@tcp(" + os.Getenv("DB_HOST_PROD") + ":" + os.Getenv("DB_PORT_PROD") + ")/" + os.Getenv("DB_NAME_PROD") + "?charset=utf8mb4&parseTime=True&loc=Local"
	}

	return connectionString

}
func GetDBInstance() *gorm.DB {
	return DB
}

func Migration(db *gorm.DB) {
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Category{})
	db.AutoMigrate(&models.Unit{})
}
