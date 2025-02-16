package main

import "errors"

/*
Create a command struct.
A command contains a name and a slice of string arguments.
For example, in the case of the login command, the name would be "login" and the handler will expect the arguments slice to contain one string, the username.
*/
type command struct {
	Name string
	Args []string
}

type commands struct {
	registeredCommands map[string]func(*state, command) error
}

// registers a new handler function for a command name.
func (c *commands) register(name string, f func(*state, command) error) {
	c.registeredCommands[name] = f
}

// runs a given command with the provided state if it exists.
func (c *commands) run(s *state, cmd command) error {
	handler, ok := c.registeredCommands[cmd.Name]
	if !ok {
		return errors.New("command not found")
	}
	return handler(s, cmd)
}
