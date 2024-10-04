package connection

import (
	"fmt"

	"github.com/handarudwiki/raport-app-be/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDatabaseConnection(cnf *config.Config) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", cnf.Database.Host, cnf.Database.User, cnf.Database.Password, cnf.Database.Name, cnf.Database.Port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	fmt.Println("Database connection success")

	return db
}
