package main

import (
	"errors"
)

type command struct {
	name string
	args []string
}

type commands struct {
	cmds map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	f, ok := c.cmds[cmd.name]
	if !ok {
		return errors.New("the command is not registered")
	}

	return f(s, cmd)
}

// func (c *commands) register(name string, f func(*state, command) error) {
// 	c.cmds[name] = f
// }

func loadCommands() commands {
	return commands{
		cmds: map[string]func(*state, command) error{
			"login":     handleLogin,
			"register":  handleRegister,
			"reset":     handleReset,
			"users":     handleListUsers,
			"agg":       handleAgg,
			"addfeed":   handleAddFeed,
			"feeds":     handleGetUsersAndFeeds,
			"allfeeds":  handleFeeds,
			"follow":    handleFollow,
			"following": handleFollowing,
		},
	}
}
func getCommandsHelp() map[string]struct {
	name        string
	description string
	usage       string
} {
	cmdStruct := map[string]struct {
		name        string
		description string
		usage       string
	}{
		"login": {
			name:        "login",
			description: "login command can be used to login a registered user, otherwise will result in an error.",
			usage:       "usage: login <name>",
		},
		"register": {
			name:        "register",
			description: "register command can be used to create a new user. It will result into an error if the user already exists.",
			usage:       "usage: register <name>",
		},
		"reset": {
			name:        "reset",
			description: "reset command can be used to reset the database to a blank state.",
			usage:       "usage: reset",
		},
		"users": {
			name:        "users",
			description: "users command is used to list all users registered in gator. The currently logged in user will have a (current) flag after the name",
			usage:       "usage: users",
		},
		"agg": {
			name:        "agg",
			description: "agg command is used to fetch feed for an rss url",
			usage:       "usage: agg",
		},
		"addfeed": {
			name:        "addfeed",
			description: "addfeed command can be used to create a feed",
			usage:       "usage: addfeed <name> <url>",
		},
		"feeds": {
			name:        "feeds",
			description: "feeds command can be used to get all feeds and their owners",
			usage:       "usage: feeds",
		},
		"allfeeds": {
			name:        "allfeeds",
			description: "allfeeds command can be used to get all feeds and their owners",
			usage:       "usage: allfeeds",
		},
		"follow": {
			name:        "follow",
			description: "follow command can be used to follow  a feed",
			usage:       "usage: follow <url>",
		},
		"following": {
			name:        "following",
			description: "following command can be used list feeds followed by a user",
			usage:       "usage: following",
		},
	}

	return cmdStruct
}
