package tcp

import (
	"io"
	"log"
	"net"
	"os"
)

type Monitor struct {
	*log.Logger
}

func (mlog *Monitor) Write(p []byte) (int, error) {
	return len(p), mlog.Output(2, string(p))
}

func ExampleMonitor() {
	monitor := &Monitor{Logger: log.New(os.Stdout, "monitor: ", 0)}

	listener, err := net.Listen("tcp", "127.0.0.1:")
	if err != nil {
		monitor.Fatal(err)
	}

	done := make(chan struct{})

	go func() {
		defer close(done)

		conn, err := listener.Accept()
		if err != nil {
			return
		}
		defer conn.Close()

		buf := make([]byte, 1024)

		r := io.TeeReader(conn, monitor)
		n, err := r.Read(buf)
		if err != nil {
			monitor.Println(err)
			return
		}

		w := io.MultiWriter(conn, monitor)
		_, err = w.Write(buf[:n])
		if err != nil && err != io.EOF {
			monitor.Println(err)
			return
		}
	}()

	conn, err := net.Dial("tcp", listener.Addr().String())
	if err != nil {
		monitor.Fatal(err)
	}

	_, err = conn.Write([]byte("Ping\n"))
	if err != nil {
		monitor.Fatal(err)
	}

	conn.Close()
	<-done
}