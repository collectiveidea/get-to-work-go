package config

import (
  "encoding/json"
  "io/ioutil"
  "os"
  "fmt"
)

type HarvestConfig struct {
  Subdomain string `json:"subdomain"`
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
    configJSON, _ := json.MarshalIndent(config, "", "  ")
    ioutil.WriteFile(path, configJSON, 0644)
    return config, nil
  }

  fileContents, e := ioutil.ReadFile(path)
  if (e == nil) {
    json.Unmarshal(fileContents, &config)
  }

  return config, e
}
