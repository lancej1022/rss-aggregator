package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	DbUrl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

// TODO: should this be a value receiver instead of a pointer receiver?
func (c *Config) SetUser(name string) {
	c.CurrentUserName = name

	// homeDir, err := os.UserHomeDir()
	// if err != nil {
	// 	panic(err)
	// }
	// configFilePath := filepath.Join(homeDir, configFileName)
	configPath, err := getConfigFilePath()
	if err != nil {
		panic(err)
	}
	file, err := os.Create(configPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(c); err != nil {
		panic(err)
	}
}
