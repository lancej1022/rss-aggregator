package main

import (
	"context"
	"fmt"
	"gator/internal/database"
)

func handleListUsers(s *state, cmd command) error {
	if len(cmd.Args) > 0 {
		return fmt.Errorf("usage: %s", cmd.Name)
	}

	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't get users: %w", err)
	}

	for _, user := range users {
		printUser(user, s.cfg.CurrentUserName)
	}

	return nil
}

func printUser(user database.User, currentUser string) {
	// fmt.Printf(" * ID:      %v\n", user.ID)
	if user.Name == currentUser {
		fmt.Printf("* %s (current)\n", user.Name)
	} else {

		fmt.Printf("* %s\n", user.Name)
	}
}
