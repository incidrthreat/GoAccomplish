package models

import (
	"errors"
)

// ErrRemovingOnlyAdmin ensures there is always an admin in the system
var ErrRemovingOnlyAdmin = errors.New("Cannot remove the final admin")

// User struct outlines user data
type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username" sql:"not null;unique"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

// GetUser returns a user based on the ID, no user found system returns an error.
func GetUser(id int64) (User, error) {
	// TODO
	return *&User{}, nil
}

// GetUsers returns an array of all registered users
func GetUsers() ([]User, error) {
	// TODO
	return nil, nil
}

// GetUserByUsername retreives data from db based on username
func GetUserByUsername(username string) (User, error) {
	// TODO
	return *&User{}, nil
}

// UpdateUser updates user data
func UpdateUser(user *User) error {
	// TODO
	return nil
}

// DeleteUser deletes a user from the db
func DeleteUser(id int64) error {
	// TODO
	return nil
}
