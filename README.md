# DNSBL

dnsbl service that can be used to check DNSBL records for an IPv4 address
dnsbl offers a GraphQL API for checking, and enqueing checks on IP addresses against a DNSBL service.

### Running dnsbl
To run dnsbl locally, you can use docker-compose.yml to get started.
You can also run the binary locally using the following flow:
Follow instructions here to install mage https://magefile.org/ 

```zsh
mage -v go:build
./bin/dnsbl
```

Configuration can be provided either by environment variables, or CLI flags.
See list of configuration using `-h` on the binary.

### External Libraries Used

- `entgo.io/ent v0.9.2-0.20210821141344-368a8f7a2e9a` -- ent is an entity relationship library
- `entgo.io/contrib v0.2.0` -- ent-contrib provides ent and gqlgen templates for codegen
- `github.com/hashicorp/go-multierror v1.1.1` -- used in entc generated code 
- `github.com/vektah/gqlparser/v2 v2.2.0` -- used in entc generated code for pagination
- `github.com/vmihailenco/msgpack/v5 v5.3.5` --used in entc generated code for pagination
- `github.com/99designs/gqlgen v0.14.0` --  gqlgen is for resolver generation, and gql service scaffolding 
- `github.com/elithrar/simple-scrypt v1.3.0` -- Generate a scrypt derived key with secure salt 
- `github.com/google/uuid v1.3.0` -- used to generate new UUID-based IDs
- `github.com/gammazero/workerpool v1.1.2` -- workerpool is an in-memory woker pool using goroutines
- `github.com/jessevdk/go-flags v1.5.0` -- go-flags is used for flag and argument parsing on the CLI
- `github.com/magefile/mage v1.11.0` -- Mage is for make for go
- `github.com/mattn/go-sqlite3 v1.14.9` -- sqlite3 for portable database, and in-memory testing
- `github.com/onsi/gomega v1.17.0` -- assertions based testing library
- `go.uber.org/zap v1.19.1` -- typed logging library 

### Internal Packages

- `app/auth` -- authentication for GraphQL service
- `app/cmd` -- cmd line options
- `app/dnsbl` -- DNSBL service
- `app/ent` -- entity framework with gqlgen integration
- `app/ent/schema` -- entity schema definitions
- `app/ent/mixin` -- entity mixins
- `app/server` -- gqlgen-based GraphQL generation and server
- `app/worker` -- background worker library
