package expose

import (
	"log"
	"io"
	"net"
)

// Forward TCP connections received on listener to the address identified by
// forwardTo.  Call this in a goroutine.  It will stop when the listener stops.
func forwardPort(forwardFrom net.Listener, forwardTo string) {
	log.Printf("in forwardPort")
	for {
		log.Printf("in forwardPort, calling accept")
		fromConn, err := forwardFrom.Accept()
		if err != nil {
		log.Printf("in forwardPort, accept returned error")
			// for non-temporary or unrecognized errors, just return (this
			// typically means the listener has shut down)
			return
		}

		log.Printf("in forwardPort, accept returned success")
		toConn, err := net.Dial("tcp", forwardTo)
		log.Printf("in forwardPort, dialing")
		if err != nil {
		log.Printf("in forwardPort, dialing failed")
			// simulate a connection refused to the fromConn, although this will
			// appear as an immediate close, rather than a refusal
			fromConn.Close()
			continue
		}

		log.Printf("in forwardPort, dialing succeeded")
		forward := func(f net.Conn, t net.Conn) {
		log.Printf("in forwardPort, copying data")
			defer func() {
		log.Printf("in forwardPort, closing t")
			t.Close()
		}()
			_, _ = io.Copy(t, f)
		}
		go forward(fromConn, toConn)
		go forward(toConn, fromConn)
	}
	// unreachable
	//log.Printf("end of forwardPort")
}
