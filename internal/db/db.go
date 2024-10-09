package db

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
    "log"
)

var DB *gorm.DB

func ConnectDatabase() {
    var err error
    DB, err = gorm.Open("postgres", "host=localhost user=gorm dbname=gorm 
password=gorm sslmode=disable")
    if err != nil {
        log.Fatalf("Error connecting to database: %v", err)
    }
}
