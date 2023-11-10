package network

import "net"

func GetHostAddress(interfaceName string) ([]string, error) {
	var ipList []string

	interfaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	for _, iface := range interfaces {
		if (iface.Flags&net.FlagUp) == 0 || (iface.Flags&net.FlagLoopback) != 0 {
			continue
		}

		if interfaceName != "" && interfaceName != iface.Name {
			continue
		}

		addrs, err := iface.Addrs()
		if err != nil {
			return nil, err
		}

		for _, addr := range addrs {
			ip, _, err := net.ParseCIDR(addr.String())
			if err != nil {
				return nil, err
			}

			if ip.To4() == nil {
				// Skip IPv6 addresses
				continue
			}

			ipList = append(ipList, ip.String())
		}
	}

	return ipList, nil
}
