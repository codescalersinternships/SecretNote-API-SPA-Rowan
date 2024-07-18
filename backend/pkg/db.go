package pkg

import (
	"github.com/glebarez/sqlite"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"log"
	"time"
)

type DB struct {
	db *gorm.DB
}
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

type User struct {
	gorm.Model
	Username string `gorm:"unique"`
	Password string
}

func NewDB() (DB, error) {
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("cannot launch database, we are in big trouble")
	}
	newDB := DB{db: db}
	return newDB, newDB.Migrate()
}
func (dbm *DB) Migrate() error {
	return dbm.db.AutoMigrate(&User{}, &Note{})
}

func (note *Note) BeforeCreate(tx *gorm.DB) (err error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return err
	}

	note.ID = id
	return
}

func (dbm *DB) CreateUser(username, password string) error {
	user := User{Username: username, Password: password}
	result := dbm.db.Create(&user)
	return result.Error
}

func (dbm *DB) GetUserByUsername(username string) (User, error) {
	var user User
	result := dbm.db.First(&user, "Username = ?", username)
	return user, result.Error
}

func (dbm *DB) GetUserByID(id uint) (User, error) {
	var user User
	result := dbm.db.First(&user, id)
	return user, result.Error
}

func (dbm *DB) CreateNote(note Note, user User) (Note, error) {
	note.UserID = user.ID
	note.User = user
	note.Expired = false
	note.Views = 0
	result := dbm.db.Create(&note)
	return note, result.Error
}

func (dbm *DB) GetNote(id uuid.UUID) (Note, error) {
	var note Note
	result := dbm.db.First(&note, "id = ?", id)
	if !dbm.IsExpired(note) {
		note.Views += 1
		dbm.db.Save(&note)
	}
	return note, result.Error
}

func (dbm *DB) GetNotes(user User) ([]Note, error) {
	var notes []Note
	result := dbm.db.Where("user_id = ?", uint(user.ID)).Find(&notes)
	return notes, result.Error
}

func (dbm *DB) GetExpiredNotes(user User) ([]Note, error) {
	var notes []Note
	result := dbm.db.Where("user_id = ?", uint(user.ID)).Where("expired = ?", true).Find(&notes)
	return notes, result.Error
}

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
