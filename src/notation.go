package main

import (
	"strconv"
	"strings"
)

type tNotation struct {
	Exp    string
	Pre    string
	Op     string
	Suf    string
	SufInt int
}

func notate(str string) (nt tNotation) {
	nt.Exp = str
	arr := rxSplit(`[\+|\-|\/]`, str)
	if len(arr) > 0 {
		nt.Pre = arr[0]
	}
	if len(arr) > 1 {
		nt.Suf = arr[1]
	}
	nt.SufInt, _ = strconv.Atoi(strings.Trim(nt.Suf, "."))
	nt.Op = rxFind(`[\+|\-|\/]`, str)

	if isIPv4Full(nt.Pre) {
		nt.Exp = nt.Pre
	} else {
		localIP = getLocalIP().String()
		nt.Exp = replaceTail(localIP, nt.Pre)
	}

	return
}
