package main

import (
	"log"
	"os"

	"github.com/deexth/gator/internal/config"
)

type state struct {
	conf *config.Config
}

func main() {
	conf, err := config.Read()
	if err != nil {
		log.Fatalf("somehing went wrong: %v\n", err)
	}

	s := state{
		conf: &conf,
	}

	cmds := commands{
		cmds: make(map[string]func(*state, command) error),
	}

	cmdLineArgs := os.Args
	if len(cmdLineArgs) < 2 {
		log.Fatal("Usage: cli <command> [args...]")
	}

	cmd := command{
		name: cmdLineArgs[1],
		args: cmdLineArgs[2:],
	}

	cmds.register(cmd.name, handleLogin)

	if err := cmds.run(&s, cmd); err != nil {
		log.Fatal(err)
	}
}
