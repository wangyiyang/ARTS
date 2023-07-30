package main

import (
	"fmt"
	"strconv"
	"strings"
)

func validIPAddress(IP string) string {
	if checkIPv4(IP) {
		return "IPv4"
	}
	if checkIPv6(IP) {
		return "IPv6"
	}
	return "Neither"
}

func checkIPv4(IP string) bool {
	strs := strings.Split(IP, ".")
	if len(strs) != 4 {
		return false
	}
	for _, s := range strs {
		if len(s) == 0 || (len(s) > 1 && s[0] == '0') {
			return false
		}
		// if s[0] < '0' || s[0] > '9' {
		// 	return false
		// }
		n, err := strconv.Atoi(s)
		if err != nil {
			return false
		}
		if n < 0 || n > 255 {
			return false
		}
	}
	return true
}

func checkIPv6(IP string) bool {
	strs := strings.Split(IP, ":")
	if len(strs) != 8 {
		return false
	}
	for _, s := range strs {
		if len(s) <= 0 || len(s) > 4 {
			return false
		}
		for i := 0; i < len(s); i++ {
			if s[i] >= '0' && s[i] <= '9' {
				continue
			}
			if s[i] >= 'A' && s[i] <= 'F' {
				continue
			}
			if s[i] >= 'a' && s[i] <= 'f' {
				continue
			}
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(validIPAddress("172.16.254.1"))
}
