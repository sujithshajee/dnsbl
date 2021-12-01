package graphql

import (
	"log"

	"github.com/jessevdk/go-flags"
	// "github.com/sujithshajee/dnsbl/app/graphql"
)

type cmd struct{}

func Register(p *flags.Parser) {
	c, err := p.AddCommand("graphql", "run graphql", "", &cmd{})
	if err != nil {
		log.Fatalln(err)
	}

	_, err = c.AddGroup("server", "server options", &cmd{})
	if err != nil {
		log.Fatalln(err)
	}
}

// Execute creates and starts a new GraphQL service
// func (*cmd) Execute(args []string) error {
// 	s, err := graphql.New()
// 	if err != nil {
// 		log.Fatalf("unable to start server: %s\n", err)
// 	}

// 	return s.Start()
// }
