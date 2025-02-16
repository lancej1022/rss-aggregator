package main

import (
	"gator/internal/config"
	"log"
	"os"
)

type state struct {
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	programState := &state{cfg: &cfg}

	cmds := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)

	args := os.Args
	// If there are fewer than 2 arguments, print an error message to the terminal and exit.
	// Why two? The first argument is automatically the program name, which we ignore, and we require a command name.
	if len(args) < 2 {
		log.Fatal("Usage: cli <command> [args...]")
		return
	}

	// You'll need to split the command-line arguments into the command name and the arguments slice to create a command instance.
	// Use the commands.run method to run the given command and print any errors returned.
	cmd := command{
		Name: args[1],
		Args: args[2:],
	}

	err = cmds.run(programState, cmd)
	if err != nil {
		log.Fatal(err)
	}
}
