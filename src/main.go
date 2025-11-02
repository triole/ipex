package main

import (
	"os"
)

func main() {
	cidrs := parseArgs(os.Args[1:])
	for _, cidr := range cidrs {
		displayIPs(cidr)
	}
}
