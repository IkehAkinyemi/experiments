package main

import (
	"fmt"
	"net"
	"os"
	"syscall"
)

func main() {
	// Create a listening socket on port 8080
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error creating server:", err)
		os.Exit(1)
	}
	defer ln.Close()

	// Create an epoll file descriptor
	epollFD, err := syscall.EpollCreate1(0)
	if err != nil {
		fmt.Println("Error creating epoll FD:", err)
		os.Exit(1)
	}

	// Add the listening socket to the epoll interest list
	event := syscall.EpollEvent{
		Events: syscall.EPOLLIN,
		Fd:     int32(ln.(*net.TCPListener).Fd()),
	}
	if err := syscall.EpollCtl(epollFD, syscall.EPOLL_CTL_ADD, int(ln.(*net.TCPListener).Fd()), &event); err != nil {
		fmt.Println("Error adding listening socket to epoll interest list:", err)
		os.Exit(1)
	}

	// Start the event loop
	events := make([]syscall.EpollEvent, 10)
	for {
		n, err := syscall.EpollWait(epollFD, events, -1)
		if err != nil {
			fmt.Println("Error in epoll wait:", err)
			continue
		}

		// Handle incoming events
		for i := 0; i < n; i++ {
			if int(events[i].Fd) == int(ln.(*net.TCPListener).Fd()) {
				// Handle incoming connection
				conn, err := ln.Accept()
				if err != nil {
					fmt.Println("Error accepting connection:", err)
					continue
				}

				// Add the new client socket to the epoll interest list
				clientFD := int(conn.(*net.TCPConn).Fd())
				event := syscall.EpollEvent{
					Events: syscall.EPOLLIN | syscall.EPOLLET,
					Fd:     int32(clientFD),
				}
				if err := syscall.EpollCtl(epollFD, syscall.EPOLL_CTL_ADD, clientFD, &event); err != nil {
					fmt.Println("Error adding client socket to epoll interest list:", err)
					continue
				}
			} else {
				// Handle incoming data on client socket
				conn, err := net.FileConn(os.NewFile(uintptr(events[i].Fd), ""))
				if err != nil {
					fmt.Println("Error converting file descriptor to net.Conn:", err)
					continue
				}
				buf := make([]byte, 1024)
				n, err := conn.Read(buf)
				if err != nil {
				
				fmt.Println("Error reading from client socket:", err)
				continue
			}
			fmt.Printf("Received data from client: %s\n", buf[:n])

			// Close the client connection
			conn.Close()

			// Remove the client socket from the epoll interest list
			if err := syscall.EpollCtl(epollFD, syscall.EPOLL_CTL_DEL, int(events[i].Fd), nil); err != nil {
				fmt.Println("Error removing client socket from epoll interest list:", err)
				continue
			}
		}
	}
}
}