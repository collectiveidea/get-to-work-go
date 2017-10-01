package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

const defaultConfigPath = ".get-to-work"

// HarvestConfig is a struct that saves Harvest configuration information
type HarvestConfig struct {
	Subdomain string `json:"subdomain"`
	Username  string `json:"username"`
}

// Config is a struct that contains configuration information
type Config struct {
	Harvest HarvestConfig `json:"harvest"`
}

// FromFile returns a Config given a file path.
func FromFile(path string) (cfg Config, e error) {
	fmt.Print(path)
	config := Config{}

	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		// Create the file
		config.Save(path)
		return config, nil
	}

	fileContents, e := ioutil.ReadFile(path)
	if e == nil {
		json.Unmarshal(fileContents, &config)
	}

	return config, e
}

// Save persists the current state of the config struct to a fileContents
func (config *Config) Save(path string) (err error) {
	configJSON, err := json.MarshalIndent(config, "", "  ")
	ioutil.WriteFile(path, configJSON, 0644)

	return err
}

// DefaultConfig returns the defualt config from ".get-to-work"
func DefaultConfig() (config Config, e error) {
	cfg, err := FromFile(defaultConfigPath)
	return cfg, err
}
