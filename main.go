package main

import (
	"fmt"
	"gator/internal/config"
	"os"
)

// Read the config file again and print the contents of the config struct to the terminal.
func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println("Error reading config:", err)
		os.Exit(1)
	}

	err = cfg.SetUser("Lance")

	if err != nil {
		fmt.Println("Error writing config:", err)
		os.Exit(1)
	}

	cfg, err = config.Read()
	if err != nil {
		fmt.Println("Error reading config:", err)
		os.Exit(1)
	}

	fmt.Printf("Config: %+v\n", cfg)
}
