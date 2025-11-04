package main

import (
	"net/netip"
)

func parseInput(args []string) (cidrs []netip.Addr) {
	for _, arg := range args {
		nt := notate(arg)

		var app []netip.Addr
		switch nt.Op {
		case "":
			app = []netip.Addr{netip.MustParseAddr(nt.Exp)}
		case "/":
			for _, el := range parseCIDRtoIPList(nt.Exp + nt.Op + nt.Suf) {
				app = append(app, el)
			}
		case "+":
			ip := netip.MustParseAddr(nt.Exp)
			for i := 0; i < nt.SufInt; i++ {
				app = append(app, ip)
				ip = ip.Next()
			}
		case "-":
			ip := netip.MustParseAddr(nt.Exp)
			for i := 0; i < nt.SufInt; i++ {
				app = append(app, ip)
				ip = ip.Prev()
			}
		}
		if len(app) > 0 {
			cidrs = append(cidrs, app...)
		}
	}
	return cidrs
}
