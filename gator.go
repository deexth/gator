package main

import (
	"errors"
	"fmt"
)

func handleLogin(s *state, cmd command) error {
	if cmd.name == "login" && len(cmd.args) == 0 {
		return errors.New("the login command expects a username")
	}

	if err := s.conf.SetUser(cmd.args[0]); err != nil {
		return fmt.Errorf("something went wrong: %v", err)
	}

	fmt.Println("user has been set")

	return nil
}
