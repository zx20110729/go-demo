package main

import "fmt"

type IpAddr [4]byte

func (ipAddr IpAddr) String() string {
	l := len(ipAddr)
	var result string
	for i := 0; i < l; i++ {
		result = fmt.Sprintf("%v%v", result, ipAddr[i])
		if i < l-1 {
			result = result + "."
		}
	}
	return result
}
func main() {
	hosts := map[string]IpAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v , %v\n", name, ip)
	}
}
