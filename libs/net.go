package fynelibs

import (
	"fmt"

	"context"
	"net"
	"time"
)

func makeResolver(resolver_addr string, resolver_port int) *net.Resolver {
	if resolver_addr == "" {
		resolver_addr = "8.8.8.8"
	}

	ResolverFullAddr := fmt.Sprintf("%v:%d", resolver_addr, resolver_port)

	DnsResolver := &net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			d := net.Dialer{
				Timeout: time.Millisecond * time.Duration(10000),
			}
			return d.DialContext(ctx, network, ResolverFullAddr)
		},
	}
	return DnsResolver
}

func LookupMyIp() ([]string, error) {
	OpendnsReverseResolver := makeResolver("resolver1.opendns.com", 53)

	ip, err := OpendnsReverseResolver.LookupHost(context.Background(), "myip.opendns.com")

	return ip, err
}