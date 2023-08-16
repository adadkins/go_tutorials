package data_layer

import (
	"fmt"
	"go_tutorials/go_jwt_api/models"

	"github.com/jmoiron/sqlx"
)

type DataLayer interface {
	InsertUser(user *models.User) error
	GetUser(email string) (*models.User, error)
	UpdateUser(user *models.User) error
}

type Database struct {
	db *sqlx.DB
}

func New() (Database, error) {
	dl := Database{db: connect()}
	return dl, nil
}

func connect() *sqlx.DB {
	connectionString := "./mydatabase.db" // Change this to the path of your SQLite database file

	// Open a connection to the SQLite database
	db, err := sqlx.Open("sqlite3", connectionString)
	if err != nil {
		fmt.Println("Error opening database:", err)
		return nil
	}
	// defer db.Close()
	fmt.Println("Ping...")

	err = db.Ping()
	if err != nil {
		fmt.Println("Error opening database:", err)
		return nil
	}
	fmt.Println("Pong")
	return db
}

func (d *Database) CloseDB() {
	d.db.DB.Close()
}
