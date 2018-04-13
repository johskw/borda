package model

import (
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open("mysql", os.Getenv("BORDA_DB_CONNECTION"))
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Event{}, &Choice{}, &Vote{}, &Score{})
}
