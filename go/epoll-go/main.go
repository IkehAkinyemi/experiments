package main

import (
	"fmt"
	"log"
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

	file, err := ln.(*net.TCPListener).File()
	if err != nil {
		log.Fatal(err)
	}
	fd := file.Fd()

	// Add the listening socket to the epoll interest list
	event := syscall.EpollEvent{
		Events: syscall.EPOLLIN,
		Fd:     int32(fd),
	}
	if err := syscall.EpollCtl(epollFD, syscall.EPOLL_CTL_ADD, int(fd), &event); err != nil {
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
			if int(events[i].Fd) == int(fd) {
				// Handle incoming connection
				_, err := ln.Accept()
				if err != nil {
					fmt.Println("Error accepting connection:", err)
					continue
				}

				// Add the new client socket to the epoll interest list
				clientFD := fd
				event := syscall.EpollEvent{
					Events: uint32(syscall.EPOLLIN),
					Fd:     int32(clientFD),
				}
				if err := syscall.EpollCtl(epollFD, syscall.EPOLL_CTL_ADD, int(clientFD), &event); err != nil {
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