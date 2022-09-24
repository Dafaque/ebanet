package scanner

import (
	"fmt"
	"log"
	"net"
	"sync"
	"sync/atomic"
	"time"
)

const maxPending uint32 = 10

var mu sync.Mutex
var pending uint32

func dialSsh(addr string) {
	if pending > maxPending {
		mu.Lock()
		defer mu.Unlock()
	}
	atomic.AddUint32(&pending, 1)
	log.Printf("checking %s:22", addr)
	conn, err := net.DialTimeout(
		"tcp",
		fmt.Sprintf("%s:22", addr),
		10*time.Second,
	)
	if err == nil {
		conn.Close()
		log.Printf("%s listens on 22 port", addr)
	}
	atomic.AddUint32(&pending, ^uint32(0))
}
