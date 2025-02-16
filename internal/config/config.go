package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DbUrl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

// TODO: should this be a value receiver instead of a pointer receiver?
func (c *Config) SetUser(name string) error {
	c.CurrentUserName = name
	return write(*c)
}

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	configPath := filepath.Join(homeDir, configFileName)
	return configPath, nil
}

func Read() (Config, error) {
	configPath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	file, err := os.Open(configPath)
	if err != nil {
		return Config{}, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	cfg := Config{}
	if err := decoder.Decode(&cfg); err != nil {
		// TODO: the error occurs here...
		return Config{}, err
	}

	return cfg, nil
}

func write(c Config) error {
	configPath, err := getConfigFilePath()
	if err != nil {
		panic(err)
	}
	file, err := os.Create(configPath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(c); err != nil {
		return err
	}
	return nil
}
