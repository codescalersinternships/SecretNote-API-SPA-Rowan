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

func main() {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	print(db)
	print(err)
}
