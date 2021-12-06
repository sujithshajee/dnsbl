package worker

import (
	"fmt"
	"log"

	gflags "github.com/jessevdk/go-flags"
	"github.com/sujithshajee/dnsbl/app/dnsbl"
	"github.com/sujithshajee/dnsbl/app/ent"
	"github.com/sujithshajee/dnsbl/app/ent/task"
	"github.com/sujithshajee/dnsbl/app/worker"
	backgroundjob "github.com/sujithshajee/dnsbl/app/worker/backgroundjobs/dnsbl"
	"go.uber.org/zap"
)

type cmd struct{}

func Register(p *gflags.Parser) {
	_, err := p.AddCommand("worker", "worker service", "", &cmd{})
	if err != nil {
		log.Fatal(err)
	}
}

func (c *cmd) Execute(args []string) error {
	cl, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
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
