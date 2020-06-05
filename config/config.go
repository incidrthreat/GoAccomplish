package config

import (
	"encoding/json"
	"io/ioutil"

	"github.com/pkg/errors"
)

// Configuration - the app config
type Configuration struct {
	ListenInterface string   `json:"listen_interface"`
	DB              Database `json:"database"`
}

// Database - config for the database
type Database struct {
	Host     string `json:"host"`
	Name     string `json:"name"`
	User     string `json:"user"`
	Password string `json:"password"`
}

// ConfigurationFromFile parses the given file and returns the config
func ConfigurationFromFile(fileName string) (Configuration, error) {
	var conf Configuration

	confjson, err := ioutil.ReadFile(fileName)
	if err != nil {
		return conf, errors.Wrapf(err, "Failed to open the config file at: %s", fileName)
	}

	if err := json.Unmarshal(confjson, &conf); err != nil {
		return conf, errors.Wrapf(err, "Unable to parse the config file at: %s", fileName)
	}

	return conf, nil
}
