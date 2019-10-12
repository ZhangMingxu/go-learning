package main

import (
	"fmt"
	"strconv"
)

type IPAddr [4]byte

// TODO: 给 IPAddr 添加一个 "String() string" 方法
func (ip IPAddr) String() string {
	s := ""
	for i, v := range ip {
		if i == len(ip)-1 {
			s += strconv.Itoa(int(v))
		} else {
			s = s + strconv.Itoa(int(v)) + "."
		}
	}
	return s
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
