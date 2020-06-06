package models

import (
	"time"
)

// Task struct outlines task data
type Task struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	Task      string    `json:"task"`
	Completed bool      `json:"completed"`
	DateEnter time.Time `json:"date_enter"`
	DateComp  string    `json:"date_comp"`
}

// GetUserTask retreives task for a specific user using the id of the task
func GetUserTask(username string, id int64) (Task, error) {
	// TODO
	return *&Task{}, nil
}

// GetUserTasks retreives all user tasks
func GetUserTasks(username string) ([]Task, error) {
	// TODO
	return nil, nil
}

// UpdateUserTask updates the task of a specific user
func UpdateUserTask(username string, id int64) error {
	// TODO
	return nil
}

// DeleteUserTask removes specific task from user tasks
func DeleteUserTask(usernam string, id int64) error {
	// TODO
	return nil
}
