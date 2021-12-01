package dnsbl

import (
	"context"
	"fmt"
	"net"
	"strings"
)

type haus struct {
	network  string
	resolver resolver
}

type options struct {
	network    string
	dnsServer  string
	goResolver bool
}

type Option func(*options)

type resolver interface {
	LookupIP(context.Context, string, string) ([]net.IP, error)
}

func WithDNSServer(server string) Option {
	return func(o *options) {
		o.dnsServer = server
	}
}

func NewSpamhaus(opts ...Option) DNSBL {
	options := &options{}
	for _, opt := range opts {
		opt(options)
	}
	if options.network == "" {
		options.network = "ip"
	}

	r := &net.Resolver{
		PreferGo: options.goResolver,
	}
	if options.dnsServer != "" {
		r.Dial = customDialer(options.network, options.dnsServer)
	}

	return &haus{
		network:  options.network,
		resolver: r,
	}
}

func (sh *haus) Query(ctx context.Context, ip string) (*Response, error) {
	ip = ipReversal(ip)
	ip = fmt.Sprintf("%s.zen.spamhaus.org", ip)

	resp, err := sh.resolver.LookupIP(ctx, sh.network, ip)
	if err != nil {
		return nil, fmt.Errorf("spamhaus query: %w", err)
	}

	r := freshResponse()
	for _, ip := range resp {
		r.ResponseCodes = append(r.ResponseCodes, ResponseCode(ip.String()))
	}

	return r, nil
}

type dialerFunc func(ctx context.Context, network, addr string) (net.Conn, error)

func customDialer(network, server string) dialerFunc {
	return func(ctx context.Context, network, addr string) (net.Conn, error) {
		return net.Dial(network, server)
	}
}

func ipReversal(str string) string {
	octs := strings.Split(str, ".")

	last := len(octs) - 1
	for i := 0; i < len(octs)/2; i++ {
		octs[i], octs[last-i] = octs[last-i], octs[i]
	}

	return strings.Join(octs, ".")
}
