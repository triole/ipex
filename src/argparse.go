package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/alecthomas/kong"
)

var (
	BUILDTAGS      string
	appName        = "ipex"
	appDescription = "IP expander, transform strings into lists of IP addresses"
	appMainversion = "0.1"
)

var CLI struct {
	Input       []string `help:"input strings, i.e. 192.168.33.1, 192.168.33.1/29, 33.1/30, 1/28, 1, 33+12, 22.5-10" arg:"" optional:"" passthrough:""`
	BaseIP      string   `help:"base IP, default is self address" optional:"" short:"b"`
	Sort        bool     `help:"sort final list before output" optional:"" short:"s"`
	Uniq        bool     `help:"uniq, remove duplicates from the output" optional:"" short:"u"`
	Separator   string   `help:"separator between printed output addresses" optional:"" default:"\n" short:"p"`
	VersionFlag bool     `help:"display version" short:"V"`
}

func parseArgs() {
	ctx := kong.Parse(&CLI,
		kong.Name(appName),
		kong.Description(appDescription),
		kong.UsageOnError(),
		kong.ConfigureHelp(kong.HelpOptions{
			Compact: true,
			Summary: true,
		}),
		kong.Vars{},
	)
	_ = ctx.Run()

	if CLI.VersionFlag {
		printBuildTags(BUILDTAGS)
		os.Exit(0)
	}
	// ctx.FatalIfErrorf(err)
}

func printBuildTags(buildtags string) {
	regexp, _ := regexp.Compile(`({|}|,)`)
	s := regexp.ReplaceAllString(buildtags, "\n")
	s = strings.Replace(s, "_subversion: ", "Version: "+appMainversion+".", -1)
	fmt.Printf("%s\n", s)
}
