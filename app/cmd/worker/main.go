package worker

import (
	"fmt"
	"log"

	gflags "github.com/jessevdk/go-flags"
	flags "github.com/sujithshajee/dnsbl/app"
	"github.com/sujithshajee/dnsbl/app/dnsbl"
	"github.com/sujithshajee/dnsbl/app/ent"
	"github.com/sujithshajee/dnsbl/app/ent/task"
	"github.com/sujithshajee/dnsbl/app/worker"
	backgroundjob "github.com/sujithshajee/dnsbl/app/worker/backgroundjobs/dnsbl"
	"go.uber.org/zap"
)

type cmd struct {
	flags.Database
}

// Register registers the worker runner with the flags parser given
func Register(p *gflags.Parser) {
	_, err := p.AddCommand("worker", "worker service", "", &cmd{})
	if err != nil {
		log.Fatal(err)
	}
}

// Execute initializes and starts a new worker pool
func (c *cmd) Execute(args []string) error {
	cl, err := ent.Open(c.DatabaseDriver, c.DatabaseURI)
	if err != nil {
		return fmt.Errorf("opening database: %w", err)
	}

	log, err := zap.NewProduction()
	if err != nil {
		return fmt.Errorf("getting logger: %w", err)
	}

	wp := worker.New(cl, log.Sugar())
	if err != nil {
		return fmt.Errorf("initializing worker: %w", err)
	}

	sh := dnsbl.NewSpamhaus()
	j := backgroundjob.NewDNSBLJob(cl, sh)

	wp.Register(task.TypeIPDNSBL, j)

	return wp.Start()
}
