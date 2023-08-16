package data_layer

import (
	"go_tutorials/go_jwt_api/models"
)

func (a *Database) UpdateUser(user *models.User) error {
	query := "update users SET name = :name, username= :username, email=:email, password=:password where email = :email"
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
