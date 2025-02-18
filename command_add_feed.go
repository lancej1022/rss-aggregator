package main

import (
	"context"
	"fmt"
	"gator/internal/database"
	"time"

	"github.com/google/uuid"
)

func handleAddFeed(s *state, cmd command) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("usage: %s <name> <url>", cmd.Name)
	}

	feedName := cmd.Args[0]
	feedUrl := cmd.Args[1]
	currUsername := s.cfg.CurrentUserName
	// At the top of the handler, get the current user from the database and connect the feed to that user.
	// If everything goes well, print out the fields of the new feed record.
	currUser, err := s.db.GetUser(context.Background(), currUsername)
	if err != nil {
		return fmt.Errorf("couldn't find current user: %w", err)
	}

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      feedName,
		Url:       feedUrl,
		UserID:    currUser.ID,
	})
	if err != nil {
		return fmt.Errorf("couldn't create feed: %w", err)
	}
	fmt.Printf("Successfully created feed %v\n", feed)
	return nil
}
