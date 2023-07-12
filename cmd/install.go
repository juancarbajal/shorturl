package main

import (
	"github.com/juancarbajal/shorturl/configs"
	"github.com/juancarbajal/shorturl/entities"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, _ := gorm.Open(sqlite.Open(configs.DB_NAME), &gorm.Config{})
	db.AutoMigrate(&entities.Shorten{})
}
