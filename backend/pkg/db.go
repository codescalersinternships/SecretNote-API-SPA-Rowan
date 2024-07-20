package pkg

import (
	"github.com/glebarez/sqlite"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"log"
	"time"
)

// DB struct has an instance of sqlite database
type DB struct {
	db *gorm.DB
}

// Note struct is a db model
// It has a many to 1 relation with User table
type Note struct {
	gorm.Model
	ID      uuid.UUID `gorm:"type:uuid;default`
	Title   string
	Content string
	Views   int64
	Expired bool
	UserID  uint
	User    User
}

// User struct is a db model
// It has a 1 to many relation with Note table
type User struct {
	gorm.Model
	Username string `gorm:"unique"`
	Password string
}

// NewDB initializes database and auto migrates tables
func NewDB(file ...string) (DB, error) {
	filePath := "gorm.db"
	if len(file) != 0 {
		filePath = file[0]
	}
	db, err := gorm.Open(sqlite.Open(filePath), &gorm.Config{})
	if err != nil {
		log.Fatal("cannot launch database, we are in big trouble")
	}
	newDB := DB{db: db}
	return newDB, newDB.Migrate()
}

// Migrate function handles migration of current db tables
func (dbm *DB) Migrate() error {
	return dbm.db.AutoMigrate(&User{}, &Note{})
}

// BeforeCreate is a hook db function that's called before note creation
// Sqlite in gorm can't have uuid default creation --> produces an error
func (note *Note) BeforeCreate(tx *gorm.DB) (err error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return err
	}

	note.ID = id
	return
}

// CreateUser creates a user in db, returns an error if it can't be created
func (dbm *DB) CreateUser(username, password string) error {
	user := User{Username: username, Password: password}
	result := dbm.db.Create(&user)
	return result.Error
}

// GetUserByUsername returns user and an error if db can't retreive
func (dbm *DB) GetUserByUsername(username string) (User, error) {
	var user User
	result := dbm.db.First(&user, "Username = ?", username)
	return user, result.Error
}

// GetUserByID uses ID to fetch user, returns user and an error if db can't retreive
func (dbm *DB) GetUserByID(id uint) (User, error) {
	var user User
	result := dbm.db.First(&user, id)
	return user, result.Error
}

// CreateNote creates a note in db, it also sets userID returns an error if it can't be created
func (dbm *DB) CreateNote(note Note, user User) (Note, error) {
	note.UserID = user.ID
	note.User = user
	note.Expired = false
	note.Views = 0
	result := dbm.db.Create(&note)
	return note, result.Error
}

// GetNote fetches note returns note and an error if db can't retreive
func (dbm *DB) GetNote(id uuid.UUID) (Note, error) {
	var note Note
	result := dbm.db.First(&note, "id = ?", id)
	if !dbm.IsExpired(note) {
		note.Views++
		dbm.db.Save(&note)
	}
	return note, result.Error
}

// GetNotes fetches notes of a certain user and returns an error if db can't retreive
func (dbm *DB) GetNotes(user User) ([]Note, error) {
	var notes []Note
	result := dbm.db.Where("user_id = ?", uint(user.ID)).Find(&notes)
	return notes, result.Error
}

// GetExpiredNotes fetches expired notes of a certain user and returns an error if db can't retreive
func (dbm *DB) GetExpiredNotes(user User) ([]Note, error) {
	var notes []Note
	result := dbm.db.Where("user_id = ?", uint(user.ID)).Where("expired = ?", true).Find(&notes)
	return notes, result.Error
}

// IsExpired validates expiration of note by both View count and Creation date
func (dbm *DB) IsExpired(note Note) bool {
	if note.Views >= 10 {
		note.Expired = true
	}
	currTime := time.Now().Add(-2 * time.Hour)
	if note.CreatedAt.Year() < currTime.Year() || note.CreatedAt.Month() < currTime.Month() || note.CreatedAt.Day() < currTime.Day() {
		note.Expired = true
	}
	if note.CreatedAt.Hour() < currTime.Hour() {
		note.Expired = true
	}
	dbm.db.Save(&note)
	return note.Expired
}
