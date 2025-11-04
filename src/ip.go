package main

import (
	"fmt"
	"log"
	"net"
	"strings"
)

func getLocalIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP
}

func removeTail(ip string, x int) (r string) {
	arr := strings.Split(ip, ".")
	if len(arr) > 0 && x > 0 {
		r = strings.Join(arr[:len(arr)-x], ".")
	}
	return
}

func replaceTail(ip, newTail string) (r string) {
	newTail = strings.Trim(newTail, ".")
	x := strings.Count(newTail, ".")
	r = removeTail(ip, x+1) + "." + newTail
	return
}

func displayIPs(cidr string) {
	var ips []string

	if isIPv4Full(cidr) {
		fmt.Println(cidr)
		return
	}

	ipAddr, ipNet, err := parseCIDRString(cidr)
	if err != nil {
		log.Print(err)
		return
	}

	for ip := ipAddr.Mask(ipNet.Mask); ipNet.Contains(ip); incrementIP(ip) {
		ips = append(ips, ip.String())
	}

	if len(ips) <= 2 {
		for _, ip := range ips {
			fmt.Printf("%s\n", ip)
		}
		return
	}

	for _, ip := range ips[1 : len(ips)-1] {
		fmt.Println(ip)
	}
}

func incrementIP(ip net.IP) {
	for i := len(ip) - 1; i >= 0; i-- {
		ip[i]++
		if ip[i] != 0 {
			break
		}
	}
}

func parseCIDRString(cidr string) (ipAddr net.IP, ipNet *net.IPNet, err error) {
	ipAddr, ipNet, err = net.ParseCIDR(cidr)
	if err != nil {
		log.Print(err)
		return
	}
	return
}

func parseToIPv4(ipstr string) (ip net.IP) {
	ip = net.ParseIP(ipstr)
	ip = ip.To4()
	if ip == nil {
		log.Fatal("non ipv4 address")
	}
	return
}
