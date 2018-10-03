package datatools

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// AppConfig holds the configuration from the loaded yaml file.
var AppConfig Config

// Config is built from the yaml configuration file.
// It provides all configuration values for the application.
type Config struct {
	Settings struct {
		Debug     bool   `yaml:"debug"`
		Verbose   bool   `yaml:"verbose"`
		FixedDate string `yaml:"fixedDate"`
		Output    string `yaml:"outputDir"`
		Input     string `yaml:"inputDir"`
		Logfile   string `yaml:"logfile"`
	} `yaml:"settings"`
}

// LoadConfig expects a file in yaml format.
// It returns any error encountered reading the file.
func LoadConfig(file string) error {
	configFile, err := ioutil.ReadFile(file)
	Check(err)

	Check(yaml.Unmarshal(configFile, &AppConfig))

	return err
}
