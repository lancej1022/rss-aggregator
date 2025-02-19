package main

import (
	"context"
	"fmt"
	"gator/internal/database"
	"time"

	"github.com/google/uuid"
)

func handleFollowFeed(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <feed_url>", cmd.Name)
	}
	currUser, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName) // TODO: idk if this is right...
	if err != nil {
		return fmt.Errorf("couldn't get current user: %w", err)
	}

	feedUrl := cmd.Args[0]
	feed, err := s.db.GetFeedByUrl(context.Background(), feedUrl)
	if err != nil {
		return fmt.Errorf("couldn't get feed: %w", err)
	}

	feedFollow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    currUser.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return fmt.Errorf("couldn't create feed follow: %w", err)
	}

	fmt.Printf("Successfully followed feed %v\n", feedFollow.FeedName)
	fmt.Printf("Current user: %v\n", feedFollow.UserName)
	return nil
}

// print all the names of the feeds the current user is following.
func handleFollowing(s *state, cmd command) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("usage: %s", cmd.Name)
	}

	currUser, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("couldn't get current user: %w", err)
	}

	feedFollows, err := s.db.GetFeedFollowsForUser(context.Background(), currUser.ID)
	if err != nil {
		return fmt.Errorf("couldn't get feedFollows: %w", err)
	}

	if len(feedFollows) == 0 {
		fmt.Println("No feed follows found for this user.")
		return nil
	}

	fmt.Printf("Feed follows for user %s:\n", currUser.Name)
	for _, f := range feedFollows {
		fmt.Printf("* %s\n", f.FeedName)
	}
	return nil
}
