package scanner

import (
	"log"
	"net"
)

func internalAddrs() []net.Addr {
	defer func() {
		if rec := recover(); rec != nil {
			log.Println(rec)
		}
	}()
	var ifaces []net.Interface
	{ //Shadow err, cause ih tut budet dahooya
		var err error
		if ifaces, err = net.Interfaces(); err != nil {
			panic(err)
		}
	}
	var addrs []net.Addr = make([]net.Addr, 0)
	for _, iface := range ifaces {
		if ifAddrs, err := iface.Addrs(); err != nil {
			panic(err)
		} else {
			addrs = append(addrs, ifAddrs...)
		}
	}
	return addrs
}
