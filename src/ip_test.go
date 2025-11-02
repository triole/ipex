package main

import "testing"

func TestReplaceTail(t *testing.T) {
	assertReplaceTail("192.168.33.1", "9", "192.168.33.9", t)
	assertReplaceTail("192.168.33.1", ".9", "192.168.33.9", t)
	assertReplaceTail("192.168.33.1", "11.9", "192.168.11.9", t)
	assertReplaceTail("192.168.33.1", ".11.9", "192.168.11.9", t)
	assertReplaceTail("192.168.33.1", "22.11.9", "192.22.11.9", t)
	assertReplaceTail("192.168.33.1", ".22.11.9", "192.22.11.9", t)
}

func assertReplaceTail(ip, newTail, exp string, t *testing.T) {
	res := replaceTail(ip, newTail)
	if res != exp {
		t.Errorf("replaceTail failed: %s != %s", res, exp)
	}
}
