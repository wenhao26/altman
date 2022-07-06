package initialize

import (
	"altman/models"
	"fmt"
	"gorm.io/gorm"
	"os"
)

func InitDB() *gorm.DB {
	return GetMySqlDB()
}

func RegisterTables(db *gorm.DB) {
	err := db.AutoMigrate(
		models.Admin{},
	)
	if err != nil {
		fmt.Println("register table failed", err)
		os.Exit(0)
	}
}
