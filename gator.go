package main

import (
	"errors"
	"fmt"

	"github.com/deexth/gator/internal/database"
	"github.com/deexth/gator/rss"
	"github.com/google/uuid"
)

func checkArgs(args []string) (string, error) {
	if len(args) == 0 {
		return "", errors.New("usage: command <argument>")
	}

	return args[0], nil
}

func handleRegister(s *state, cmd command) error {
	arg, err := checkArgs(cmd.args)
	if err != nil {
		return err
	}

	id := uuid.New()

	params := database.CreateUserParams{
		ID:   id,
		Name: arg,
	}

	user, err := s.db.CreateUser(s.ctx, params)
	if err != nil {
		return err
	}

	if err := s.conf.SetUser(user.Name); err != nil {
		return fmt.Errorf("something went wrong: %v", err)
	}

	fmt.Println("user has been registered successfully")

	printUser(user)

	return nil
}

func handleLogin(s *state, cmd command) error {
	arg, err := checkArgs(cmd.args)
	if err != nil {
		return err
	}

	user, err := s.db.GetUser(s.ctx, arg)
	if err != nil {
		return err
	}

	if err := s.conf.SetUser(user.Name); err != nil {
		return fmt.Errorf("something went wrong: %v", err)
	}

	fmt.Println("user logged in successfully")

	return nil
}

func handleListUsers(s *state, cmd command) error {
	allUsers, err := s.db.GetUsers(s.ctx)
	if err != nil {
		return fmt.Errorf("issue getting users list: %v", err)
	}

	for _, user := range allUsers {
		if s.conf.UserName == user.Name {
			fmt.Printf("* %s (current)\n", user.Name)
		} else {
			fmt.Printf("* %s\n", user.Name)
		}
	}

	return nil
}

func handleAgg(s *state, cmd command) error {
	url := "https://www.wagslane.dev/index.xml"
	feed, err := rss.FetchFeed(s.ctx, url)
	if err != nil {
		return fmt.Errorf("an issue occured while fetching the feed at %s: %v", url, err)
	}

	fmt.Printf("Below is the feed at %s: %v\n", url, feed)

	return nil
}

func handleAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.args) < 2 {
		return errors.New("usage: addfeed <name> <url>")
	}

	params := database.AddFeedParams{
		ID:     uuid.New(),
		Name:   cmd.args[0],
		Url:    cmd.args[1],
		Name_2: user.Name,
		ID_2:   uuid.New(),
	}

	_, err := s.db.AddFeed(s.ctx, params)
	if err != nil {
		return fmt.Errorf("issue while adding feed: %v", err)
	}

	return nil
}

func handleGetUsersAndFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetUsersAndFeeds(s.ctx)
	if err != nil {
		return errors.New("issue retrieving feeds")
	}

	for _, f := range feeds {
		fmt.Printf(". %s\n", f.FeedName)
		fmt.Printf(". %s\n", f.Url)
		fmt.Printf(". %s\n", f.Username)
	}

	return nil
}

func handleFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeeds(s.ctx)
	if err != nil {
		return errors.New("issue retrieving feeds")
	}

	for _, f := range feeds {
		fmt.Printf("- %s: %s", f.Name, f.Url)
	}

	return nil
}

func handleFollow(s *state, cmd command, user database.User) error {
	url, err := checkArgs(cmd.args)
	if err != nil {
		return err
	}

	feed, err := s.db.CreateFeedFollow(s.ctx, database.CreateFeedFollowParams{
		ID:   uuid.New(),
		Name: user.Name,
		Url:  url,
	})
	if err != nil {
		return fmt.Errorf("issue following feed: %s", url)
	}

	fmt.Printf("%s now followed by %s", feed.FeedName, feed.UserName)

	return nil
}

func handleFollowing(s *state, cmd command, user database.User) error {
	feeds, err := s.db.GetFeedFollowsForUser(s.ctx, user.Name)
	if err != nil {
		return fmt.Errorf("issue retrieving feeds followed by %s", user.Name)
	}

	if len(feeds) == 0 {
		fmt.Print("Looks like you don't follow any feeds yet")
		return nil
	}

	for _, feed := range feeds {
		fmt.Printf(". %s\n", feed)
	}

	return nil
}

func handleUnfollowFeed(s *state, cmd command, user database.User) error {
	url, err := checkArgs(cmd.args)
	if err != nil {
		return err
	}

	err = s.db.UnfollowFeed(s.ctx, database.UnfollowFeedParams{
		Name: user.Name,
		Url:  url,
	})
	if err != nil {
		return fmt.Errorf("issue trying to unfollow %s: %v", url, err)
	}

	return nil
}

func handleReset(s *state, cmd command) error {
	if err := s.db.ResetDb(s.ctx); err != nil {
		return errors.New("issue reseting db")
	}

	fmt.Println("Database reset successfully!")

	return nil
}

func printUser(user database.User) {
	fmt.Printf(" * ID:      %v\n", user.ID)
	fmt.Printf(" * Name:    %v\n", user.Name)
}
