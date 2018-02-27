package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

const defaultConfigPath = ".get-to-work"

// PivotalTrackerConfig is a struct that saves PT confituration information
type PivotalTrackerConfig struct {
	ProjectID   string `json:"project_id"`
	LastStoryID int    `json:"last_story_id,omitempty"`
}

// HarvestConfig is a struct that saves Harvest configuration information
type HarvestConfig struct {
	AccountID      string `json:"account_id"`
	ProjectID      string `json:"project_id"`
	TaskID         string `json:"task_id"`
	LatTimeEntryID int    `json:"last_time_entry_id,omitempty"`
}

// Config is a struct that contains configuration information
type Config struct {
	Harvest        HarvestConfig        `json:"harvest"`
	PivotalTracker PivotalTrackerConfig `json:"pivotal_tracker"`
}

// FromFile returns a Config given a file path.
func FromFile(path string) (cfg Config, err error) {
	cfg = Config{}

	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		// Create the file
		suberr := cfg.Save(path)
		if suberr != nil {
			return
		}
	}

	fileContents, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = json.Unmarshal(fileContents, &cfg)
	return
}

// Save persists the current state of the config struct to a fileContents
func (config *Config) Save(path string) (err error) {
	configJSON, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = ioutil.WriteFile(path, configJSON, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

	return
}

// DefaultConfig returns the defualt config from ".get-to-work"
func DefaultConfig() (config Config, e error) {
	cfg, err := FromFile(defaultConfigPath)
	return cfg, err
}

// SaveDefaultConfig saves the default config file
func (config *Config) SaveDefaultConfig() (err error) {
	err = config.Save(defaultConfigPath)
	return
}
