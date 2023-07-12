package configs
import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB_NAME = "shorten.db"
func DbConnect() *gorm.DB{
	db, _ := gorm.Open(sqlite.Open(DB_NAME), &gorm.Config{})
	return db
}
