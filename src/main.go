package main

import (
	"fmt"
	"net/netip"
	"os"
	"sort"

	"golang.org/x/term"
)

var (
	localIP string
)

func main() {
	parseArgs()
	if CLI.BaseIP != "" {
		localIP = CLI.BaseIP
	}

	ipList := parseInput(CLI.Input)

	if CLI.Uniq {
		ipList = removeDuplicates(ipList)
	}

	if CLI.Sort {
		sort.Slice(ipList, func(i, j int) bool {
			return ipList[i].Less(ipList[j])
		})
	}

	for i, el := range ipList {
		if i+1 < len(ipList) {
			fmt.Printf("%s%s", el.String(), CLI.Separator)
		} else {
			fmt.Printf("%s", el.String())
		}
	}
	if term.IsTerminal(int(os.Stdout.Fd())) {
		fmt.Printf("\n")
	}
}

func removeDuplicates(ipList []netip.Addr) (newList []netip.Addr) {
	allKeys := make(map[netip.Addr]bool)
	for _, item := range ipList {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			newList = append(newList, item)
		}
	}
	return newList
}
