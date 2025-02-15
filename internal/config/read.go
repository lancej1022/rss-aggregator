package config

import (
	"encoding/json"
	"os"
)

func Read() Config {
	// homeDir, err := os.UserHomeDir()
	// if err != nil {
	// 	panic(err)
	// }

	// configPath := filepath.Join(homeDir, configFileName)
	configPath, err := getConfigFilePath()
	if err != nil {
		panic(err)
	}
	file, err := os.Open(configPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var config Config
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		panic(err)
	}

	return config
}
