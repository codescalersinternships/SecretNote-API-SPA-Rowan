package main

import (
	// "gorm.io/driver/sqlite"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"time"
)

type Note struct {
	gorm.Model
	Title         string
	Content       string
	Creation_date time.Time
	UserID        uint
	User          User
}

type User struct {
	gorm.Model
	Username string
	Password string
}
var db, _ = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
func main() {
	db.AutoMigrate(&User{})
}
