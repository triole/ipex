package main

import (
	"net/netip"
)

func parseInput(args []string) (cidrs []netip.Addr) {
	for _, arg := range args {
		nt := notateSplit(arg)
		nt.expandIP()

		var app []netip.Addr
		switch nt.Op {
		case "":
			app = []netip.Addr{netip.MustParseAddr(nt.Exp)}
		case "/":
			app = []netip.Addr{netip.MustParseAddr(nt.Exp + nt.Op + nt.Suf)}
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
