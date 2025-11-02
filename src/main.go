package main

var (
	localIP string
)

func main() {
	parseArgs()
	if CLI.BaseIP != "" {
		localIP = CLI.BaseIP
	}
	for _, cidr := range parseInput(CLI.Input) {
		displayIPs(cidr)
	}
}
