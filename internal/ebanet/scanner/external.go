package scanner

import (
	"log"
	"net"
)

func externalAddrs() []net.Addr {
	defer func() { //Code will panic, cause we do not give a fuck; but recovers, cause we do not want to stop
		if rec := recover(); rec != nil {
			log.Println(rec)
		}
	}()
	var addrs []net.Addr
	// @todo
	// Open file .cidr from assets/:
	// file, err := os.ReadFile(path)

	// Read line by line (separator \n)
	// Parse line as ip and append to result:
	// addrs = append(addrs, &net.IPAddr{IP: net.ParseIP(string)})

	// tut net proverok na oshibki i mb vashe ne rabotaet, need to check
	return addrs
}
