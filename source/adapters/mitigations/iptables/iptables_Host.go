package iptables

import "tholian-firewall/types"
import "os/exec"

func isForbiddenHost(chain string, address string) bool {

	var program string

	if types.IsIPv6(address) {

		ipv6 := types.ParseIPv6(address)

		if ipv6 != nil {
			tmp := ipv6.String()
			address = tmp[1 : len(tmp)-1]
			program = "ip6tables"
		}

	} else if types.IsIPv4(address) {

		ipv4 := types.ParseIPv4(address)

		if ipv4 != nil {
			address = ipv4.String()
			program = "iptables"
		}

	}

	var result bool = false

	if program != "" {

		if chain == "INPUT" {

			cmd := exec.Command(program, "-C", "INPUT", "-s", address, "-j", "DROP")
			_, err := cmd.Output()

			if err == nil {
				result = true
			}

		} else if chain == "OUTPUT" {

			cmd := exec.Command(program, "-C", "OUTPUT", "-d", address, "-j", "DROP")
			_, err := cmd.Output()

			if err == nil {
				result = true
			}

		}

	}

	return result

}

func forbidHost(chain string, address string) bool {

	var program string

	if types.IsIPv6(address) {

		ipv6 := types.ParseIPv6(address)

		if ipv6 != nil {
			tmp := ipv6.String()
			address = tmp[1 : len(tmp)-1]
			program = "ip6tables"
		}

	} else if types.IsIPv4(address) {

		ipv4 := types.ParseIPv4(address)

		if ipv4 != nil {
			address = ipv4.String()
			program = "iptables"
		}

	}

	var result bool = false

	if program != "" {

		if chain == "INPUT" {

			cmd := exec.Command(program, "-A", "INPUT", "-s", address, "-j", "DROP")
			_, err := cmd.Output()

			if err == nil {
				result = true
			}

		} else if chain == "OUTPUT" {

			cmd := exec.Command(program, "-A", "OUTPUT", "-d", address, "-j", "DROP")
			_, err := cmd.Output()

			if err == nil {
				result = true
			}

		}

	}

	return result

}

func permitHost(chain string, address string) bool {

	var program string

	if types.IsIPv6(address) {

		ipv6 := types.ParseIPv6(address)

		if ipv6 != nil {
			tmp := ipv6.String()
			address = tmp[1 : len(tmp)-1]
			program = "ip6tables"
		}

	} else if types.IsIPv4(address) {

		ipv4 := types.ParseIPv4(address)

		if ipv4 != nil {
			address = ipv4.String()
			program = "iptables"
		}

	}

	var result bool = false

	if program != "" {

		if chain == "INPUT" {

			cmd := exec.Command(program, "-D", "INPUT", "-s", address, "-j", "DROP")
			_, err := cmd.Output()

			if err == nil {
				result = true
			}

		} else if chain == "OUTPUT" {

			cmd := exec.Command(program, "-D", "OUTPUT", "-d", address, "-j", "DROP")
			_, err := cmd.Output()

			if err == nil {
				result = true
			}

		}

	}

	return result

}
