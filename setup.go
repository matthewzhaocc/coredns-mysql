package corednsmysql

import (
	"github.com/coredns/caddy"
	"github.com/coredns/coredns/core/dnsserver"
	"github.com/coredns/coredns/plugin"
	"fmt"
)

func init() {
	fmt.Println("initing coredns-mysql")
	plugin.Register("coredns-mysql", setup)
}

func setup(c *caddy.Controller) error {
	c.Next()
	
	if c.NextArg() {
		return plugin.Error("coredns-mysql", c.ArgErr())
	}

	dnsserver.GetConfig(c).AddPlugin(func(next plugin.Handler) plugin.Handler {
		return DnsResolver{Next: next}
	})
	return nil
}