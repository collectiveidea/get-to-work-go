package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFromFileNonExistant(t *testing.T) {
	path := ".non-existant-test"
	_, err := FromFile(path)
	if err != nil {
		t.Error("From File should not raise an error")
	}

	fileContent, fileReadErr := ioutil.ReadFile(path)
	if os.IsNotExist(fileReadErr) {
		t.Error("FromFile did not create a new file")
	}

	expectedJSON := `{
  "harvest": {
    "account_id": "",
    "project_id": "",
    "task_id": ""
  },
  "pivotal_tracker": {
    "project_id": ""
  }
}`

	assert.Equal(t, expectedJSON, string(fileContent), "Incorrect default content")
	err = os.Remove(path)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func TestFromFileThatExists(t *testing.T) {
	// Create a file with json in it
	path := ".existing-test-config"
	fileContent := `{
      "harvest": {
        "account_id": "foobar",
        "project_id": "",
        "task_id": ""
      },
			"pivotal_tracker": {
		    "project_id": ""
		  }
  }`

	err := ioutil.WriteFile(path, []byte(fileContent), 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

	config, err := FromFile(path)
	assert.Nil(t, err, "Config raied error with existing path")
	expected := Config{Harvest: HarvestConfig{AccountID: "foobar"}}

	assert.Equal(t, expected, config, "Config didn't load JSON from file")

	err = os.Remove(path)
	if err != nil {
		fmt.Println(err)
		return
	}
}
