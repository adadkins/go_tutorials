package data_layer

import (
	"go_tutorials/go_jwt_api/models"
)

func (a *Database) InsertUser(user *models.User) error {
	query := "INSERT INTO users (name, username, email, password) VALUES (:name, :username, :email, :password)"
	// Prepare the statement
	stmt, err := a.db.PrepareNamed(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Execute the statement with the user data
	_, err = stmt.Exec(user)
	return err
}
