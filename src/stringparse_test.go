package main

import (
	"testing"
)

func TestParseArgs(t *testing.T) {
	localIP := getLocalIP().String()
	assertParseArgs(
		[]string{"1"},
		[]string{replaceTail(localIP, ".1")},
		t,
	)
	assertParseArgs(
		[]string{".1"},
		[]string{replaceTail(localIP, "1")},
		t,
	)
	assertParseArgs(
		[]string{"22.1"},
		[]string{replaceTail(localIP, "22.1")},
		t,
	)
	assertParseArgs(
		[]string{".22.1"},
		[]string{replaceTail(localIP, ".22.1")},
		t,
	)

}

func assertParseArgs(inp, exp []string, t *testing.T) {
	res := parseInput(inp)
	if len(res) != len(exp) {
		t.Errorf(
			"result and expecation differ in length: %s -> %s != %s",
			inp, exp, res,
		)
	}
	for i := 1; i <= len(res)-1; i++ {
		if exp[i] != res[i] {
			t.Errorf(
				"result and expectation differ: %s -> %s != %s",
				inp, exp, res,
			)
		}
	}
}
