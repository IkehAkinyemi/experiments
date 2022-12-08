package tcp

import (
	"log"
	"net"
	"time"
)

var (
	n int // bytes written
	maxTrial = 7 // maximum retries
)

//Mock a retries for adverse network
func Retries() {
	listener, err := net.Listen("tcp", "::1")
	if err != nil {
		return // report err instead
	}
	conn, err := listener.Accept()
	if err != nil {
		return // report err instead
	}

	for ; maxTrial > 0; maxTrial-- {
		n, err = conn.Write([]byte("Mission"))
		if err != nil {
			if err, ok := err.(net.Error); ok && err.Temporary() {
				log.Println(err)
				time.Sleep(5*time.Second)
				continue
			}

			return // report err 
		}
		break
	}

	if maxTrial == 0 {
		return // Give up, I guess
	}

	log.Printf("wrote %d bytes to %s\n", n, conn.RemoteAddr())
}