package main

import (
	"regexp"
	"strconv"
	"strings"
)

const (
	rxIPv4Full     = `^(?:\d{1,3}\.){3}\d{1,3}$`
	rxIPv4Fragment = `^(?:\d{1,3}\.)+\d{1,3}$`
)

func cidrNotateSplit(str string) (arr []string) {
	if isCIDRNotation(str) {
		arr = strings.Split(str, "/")
	} else {
		arr = []string{str, "32"}
	}
	return
}

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
	match, _ := regexp.MatchString(
		rxIPv4Full, strings.Trim(str, "."),
	)
	return match
}

func isIPv4Fragment(str string) bool {
	match, _ := regexp.MatchString(
		rxIPv4Fragment, strings.Trim(str, "."),
	)
	return match
}
