package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const (
	rxIPv4Full     = `^(?:\d{1,3}\.){3}\d{1,3}$`
	rxIPv4Fragment = `^(?:\d{1,3}\.)+\d{1,3}$`
)

func isCIDRNotation(str string) bool {
	return strings.Contains(strings.Trim(str, "."), "/")
}

func isInt(str string) bool {
	if _, err := strconv.Atoi(strings.Trim(str, ".")); err == nil {
		return true
	}
	return false
}

func isIPv4Full(str string) bool {
	return rxMatch(rxIPv4Full, strings.Trim(str, "."))
}

func isIPv4Fragment(str string) bool {
	return rxMatch(rxIPv4Fragment, strings.Trim(str, "."))
}

func rxCompile(str string) (r *regexp.Regexp) {
	var err error
	r, err = regexp.Compile(str)
	if err != nil {
		fmt.Printf("can not compile regex: %s\n", err)
		os.Exit(1)
	}
	return
}

func rxSplit(rx, str string) []string {
	re := regexp.MustCompile(rx)
	return re.Split(str, -1)
}

func rxFind(rx string, content string) (r string) {
	temp := rxCompile(rx)
	r = temp.FindString(content)
	return
}

func rxMatch(rx string, str string) (b bool) {
	re := rxCompile(rx)
	b = re.MatchString(str)
	return
}
