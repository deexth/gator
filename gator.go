package main

import (
	"errors"
	"fmt"

	"github.com/deexth/gator/internal/database"
	"github.com/google/uuid"
)

func checkArgs(args []string) (string, error) {
	if len(args) == 0 {
		return "", errors.New("usage: command <name>")
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

func printUser(user database.User) {
	fmt.Printf(" * ID:      %v\n", user.ID)
	fmt.Printf(" * Name:    %v\n", user.Name)
}
