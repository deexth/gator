package main

import (
	"context"
	"database/sql"
	"log"
	"os"

	"github.com/deexth/gator/internal/config"
	"github.com/deexth/gator/internal/database"
	_ "github.com/lib/pq"
)

type state struct {
	db   *database.Queries
	conf *config.Config
	ctx  context.Context
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

	db, err := sql.Open("postgres", s.conf.DBURL)
	if err != nil {
		log.Fatalf("something went wrong: %v\n", err)
	}
	defer db.Close()

	dbQueries := database.New(db)

	s.db = dbQueries
	s.ctx = context.Background()

	cmds.register(cmd.name, handleLogin)
	cmds.register(cmd.name, handleRegister)

	if err := cmds.run(&s, cmd); err != nil {
		log.Fatal(err)
	}
}
