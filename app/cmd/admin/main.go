package admin

import (
	"context"
	"fmt"
	"log"

	gflags "github.com/jessevdk/go-flags"
	flags "github.com/sujithshajee/dnsbl/app"
	"github.com/sujithshajee/dnsbl/app/auth"
	"github.com/sujithshajee/dnsbl/app/ent"
	"github.com/sujithshajee/dnsbl/app/ent/migrate"
)

type cmd struct{}
type generateUserCmd struct {
	flags.Database
}

func Register(p *gflags.Parser) {
	c, err := p.AddCommand("run", "trigger command", "", &cmd{})
	if err != nil {
		log.Fatal(err)
	}

	_, err = c.AddCommand("generate-user", "", "create a new user", &generateUserCmd{})
	if err != nil {
		log.Fatal(err)
	}
}

func (c *cmd) Run(args []string) error {
	return fmt.Errorf("no subcommand provided")
}

func (c *generateUserCmd) Run(args []string) error {
	if len(args) != 2 {
		return fmt.Errorf("%d missing arguments: required 2", len(args))
	}
	ctx := context.Background()

	cl, err := ent.Open(c.DatabaseDriver, c.DatabaseURI)
	if err != nil {
		return fmt.Errorf("opening database: %w", err)
	}
	err = cl.Schema.Create(ctx, migrate.WithDropColumn(true), migrate.WithDropIndex(true))
	if err != nil {
		return fmt.Errorf("migration: %w", err)
	}

	return auth.GenerateUser(ctx, cl, args[0], args[1])
}
