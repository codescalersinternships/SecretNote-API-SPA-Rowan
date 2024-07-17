package pkg

import (
	"github.com/glebarez/sqlite"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"log"
)

type DB struct {
	db *gorm.DB
}
type Note struct {
	gorm.Model
	ID      uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Title   string
	Content string
	User    User
}

type User struct {
	gorm.Model
	Username string `gorm:"unique"`
	Password string
}

func NewDB() *DB {
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("cannot launch database, we are in big trouble")
	}
	newDB := DB{db: db}
	return &newDB
}

// imp check if user already exists !!
func (dbm *DB) CreateUser(username, password string) error {
	user := User{Username: username, Password: password}
	result := dbm.db.Create(&user)
	return result.Error
}

func (dbm *DB) GetUserByUsername(username string) (User, error) {
	var user User
	result := dbm.db.First(&user, "username = ?", username)
	return user, result.Error
}

func (dbm *DB) GetUserByID(id uint) (User, error) {
	var user User
	result := dbm.db.First(&user, id)
	return user, result.Error
}

func (dbm *DB) CreateNote(note Note) error {
	result := dbm.db.Create(&note)
	return result.Error
}

func (dbm *DB) GetNote(id uuid.UUID) (Note, error) {
	var note Note
	result := dbm.db.First(&note, "id = ?", id)
	return note, result.Error
}
