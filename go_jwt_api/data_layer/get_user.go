package data_layer

import (
	"database/sql"
	"fmt"
	"go_tutorials/go_jwt_api/models"
)

func (a *Database) GetUser(email string) (*models.User, error) {
	query := "SELECT * FROM users WHERE email = ?"

	// Initialize a User struct to hold the result
	var user models.User

	// Execute the query and scan the result into the user struct
	err := a.db.Get(&user, query, email)
	if err != nil {
		if err == sql.ErrNoRows {
			// No user found with the given email
			return &user, fmt.Errorf("User not found")
		}
		return &user, err
	}

	return &user, nil
}
