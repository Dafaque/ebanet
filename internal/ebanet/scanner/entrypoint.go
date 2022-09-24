package scanner

import (
	"net"
	"net/netip"
)

func Run() {
	var ranges []net.Addr
	ranges = internalAddrs()
	ranges = append(ranges, externalAddrs()...)
	for _, cidr := range ranges {
		prefix, err := netip.ParsePrefix(cidr.String())
		if err != nil {
			continue
		}
		if !prefix.Addr().Is4() { //TMP not use v6 net
			continue
		}
		var addr netip.Addr
		for addr = prefix.Addr(); prefix.Contains(addr); addr = addr.Next() {
			ip := addr.String()
			go dialSsh(ip)
		}
	}
}
