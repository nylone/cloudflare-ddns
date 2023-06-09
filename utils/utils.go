package utils

import (
	"errors"
	"fmt"
	"net"
	"strconv"
)

func FindOwnInterfaceIP(router string, netmaskLen uint8) (string, error) {
	// extract current netmask from the router ip
	_, network, err := net.ParseCIDR(router + "/" + strconv.FormatInt(int64(netmaskLen), 10))
	if err != nil {
		return "", err
	}

	// get list of available addresses
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	for _, addr := range addrs {
		if ip, ok := addr.(*net.IPNet); ok && network.Contains(ip.IP) {
			return ip.IP.String(), nil
		}
	}

	return "", errors.New("no ip was found for router address: " + router + "/" + strconv.FormatInt(int64(netmaskLen), 10))
}
