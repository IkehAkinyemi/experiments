package udp

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type loggerRT struct {
	http.RoundTripper
}

func (rt *loggerRT) RoundTrip(req *http.Request) (*http.Response, error) {
	log.Printf("%v %v", req.Method, req.URL)
	return rt.RoundTripper.RoundTrip(req)
}

func Trip() {
	roundtripper := &loggerRT{
		RoundTripper: http.DefaultTransport,
	}

	client := http.Client {
		Transport: roundtripper,
	}

	buf := make([]byte, 1<<9)
	var n int

	res, err := client.Get("lighten.ikehakinyemi.net")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for {
		n, err = res.Body.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err.Error())
			break
		}
	}

	fmt.Println(buf[:n])
}