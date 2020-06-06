package main

import (
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"github.com/incidrthreat/GoAccomplish/config"
	"github.com/incidrthreat/GoAccomplish/models"
)

const (
	htmlDir   = "./ui/templates"
	staticDir = "./ui/static"
)

func main() {
	log.Println("Starting Server...")

	config, err := config.ConfigurationFromFile("config.json")
	if err != nil {
		panic(err)
	}

	dc := models.DatabaseConn{
		Host:     config.DB.Host,
		User:     config.DB.User,
		Password: config.DB.Password,
		Database: config.DB.Name,
	}

	db, err := models.BuildDatabaseConn("mysql", dc)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	dataService := models.NewDataService(db)

	version, _ := dataService.GetVersion()

	log.Printf("Connection to Database established. Version: %s", version)

	store := sessions.NewCookieStore(
		securecookie.GenerateRandomKey(64), // authentication key
		securecookie.GenerateRandomKey(32), // encryption key (AES-256)
	)

	store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7, // Sessions good for 1 week
		HttpOnly: true,
	}

	log.Printf("Starting server on %s", config.ListenInterface)

	err = http.ListenAndServe(config.ListenInterface, nil)
	log.Fatal(err)
}
