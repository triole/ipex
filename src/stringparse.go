package main

import (
	"fmt"
)

func parseInput(args []string) (cidrs []string) {
	for _, arg := range args {
		nt := notateSplit(arg)

		app := nt.Pre
		switch nt.Op {
		case "/", "":
			if isIPv4Full(nt.Pre) {
				app = nt.Pre
			} else if isIPv4Fragment(nt.Pre) || isInt(nt.Pre) {
				if !isIPv4Full(localIP) {
					localIP = getLocalIP().String()
				}
				app = replaceTail(localIP, nt.Pre)
			}
			if nt.Suf != "" {
				app += nt.Op + nt.Suf
			}
		case "+":
			fmt.Println("Tomorrow.")
		case "-":
			fmt.Println("In two days.")
		}

		if app != "" {
			cidrs = append(cidrs, app)
		}
	}
	return cidrs
}
