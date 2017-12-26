package config

import (
	"io/ioutil"
	"log"

	yaml "gopkg.in/yaml.v2"
)

// Config  server configuration
type Config struct {

	// DBSetting database settings
	DBSetting struct {
		Host     string `yaml:"host"`
		Database string `yaml:"database"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
	} `yaml:"db_settings"`
}

// NewConfig read config file
func NewConfig(name string) (*Config, error) {
	data, err := ioutil.ReadFile(name)
	if err != nil {
		log.Println("error in load config file")
		return nil, err
	}

	conf := &Config{}
	err = yaml.Unmarshal(data, conf)
	if err != nil {
		log.Println("error in unmarshal config file")
		return nil, err
	}

	return conf, nil
}
