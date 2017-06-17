package methods

import "fmt"

// IPAddr : Type for ipaddr
type IPAddr [4]byte

// Stringer : to format string output by print
func (ip IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", ip[0], ip[1], ip[2], ip[3])
}

// PrintExStringer : Demo Exercise
func PrintExStringer() {
	hosts := map[string]IPAddr{
		"googleDNS": {8, 8, 8, 8},
		"localhost": {127, 0, 0, 1},
	}

	for name, ipaddr := range hosts {
		fmt.Printf("%v: %v\n", name, ipaddr)
	}
}
