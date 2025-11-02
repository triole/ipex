package main

func parseInput(args []string) (cidrs []string) {
	for _, arg := range args {
		cidrSplit := cidrNotateSplit(arg)
		app := ""

		if isIPv4Full(cidrSplit[0]) {
			app = cidrSplit[0]
		} else if isIPv4Fragment(cidrSplit[0]) || isInt(cidrSplit[0]) {
			if !isIPv4Full(localIP) {
				localIP = getLocalIP().String()
			}
			app = replaceTail(localIP, cidrSplit[0])
		}
		if app != "" {
			cidrs = append(cidrs, app+"/"+cidrSplit[1])
		}
	}
	return cidrs
}
