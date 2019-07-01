package cfg

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	_, err := LoadConfig("foobar")
	if err == nil {
		t.Error("Loading an invalid config file should return an error")
	}

	conf, err := LoadConfig(filepath.Join("testdata", "beehive.conf"))
	if err != nil {
		t.Error("Error loading config file fixture")
	}

	if conf.Bees[0].Name != "echo" {
		t.Error("The first bee should be an exec bee named echo")
	}
}

func TestSaveConfig(t *testing.T) {
	tmpfile, err := ioutil.TempFile("", "example")
	if err != nil {
		t.Error("Could not create temp file")
	}
	defer os.Remove(tmpfile.Name()) // clean up

	testConfPath := filepath.Join("testdata", "beehive.conf")
	testConf, err := LoadConfig(testConfPath)
	if err != nil {
		t.Errorf("Failed to load config fixture %s: %v", testConfPath, err)
	}

	configFile := tmpfile.Name()
	err = SaveConfig(configFile, testConf)
	if err != nil {
		t.Errorf("Failed to save the config to %s", configFile)
	}

	if !Exist(configFile) {
		t.Error("Configuration file wasn't saved")
	}
}

func TestSaveCurrentConfig(t *testing.T) {
	tmpfile, err := ioutil.TempFile("", "example")
	if err != nil {
		t.Error("Could not create temp file")
	}
	defer os.Remove(tmpfile.Name()) // clean up

	t.Run("configFile empty", func(t *testing.T) {
		err = SaveCurrentConfig("")
		if err == nil {
			t.Error("Configuration file should not be saved")
		}
	})

	t.Run("configFile tmpfile", func(t *testing.T) {
		configFile := tmpfile.Name()
		err = SaveCurrentConfig(configFile)
		if err != nil {
			t.Error("Configuration file should have been saved")
		}
	})
}
