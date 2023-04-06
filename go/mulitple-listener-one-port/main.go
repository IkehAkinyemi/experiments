package main

import (
	"context"
	"log"
	"net"
	"syscall"

	"golang.org/x/sys/unix"
)

func main() {
 // Create Listener Config
 lc := net.ListenConfig{
  Control: func(network, address string, c syscall.RawConn) error {
   return c.Control(func(fd uintptr) {
    // Enable SO_REUSEADDR
    err := unix.SetsockoptInt(int(fd), unix.SOL_SOCKET, unix.SO_REUSEADDR, 1)
    if err != nil {
     log.Printf("Could not set SO_REUSEADDR socket option: %s", err)
    }

    // Enable SO_REUSEPORT
    err = unix.SetsockoptInt(int(fd), unix.SOL_SOCKET, unix.SO_REUSEPORT, 1)
    if err != nil {
     log.Printf("Could not set SO_REUSEPORT socket option: %s", err)
    }
   })
  },
 }

 // Start Listener
 l, err := lc.Listen(context.Background(), "tcp", "0.0.0.0:9000")
 if err != nil {
  log.Printf("Could not start TCP listener: %s", err)
  return
 }

 // Wait for new connections on a listener
 for {
  // Accept new connections
  c, err := l.Accept()
  if err != nil {
   log.Printf("Listener returned: %s", err)
   break
  }

  // Kickoff a Goroutine to handle the new connection
  go func() {
   defer c.Close()
   log.Printf("New connection created")

   // Write a hello world and close the session
   _, err := c.Write([]byte("Hello World"))
   if err != nil {
    log.Printf("Unable to write on connection: %s", err)
   }
  }()
 }
}