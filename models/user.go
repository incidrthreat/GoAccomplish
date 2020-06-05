package models

import (
	"errors"
)

// ErrRemovingOnlyAdmin ensures there is always an admin in the system
var ErrRemovingOnlyAdmin = errors.New("Cannot remove the final admin")

// User struct outlines user data
type User struct {
	ID	int64 `json:"id"`
	Username	string `json:"username" sql:"not null;unique"`
	Password	string `json:"password"`
	Role	string `json:"role"`
}

// GetUser returns a user based on the ID
func GetUser(id int64) (User, error) {
	return nil, nil
}