package corednsmysql

import (
	"context"
	"fmt"
	"github.com/coredns/coredns/plugin"

	"github.com/miekg/dns"
)

type DnsResolver struct {
	Next plugin.Handler
}

type ResponsePrinter struct {
	dns.ResponseWriter
}

func NewResponsePrinter(w dns.ResponseWriter) *ResponsePrinter {
	return &ResponsePrinter{ResponseWriter: w}
}

func (r *ResponsePrinter) WriteMsg(res *dns.Msg) error {
	return r.ResponseWriter.WriteMsg(res)
}

func (resolver DnsResolver) ServeDNS(ctx context.Context, w dns.ResponseWriter, r *dns.Msg) (int, error){
	pw := NewResponsePrinter(w)

	fmt.Println("xd")
	return plugin.NextOrFailure(resolver.Name(), resolver.Next, ctx, pw, r)
}

func (resolver DnsResolver) Name() string {
	return "coredns-mysql"
}