package server

import (
	"context"
	"fmt"
	"net/http"

	// database support
	"entgo.io/contrib/entgql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
	"github.com/sujithshajee/dnsbl/app/auth"
	"github.com/sujithshajee/dnsbl/app/ent"
	"github.com/sujithshajee/dnsbl/app/ent/migrate"
	"github.com/sujithshajee/dnsbl/app/server/graph/resolver"
	"go.uber.org/zap"
)

// GraphQLFlags are the flags for the GraphQL HTTP service
type GraphQLFlags struct {
	Port   string `short:"p" long:"port" env:"PORT" description:"Port to run the server on, listens on 0.0.0.0:PORT"`
	Listen string `short:"l" long:"listen" env:"LISTEN" description:"Listen address to run the server on" default:":8080"`
	Debug  bool   `long:"debug" env:"DEBUG" description:"If the server should be in debug mode"`
}

// Flags are the configurable set of GraphQLFlags
var Flags GraphQLFlags

type options struct {
	logger *zap.SugaredLogger
	client *ent.Client
}

// Option is optional configuration for the GraphQL server
type Option func(*options)

// WithLogger returns an Option to configure a logger
func WithLogger(l *zap.SugaredLogger) Option {
	return func(o *options) {
		o.logger = l
	}
}

// WithDatabase returns an Option to configure a database client connection to
// use for the GraphQL server
func WithDatabase(cl *ent.Client) Option {
	return func(o *options) {
		o.client = cl
	}
}

// Server represents the HTTP server that handles GraphQL requests
type Server struct {
	logger  *zap.SugaredLogger
	client  *ent.Client
	handler http.Handler
}

// Start starts the Server, and blocks
func (s *Server) Start() error {
	if Flags.Port != "" {
		Flags.Listen = fmt.Sprintf(":%s", Flags.Port)
	} else {
		Flags.Listen = fmt.Sprintf(":%s", "8080")
	}

	s.logger.Infof("listening on %s", Flags.Listen)
	return http.ListenAndServe(Flags.Listen, s.handler)
}

// Stop stops the Server
func (s *Server) Stop() error {
	return s.client.Close()
}

func New(opts ...Option) (*Server, error) {
	options := &options{}
	for _, opt := range opts {
		opt(options)
	}

	if options.logger == nil {
		l, err := zap.NewProduction()
		if err != nil {
			return nil, fmt.Errorf("unable to start logger: %w", err)
		}
		options.logger = l.Sugar()
	}

	if options.client == nil {
		entOpts := []ent.Option{
			ent.Log(func(vs ...interface{}) {
				if len(vs) == 0 {
					return
				}

				msg := vs[0]
				options.logger.Debugf(msg.(string), vs[1:]...)
			}),
		}

		if Flags.Debug {
			entOpts = append(entOpts, ent.Debug())
		}

		cl, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1", entOpts...)
		if err != nil {
			return nil, fmt.Errorf("unable to open database connetion: %w", err)
		}
		options.client = cl
	}

	options.logger.Info("migrating schema")
	ctx := context.Background()
	err := options.client.Schema.Create(ctx, migrate.WithDropColumn(true), migrate.WithDropIndex(true))
	if err != nil {
		return nil, fmt.Errorf("migration: %w", err)
	}

	cl, er := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if er != nil {
		return nil, fmt.Errorf("unable to open database connetion: %w", err)
	}
	auth.GenerateUser(ctx, cl, "secureworks", "supersecret")

	r := mux.NewRouter()

	srv := handler.NewDefaultServer(resolver.NewSchema(options.client))
	srv.Use(entgql.Transactioner{TxOpener: options.client})

	r.Handle("/", http.RedirectHandler("/graphql/playground", http.StatusPermanentRedirect))
	r.Handle("/graphql/playground", playground.Handler("GraphQL Playground", "/graphql"))
	r.Handle("/graphql", auth.BasicAuth(options.client, srv))

	return &Server{
		logger:  options.logger,
		client:  options.client,
		handler: r,
	}, nil
}
