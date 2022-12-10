package udp

import (
	"bytes"
	"context"
	"net"
	"testing"
)

func TestUDPServer(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	serverAddr, err := UdpServer(ctx, "127.0.0.1:")
	if err != nil {
		t.Fatal(err)
	}
	
	client, err := net.ListenPacket("udp", "127.0.0.1:")
	if err != nil {
		t.Fatal(err)
	}
	defer client.Close()

	msg := []byte("ping")
	_, err = client.WriteTo(msg, serverAddr)
	if err != nil {
		t.Fatal(err)
	}

	buf := make([]byte, 1024)
	n, addr, err := client.ReadFrom(buf)
	if err != nil {
		t.Fatal(err)
	}

	if addr.String() != serverAddr.String() {
		t.Fatalf("expected %s; actual %s", serverAddr, addr)
	}

	if !bytes.Equal(msg, buf[:n]) {
		t.Errorf("expected %s; actual %s", msg, string(buf[:n]))
	}
}