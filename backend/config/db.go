package config

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
    dsn := "root:@tcp(localhost:3306)/ecommerce?parseTime=true"
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Database unreachable: %v", err)
    }

    DB = db
    log.Println("Database connected successfully")
}
