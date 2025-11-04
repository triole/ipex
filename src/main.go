package main

import (
	"net/netip"
	"sort"
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

	for _, el := range ipList {
		displayIPs(el.String())
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
