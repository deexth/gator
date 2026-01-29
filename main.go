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

	cmds := loadCommands()

	cmdLineArgs := os.Args
	if len(cmdLineArgs) < 2 {
		log.Fatal("Usage: cli <command> [args...]")
	}

	cmd := command{
		name: cmdLineArgs[1],
		args: cmdLineArgs[2:],
	}

	if _, ok := cmds.cmds[cmd.name]; !ok {
		log.Fatalf("no command %v for gator", cmd.name)
	}

	db, err := sql.Open("postgres", s.conf.DBURL)
	if err != nil {
		log.Fatalf("something went wrong: %v\n", err)
	}
	defer db.Close()

	dbQueries := database.New(db)

	s.db = dbQueries
	s.ctx = context.Background()

	if err := cmds.run(&s, cmd); err != nil {
		log.Fatal(err)
	}
}
