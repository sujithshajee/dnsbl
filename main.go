package main

import (
	"os"

	"github.com/jessevdk/go-flags"
	"github.com/sujithshajee/dnsbl/app/cmd/admin"
	"github.com/sujithshajee/dnsbl/app/cmd/graphql"
	"github.com/sujithshajee/dnsbl/app/cmd/worker"
)

func main() {
	p := flags.NewParser(nil, flags.Default)
	p.SubcommandsOptional = true

	worker.Register(p)
	graphql.Register(p)
	admin.Register(p)

	if _, err := p.Parse(); err != nil {
		os.Exit(1)
	}
}
