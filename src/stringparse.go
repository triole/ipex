package main

import (
	"net/netip"
)

func parseInput(args []string) (cidrs []string) {
	for _, arg := range args {
		nt := notateSplit(arg)
		nt.expandIP()

		var app []string
		switch nt.Op {
		case "":
			app = []string{nt.Exp}
		case "/":
			app = []string{nt.Exp + nt.Op + nt.Suf}
		case "+":
			ip := netip.MustParseAddr(nt.Exp)
			for i := 0; i < nt.SufInt; i++ {
				app = append(app, ip.String())
				ip = ip.Next()
			}
		case "-":
			ip := netip.MustParseAddr(nt.Exp)
			for i := 0; i < nt.SufInt; i++ {
				app = append(app, ip.String())
				ip = ip.Prev()
			}
		}
		if len(app) > 0 {
			cidrs = append(cidrs, app...)
		}
	}
	return cidrs
}
