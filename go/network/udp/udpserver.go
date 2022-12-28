package udp

import (
	"context"
	"fmt"
	"net"
)

func UdpServer(ctx context.Context, addr string) (net.Addr, error) {
	s, err := net.ListenPacket("udp", addr)
	if err != nil {
		return nil, fmt.Errorf("binding udp %s: %w", addr, err)
	}

	go func() {
		go func() {
			<-ctx.Done()
			s.Close()
		}()

		buf := make([]byte, 1024)

		for {
			n, clientAdrr, err := s.ReadFrom(buf)
			if err != nil {
				return
			}

			_, err = s.WriteTo(buf[:n], clientAdrr)
			if err != nil {
				return
			}
		}
	}()

	return s.LocalAddr(), nil
}
