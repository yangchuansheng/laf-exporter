package config

import (
	"io/ioutil"
	"log"
	"os"
    yaml "gopkg.in/yaml.v2"
)

type Config struct {
	URL    string `yaml:"url"`
	APIKey string `yaml:"api_key"`
}

var config Config

func LoadConfig() {
	apiUrl, exists := os.LookupEnv("API_URL")
	if exists {
		config.URL = apiUrl
	}
	apiKey, exists := os.LookupEnv("API_KEY")
	if exists {
		config.APIKey = apiKey
	}

    // If the environment variables are not set, fall back to the configuration file
    if config.URL == "" || config.APIKey == "" {
        file, err := ioutil.ReadFile("config.yaml")
        if err != nil {
            log.Fatalf("Unable to load config file: %v", err)
        }

        err = yaml.Unmarshal(file, &config)
        if err != nil {
            log.Fatalf("Unable to unmarshal config: %v", err)
        }
    }
}

func GetConfig() Config {
	return config
}
