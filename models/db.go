package models

import (
	"database/sql"
	"fmt"
)

// DataService is the DI class for retriving things from the database
type Dataservice struct {
	DB *sql.DB
}

// DatabaseConn represents the information needed to connect to a database
type DatabaseConn struct {
	Host     string
	User     string
	Password string
	Database string
}

// NewDataService ...
func NewDataService(db *sql.DB) Dataservice {
	return Dataservice{
		DB: db,
	}
}

// GetVersion returns the version of the database
func (d Dataservice) GetVersion() (string, error) {
	var version string
	err := d.DB.QueryRow("SELECT VERSION()").Scan(&version)
	return version, err
}

// ToConnString formats a connection string based on the DataConnection
func (d DatabaseConn) ToConnString() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", d.User, d.Password, d.Host, d.Database)
}

// BuildDatabaseConn returns a connection to the database
func BuildDatabaseConn(driver string, dc DatabaseConn) (*sql.DB, error) {

	db, err := sql.Open(driver, dc.ToConnString())
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, err
}
